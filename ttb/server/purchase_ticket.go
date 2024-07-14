package main

import (
	"context"
	"log"
	"fmt"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) PurchaseTicket(ctx context.Context, in *pb.TicketPurchaseRequest) (*pb.TicketPurchaseResponse, error) {
	log.Printf("***** Purchase ticket invoked %v\n", in.UserId)
	user := getUser(in.UserId)
	if user == nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid user"),
		)
	}
	seat := getNextAvailableSeat(seats, in.UserId)
	if seat == nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("No more seats available"),
		)
	}
	fmt.Printf("ID: %v, seatNum: %v, section: %v, booked: %v, userId: %v\n", seat.ID, seat.seatNum, seat.section, seat.booked, seat.userId)
	log.Printf("-----Seats start------")
	for _, seat := range seats {
        fmt.Printf("ID: %v, seatNum: %v, section: %v, booked: %v, userId: %v\n", seat.ID, seat.seatNum, seat.section, seat.booked, seat.userId)
    }
	log.Printf("-----Seats end------")
	log.Printf("***** Purchase ticket completed ")
	return &pb.TicketPurchaseResponse{Id: seat.ID, Section: seat.section, SeatNum: seat.seatNum}, nil
}

func getNextAvailableSeat(seats []Seat, userId int32) (*Seat) {
	for i, seat := range seats {
		if !seat.booked {
			seats[i].booked = true;
			seats[i].userId = userId;
			return &seats[i];
		}
	}
	return nil;
}