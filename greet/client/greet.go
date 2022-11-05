package main

import (
	"context"
	"log"

	pb "github.com/harsha-fsd/Go-gRPC-Example/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Sriharsha",
	})
	if err != nil {
		log.Fatalf("Couldn't greet: %v\n", err)
	}
	log.Printf("Greeeting: %s\n", res.Result)

}
