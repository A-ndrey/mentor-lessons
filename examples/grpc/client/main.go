package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"mentor-space/examples/grpc/pb"
	"sync"
)

func main() {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := pb.NewExampleClient(conn)
	Unary(client)
	//ServerStream(client)
	//ClientStream(client)
	//BiDirectional(client)
}

func Unary(client pb.ExampleClient) {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "key1", "value1", "key2", "value2")

	response, err := client.Unary(ctx, &pb.Request{Num: 1234})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(response.GetNum())
}

func ServerStream(client pb.ExampleClient) {
	stream, err := client.ServerStream(context.Background(), &pb.Request{Num: 5})
	if err != nil {
		log.Println(err)
		return
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			return
		}

		log.Println(response.GetNum())
	}
}

func ClientStream(client pb.ExampleClient) {
	stream, err := client.ClientStream(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	for i := 1; i <= 5; i++ {
		err := stream.Send(&pb.Request{Num: int64(i)})
		if err != nil {
			log.Println(err)
			return
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(response.GetNum())
}

func BiDirectional(client pb.ExampleClient) {
	stream, err := client.BiDirectStream(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			recv, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Println(err)
				return
			}

			log.Println(recv.GetNum())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			err := stream.Send(&pb.Request{Num: int64(i)})
			if err != nil {
				log.Println(err)
				return
			}
		}

		err := stream.CloseSend()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	wg.Wait()
}
