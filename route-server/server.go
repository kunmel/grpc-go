package main

import (
	"context"
	"google.golang.org/protobuf/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "my-grpc-go/route"
)

type routeGuideServer struct {
	db []*pb.Feature
	pb.UnimplementedRouteGuideServer
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _,feature := range s.db {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil,nil
}
func (s *routeGuideServer) ListFeature(*pb.Rectangle, pb.RouteGuide_ListFeatureServer) error {
	return nil
}

func (s *routeGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error {
	return nil
}

func (s *routeGuideServer) Recommend(pb.RouteGuide_RecommendServer) error {
	return nil
}

func newServer() *routeGuideServer {
	return &routeGuideServer{
		db: []*pb.Feature{
			{Name:"北京理工大学", Location: &pb.Point{
				Latitude: 123,
				Longitude: 321,
			}},
		},
	}
}

func main() {
	lis, err := net.Listen("tcp", "10.108.16.218:5000")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	log.Fatalln(grpcServer.Serve(lis))
}
