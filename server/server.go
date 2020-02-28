package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

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
)

type CommentServer struct {
	commentCh chan string
}

func newServer() *CommentServer {
	s := &CommentServer{
		commentCh: make(chan string),
	}
	s.startScan()

	return s
}

func (s *CommentServer) startScan() {
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			s.commentCh <- scanner.Text()
		}
	}()
}

func (s *CommentServer) GetComment(req *pb.Filter, stream pb.CommentService_GetCommentServer) error {
	dummyID := 0
	for {
		c := <-s.commentCh
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
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCommentServiceServer(grpcServer, newServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
