package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"

	pb "github.com/h1de27/my-first-grpc/echo/proto"
)

func main() {
	target := "localhost:50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	client := pb.NewEchoServiceClient(conn)
	msg := os.Args[1]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Echo(ctx, &pb.EchoRequest{Message:msg})
	if err != nil {
		log.Println(err)
	}
	log.Println(r.GetMessage())
}
