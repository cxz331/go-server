package main

// server.go

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"vip_services/controller"
	"vip_services/models"
	pb "vip_services/vip_services"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetOneByApplyId(ctx context.Context, in *pb.ApplyIdRequest) (*pb.ApplyIdReply, error) {

	Apply := new(controller.ApplyController)

	applyInfo := Apply.GetOneApplyById(in)

	return &pb.ApplyIdReply{ApplyInfo: applyInfo}, nil

}

func (s *server) GetAllByApplyUid(ctx context.Context, in *pb.ApplyUidRequest) (*pb.ApplyUidReply, error) {

	Apply := new(controller.ApplyController)

	applyInfo := Apply.GetAllApplyByUserId(in)

	return &pb.ApplyUidReply{ApplyInfo: applyInfo}, nil

}

func main() {

	models.Init()

	lis, err := net.Listen("tcp", port)

	if err != nil {

		log.Fatal("failed to listen: %v", err)

	}

	s := grpc.NewServer()

	pb.RegisterDoumiBApplyServer(s, &server{})

	s.Serve(lis)

}
