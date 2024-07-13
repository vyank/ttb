package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vyank/train-ticket-app/ttb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {

	userIdInc = 0

	trains = []Train{
		{ID: 1, from: "London", to: "Paris", price: 20},
	}

	seats = []Seat{
		{ID: 1, section: "A", seatNum: 1, booked: false},
		{ID: 2, section: "A", seatNum: 2, booked: false},
		{ID: 3, section: "A", seatNum: 3, booked: false},
		{ID: 4, section: "A", seatNum: 4, booked: false},
		{ID: 5, section: "A", seatNum: 5, booked: false},
		{ID: 6, section: "B", seatNum: 1, booked: false},
		{ID: 7, section: "B", seatNum: 2, booked: false},
		{ID: 8, section: "B", seatNum: 3, booked: false},
		{ID: 9, section: "B", seatNum: 4, booked: false},
		{ID: 10, section: "B", seatNum: 5, booked: false},
	}

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterTicketBookingServer(s, &Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v\n", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
