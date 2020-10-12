package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	pb "github.com/hisheng/dataStructure/zskiplist/grpc"

)

var (
	grpcServerAddress = "0.0.0.0:8989"
	grpcCmd           = &cobra.Command{
		Use: "grpc",
		Run: func(cmd *cobra.Command, args []string) {
			println("listen: ", grpcServerAddress)
			conn, err := net.Listen("tcp", grpcServerAddress)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			grpcServer := grpc.NewServer()
			pb.RegisterAdvServer(grpcServer, pb.NewZskiplistService())
			reflection.Register(grpcServer)
			if err = grpcServer.Serve(conn); err != nil {
				log.Fatalf("grpcServer Server: %v", err)
			}
		},
	}
)

func init() {
	addr := os.Getenv("GRPC_SERVER_ADDRESS")
	if addr != "" {
		grpcServerAddress = addr
	}
}
