package main

import (
	"flag"
	"fmt"
	"io"

	pb "github.com/kvonbredow/rpc_demo/proto"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func main() {
	addr := flag.String("port", "127.0.0.1:65432", "Port for the client to dial")
	num := flag.Int("num", 8, "The number which will be sent")
	flag.Parse()

	ctx := context.Background()

	conn, err := grpc.Dial(*addr, []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		fmt.Printf("grpc.Dial(): %v", err)
		fmt.Println()
		return
	}
	defer conn.Close()

	client := pb.NewAddFiveClient(conn)
	if client == nil {
		fmt.Printf("Failed to make client")
		fmt.Println()
		return
	}

	inFlight := make(chan struct{}, 10)
	req := &pb.AddFiveRequest{
		Num: int32(*num),
	}
	stream, err := client.AddFive(ctx)
	if err != nil {
		fmt.Printf("client.AddFive(): %v", err)
		fmt.Println()
		return
	}
	defer stream.CloseSend()
	go func() {
		for {
			inFlight <- struct{}{}
			fmt.Printf("Request: %d\n", req.Num)
			stream.Send(req)
		}
	}()
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		<-inFlight
		fmt.Printf("Response: %d\n", resp.Result)
	}
}
