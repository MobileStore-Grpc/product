package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/MobileStore-Grpc/product/pb"
	"github.com/MobileStore-Grpc/product/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dail server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	mobileClient := pb.NewMobileServiceClient(conn)

	req := &pb.CreateMobileRequest{
		Mobile: sample.NewMobile(),
	}
	//set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//This pb.LaptopServiceClient will execute method CreateLaptop implemented by "LaptopServe struct" in laptop_server.go  and register in server/main.go using "pb.RegisterLaptopServiceServer(grpcServer, laptopServer)"
	res, err := mobileClient.CreateMobile(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("mobile already exists")
		} else {
			log.Fatal("cannot create mobile: ", err)
		}
	}

	log.Printf("create laptop with id: %s", res.Id)
}
