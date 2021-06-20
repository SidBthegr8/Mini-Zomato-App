package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	

	"github.com/jinzhu/gorm"
	pb "mini-zomato/backend/grpcServer"
	"google.golang.org/grpc"
	"mini-zomato/backend/Models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)

var (
	port = flag.Int("port", 10000, "The gRPC server port")
)


var db *gorm.DB
var err error



type grpcServer struct {
	pb.UnimplementedAddGetRestaurantServer
	
}

func (s *grpcServer) ListRestaurants(ctx context.Context, req *pb.GetListRequest) (*pb.RestaurantList, error) {
	
	var tempRestaurants []Models.Restaurant
	db.Order("rating desc").Find(&tempRestaurants)

	
	urlsJson, _ := json.Marshal(tempRestaurants)
	urls := &pb.RestaurantList{RestaurantList: string(urlsJson)}
	return urls, err
}


func (s *grpcServer) AddRestaurant(ctx context.Context, restaurant *pb.Restaurant) (*pb.AddedConfirmation, error) {
	
	newRestaurant := Models.Restaurant{
		Name:        restaurant.Name,
		Rating:      restaurant.Rating,
		Cuisine:     restaurant.Cuisine,
		Address:     restaurant.Address,
		OpeningTime: restaurant.OpeningTime,
		ClosingTime: restaurant.ClosingTime,
		CostForTwo:  restaurant.CostForTwo,
	}

	db.Create(&newRestaurant)
	return &pb.AddedConfirmation{AddedConfirmation: "successfully added new restaurant"}, nil
}



func main() {

	

	db, err = gorm.Open("mysql", "root:panamacanal@tcp(127.0.0.1:3306)/Football?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection Failed to open")
	} else {
		log.Println("Connection established with database")
	}
	db.AutoMigrate(&Models.Restaurant{})

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAddGetRestaurantServer(s, &grpcServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
