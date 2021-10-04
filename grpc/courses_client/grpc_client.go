package main

import (
	"context"
	"log"
	"time"

	pb "course/grpc"

	"google.golang.org/grpc"
	//tutorial fyr har noget "pb" et eller andet med her...
)

//serverens adresse
const (
	address = "localhost:8080"
)

type course struct {
	ID       string   `json:"id"`
	Rating   float64  `json:"rating"`
	Name     string   `json:"name"`
	Workload float64  `json:"workload"`
	Students []string `json:"students"`
}

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
	for i := 0; i < len(new_courses); i++ {
		r, err := client.AddCourse(ctx, &pb.NewCourse{ID: new_courses[i].ID, Name: new_courses[i].Name, Workload: new_courses[i].Workload})
		if err != nil {
			log.Fatalf("could not add course %v", err)
		}
		log.Printf(`Course details: 
		ID: %s
		NAME: %s
		WORKLOAD: %v`, r.GetID(), r.GetName(), r.GetWorkload()) //brug %f i workload for at få decimaler med(mange though)
	}

}
