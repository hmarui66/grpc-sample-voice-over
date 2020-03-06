package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/satori/go.uuid"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	pb "github.com/hmarui66/grpc-sample-voice-over/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/testdata"
)

var (
	tls      = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 8080, "The server port")
	token    = flag.String("token", "foo", "Token to call a procedure")

	clients *clientsManager
)

func init() {
	clients = newClientsManager()
}

type CommentServer struct{}

func newServer() *CommentServer {
	s := &CommentServer{}
	go s.startScan()

	return s
}

func (s *CommentServer) startScan() {
	log.Printf("starting stdin scanner...\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		clients.pushMsg(&comment{
			Type:      "NONE",
			User:      "gRPC-SERVER",
			Text:      scanner.Text(),
			Timestamp: time.Now().String(),
			Channel:   "NONE",
		})
	}
}

func (s *CommentServer) GetComment(req *pb.Filter, stream pb.CommentService_GetCommentServer) error {
	id := uuid.NewV4()
	ch := clients.add(id.String())
	defer clients.delete(id.String())

	for {
		c := <-ch
		if err := stream.Send(&pb.Comment{
			Type:      c.Type,
			User:      c.User,
			Text:      c.Text,
			Timestamp: c.Timestamp,
			Channel:   c.Channel,
		}); err != nil {
			return status.Errorf(codes.Unknown, "failed to send a comment: %v", err)
		}
	}
}

func streamAccessLogHandler(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()
	err := handler(srv, ss)
	sts := status.Convert(err)

	log.Printf("%s [duration: %d] - [status: %s] %s %v",
		info.FullMethod,
		time.Since(start).Nanoseconds()/int64(time.Millisecond),
		sts.Code(),
		sts.Message(),
		sts.Details(),
	)

	return err
}

func main() {
	flag.Parse()

	go startSlackServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	opts = append(opts,
		grpc_middleware.WithStreamServerChain(
			streamAccessLogHandler,
			grpc_auth.StreamServerInterceptor(authFunc),
		),
	)

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCommentServiceServer(grpcServer, newServer())

	log.Printf("starting grpc server...\n")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func authFunc(ctx context.Context) (context.Context, error) {
	received, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	if received != *token {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	newCtx := context.WithValue(ctx, "result", "ok")
	return newCtx, nil
}
