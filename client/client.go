package main

//client.go

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "vip_services/vip_services"
)

const (
	address  = "localhost:50051"
	apply_id = 55445929
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDoumiBApplyClient(conn)

	r, err := c.GetOneByApplyId(context.Background(), &pb.ApplyIdRequest{ApplyId: apply_id})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.ApplyInfo)
}
