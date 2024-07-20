package main

import (
	"log"
	"net"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
	"google.golang.org/grpc"
)

type User struct {
    ID    int32
    FirstName  string
    LastName string
	Email string
}

type Train struct {
	ID int32
	from string
	to string 
	price int32
}

type Seat struct {
    ID    int32
    seatNum  int32
    section string
	booked bool
	userId int32
	price int32
}

var trains []Train;
var users []User;
var seats []Seat;
var userIdInc int32

var addr string = "0.0.0.0:50051"

func main() {
	//init
	userIdInc = 0
	trains = []Train{
		{ID: 1, from: "London", to: "Paris", price: 20},
	}
	price1 := int32(20)
	seats = []Seat{
		{ID: 1, section: "A", seatNum: 1, booked: false, price: price1},
		{ID: 2, section: "A", seatNum: 2, booked: false, price: price1},
		{ID: 3, section: "A", seatNum: 3, booked: false, price: price1},
		{ID: 4, section: "A", seatNum: 4, booked: false, price: price1},
		{ID: 5, section: "A", seatNum: 5, booked: false, price: price1},
		{ID: 6, section: "B", seatNum: 1, booked: false, price: price1},
		{ID: 7, section: "B", seatNum: 2, booked: false, price: price1},
		{ID: 8, section: "B", seatNum: 3, booked: false, price: price1},
		{ID: 9, section: "B", seatNum: 4, booked: false, price: price1},
		{ID: 10, section: "B", seatNum: 5, booked: false, price: price1},
	}
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()
	log.Printf("Listening at %s\n", addr)
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	pb.RegisterTicketBookingServer(s, &Server{})
	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
