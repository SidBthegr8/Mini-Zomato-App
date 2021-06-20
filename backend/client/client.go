package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	pb "mini-zomato/backend/grpcServer"
	"google.golang.org/grpc"
	"mini-zomato/backend/Models"

)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)



func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	
}

func findRestaurants(w http.ResponseWriter, r *http.Request) {
	
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAddGetRestaurantClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := client.ListRestaurants(ctx, &pb.GetListRequest{GetListRequest: "GET"})
	if err != nil {
		log.Fatalf("Coud not get restaurants list: %v", err)
	}

	var receivedList []Models.Restaurant
	json.Unmarshal([]byte(result.RestaurantList), &receivedList)

	w.Write([]byte(result.RestaurantList))

}

func addRestaurant(w http.ResponseWriter, r *http.Request) {
	
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")


	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAddGetRestaurantClient(conn)

	var res Models.Restaurant
	s, _ := io.ReadAll(r.Body)
	json.Unmarshal([]byte(s), &res)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := client.AddRestaurant(ctx, &pb.Restaurant{
		Name:        res.Name,
		Rating:      res.Rating,
		Cuisine:     res.Cuisine,
		Address:     res.Address,
		OpeningTime: res.OpeningTime,
		ClosingTime: res.ClosingTime,
		CostForTwo:  res.CostForTwo,
	})
	if err != nil {
		log.Fatalf("could not add restaurant: %v", err)
	}

	fmt.Printf(result.AddedConfirmation)
	fmt.Printf(res.Rating)
}

func handleRequests(){
	
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/restaurants", findRestaurants)
	myRouter.HandleFunc("/add", addRestaurant)
	http.ListenAndServe(":5000", myRouter)

}
func main(){
	
	
	handleRequests()

}
