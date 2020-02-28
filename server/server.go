package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"grpc-go/status"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc/codes"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"

	pb "github.com/hmarui66/grpc-sample-voice-over/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 8080, "The server port")

	commentCh chan string
)

func init() {
	commentCh = make(chan string)
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
		commentCh <- scanner.Text()
	}
}

func (s *CommentServer) GetComment(req *pb.Filter, stream pb.CommentService_GetCommentServer) error {
	dummyID := 0
	for {
		c := <-commentCh
		dummyID += 1
		if err := stream.Send(&pb.Comment{
			Id:    string(dummyID),
			Value: c,
		}); err != nil {
			return fmt.Errorf("failed to send comment: %w", err)
		}
	}
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
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			func(
				ctx context.Context, req interface{},
				info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
			) (interface{}, error) {
				startTime := time.Now()

				resp, err := handler(ctx, req)

				duration := time.Since(startTime)

				code := codes.Unknown
				if sts, ok := status.FromError(err); ok {
					code = sts.Code()
				}

				var fn func(format string, args ...interface{})
				if err != nil {
					fn = func(format string, args ...interface{}) {
						log.Printf(format+`, error="%s"`, append(args, err)...)
					}
				}

				fn("method = %s, code = %s, duration: %0.4fms", info.FullMethod, code, toMilliseconds(duration))

				return resp, err
			},
			grpc_validator.UnaryServerInterceptor(),
		)))

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCommentServiceServer(grpcServer, newServer())

	log.Printf("starting grpc server...\n")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func toMilliseconds(d time.Duration) float64 {
	msec := d / time.Millisecond
	nsec := d % time.Millisecond
	return float64(msec) + float64(nsec)/1e6
}
