package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"

	pb "github.com/kvonbredow/rpc_demo/proto"
	grpc "google.golang.org/grpc"
)

type addFiveServer struct{}

func (*addFiveServer) AddFive(stream pb.AddFive_AddFiveServer) error {
	fmt.Println("Entering AddFive")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fmt.Printf("Received: %d\n", req.Num)
		time.Sleep(time.Second)
		resp := &pb.AddFiveResponse{
			Result: req.Num + 5,
		}
		stream.Send(resp)
	}
	fmt.Println("Leaving AddFive")
	return nil
}

func newAddFiveServer() pb.AddFiveServer {
	return &addFiveServer{}
}

func main() {
	port := flag.Int("port", 65432, "Port for the server to run on")
	flag.Parse()

	fmt.Printf("Starting AddFive Server on port %d...\n", *port)
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAddFiveServer(grpcServer, newAddFiveServer())
	grpcServer.Serve(conn)
}
