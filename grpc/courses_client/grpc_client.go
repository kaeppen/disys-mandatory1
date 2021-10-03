package main

import (
	"context"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	//tutorial fyr har noget "pb" et eller andet med her...
)

//serverens adresse
const (
	address = "localhost:8080"
)

func main() {
	//establish connection to server (insecure, blocking call)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No connection: %v", err)
	}
	//close the connection when everything else has run
	defer conn.Close()
	client := pb.NewCoursesClient(conn)

	//new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//hardcoding of new courses (in the short NewCourse format)
	var new_courses = []course{
		{ID: "1", Name: "Dømat", Workload: 9001.0},
		{ID: "2", Name: "BPAK", Workload: 0.0},
		{ID: "3", Name: "GrPro", Workload: 15.0},
	}
	for course := range new_courses {
		r, err := client.AddCourse(ctx, &pb.NewCourse{ID: course.ID, Name: course.Name, Workload: course.Workload})
		if err != nil {
			log.Fatalf("could not add course %v", err)
		}
		log.Printf(`Course details: 
		ID: %d
		NAME: %s
		WORKLOAD: %d`, r.GetID(), r.GetName(), r.GetWorkload())
	}

}
