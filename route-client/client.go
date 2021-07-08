package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "my-grpc-go/route"
)

func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude: 123,
		Longitude: 321,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}


func main() {
	conn, err := grpc.Dial("10.108.16.218:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	}

	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	runFirst(client)
}
