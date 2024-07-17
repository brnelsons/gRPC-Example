package main

import (
	"context"
	"flag"
	"gRPC/src/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	serverAddr := flag.String("server", "localhost:8080", "the server address in the format for host:port")
	flag.Parse()
	log.Println("making connection to server:", *serverAddr)

	// TODO enable TLS
	//tlsConfig := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})
	opts := []grpc.DialOption{
		//grpc.WithTransportCredentials(tlsConfig),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalln("could not create grpc client", err)
	}
	cc := pb.NewCalculatorClient(conn)

	// make the request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := cc.Sum(ctx, &pb.NumbersRequest{
		Numbers: []int64{2, 55, 72, 99},
	})
	if err != nil {
		log.Fatalln("error sending request:", err)
	}
	log.Println(res.GetResult())
}
