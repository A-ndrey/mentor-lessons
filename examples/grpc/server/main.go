package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net"
)

//protoc --go_out=. --go-grpc_out=. proto/*

type grpcHandler struct {
	pb.UnimplementedExampleServer
}

func (g *grpcHandler) Unary(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	result := request.GetNum() % 10

	return &pb.Response{Num: result}, nil
}

func (g *grpcHandler) ServerStream(request *pb.Request, server pb.Example_ServerStreamServer) error {
	for i := int64(0); i < request.GetNum(); i++ {
		err := server.Send(&pb.Response{Num: i})
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *grpcHandler) ClientStream(server pb.Example_ClientStreamServer) error {
	var sum int64
	defer func() {
		err := server.SendAndClose(&pb.Response{Num: sum})
		if err != nil {
			log.Println(err)
		}
	}()

	for {
		req, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		sum += req.GetNum()
	}

	return nil
}

func (g *grpcHandler) BiDirectStream(server pb.Example_BiDirectStreamServer) error {
	for {
		req1, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		req2, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		mul := req1.GetNum() * req2.GetNum()

		err = server.Send(&pb.Response{Num: mul})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Println(err)
		return
	}
	defer listener.Close()

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(loggerInterceptor))
	pb.RegisterExampleServer(grpcServer, &grpcHandler{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
		return
	}
}

func loggerInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Println(md)

	}

	req2, ok := req.(*pb.Request)
	if !ok {
		return nil, errors.New("unknown request type")
	}

	log.Println("request: ", req2)

	resp, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}

	resp2, ok := resp.(*pb.Response)
	if !ok {
		return nil, errors.New("unknown response type")
	}

	log.Println("response: ", resp2)

	return resp, nil
}
