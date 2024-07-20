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
	log.Printf("***** Purchase ticket invoked %v\n", in.DiscoutCode)
	user := getUser(in.UserId)
	discout := int32(0)
	if in.DiscoutCode == "D1" {
		discout = int32(10)
	} else if in.DiscoutCode == "D2" {
		discout = int32(20)
	} else if in.DiscoutCode == "D3" {
		discout = int32(30)
	} else {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid Discout code"),
		)
	}
	if user == nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid user"),
		)
	}
	seat := getNextAvailableSeat(seats, in.UserId, discout)
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

func getNextAvailableSeat(seats []Seat, userId int32, discout int32) (*Seat) {
	for i, seat := range seats {
		if !seat.booked {
			seats[i].booked = true;
			seats[i].userId = userId;
			if seats[i].price - discout < int32(0) {
				seats[i].price = int32(0);
			} else {
				seats[i].price = seats[i].price - discout;
			}
			
			return &seats[i];
		}
	}
	return nil;
}