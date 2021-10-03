package main

import (
	"context"
	"log"
	"net"

	pb "./courses.pb.go"

	"google.golang.org/grpc"
	//tutorial fyr har noget "pb" et eller andet med her...
)

//porte for the gRPC server to run on
const (
	port = ":8080"
)

type courseServer struct {
	pb.UnimplementedCoursesServer
}

func (s *courseServer) AddCourse(ctx context.Context, in *pb.NewCourse) (*pb.Course, error) {
	log.Printf("Recieved: %v", in.GetName())
	return &pb.Course{Id: in.GetID(), Name: in.getName(), Workload: in.GetWorkload()}, nil
}

/*
func (s *courseServer) DeleteCourse(ctx context.Context, in *pb.Int32Value) (*pb.Course, error) {
	log.Printf("Recieved: %v", in)
	//hvor skal de lagres henne?
}*/

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCoursesServer(server, &CoursesServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
