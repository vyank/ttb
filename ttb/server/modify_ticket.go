package main

import (
	"context"
	"log"
	"fmt"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) ModifySeat(ctx context.Context, in *pb.ModifySeatRequest) (*pb.TicketPurchaseResponse, error) {
	log.Printf("******* Modify ticket invoked %v\n", in.UserId)
	seat := modifySeat(in.CurrTicketId, in.NewSeatNum, in.NewSection, in.UserId)
	if seat == nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Seat can not be modified"),
		)
	}
	log.Printf("-----Seats start------")
	for _, seat := range seats {
        fmt.Printf("ID: %v, seatNum: %v, section: %v, booked: %v, userId: %v\n", seat.ID, seat.seatNum, seat.section, seat.booked, seat.userId)
    }
	log.Printf("-----Seats end------")
	log.Printf("******* Modify ticket completed\n")
	return &pb.TicketPurchaseResponse{Id: seat.ID, Section: seat.section, SeatNum: seat.seatNum}, nil
}

func modifySeat(currTicketId int32, newSeatNum int32, newSection string, userId int32) (*Seat) {
	newSeatIndex := -1
	for i, seat := range seats {
		if seat.seatNum == newSeatNum && seat.section == newSection && !seat.booked {
			newSeatIndex = i
		}
	}
	if newSeatIndex == -1 {
		log.Printf("Requested seat is not available")
		return nil
	} else {
		cacelledCurrentTicket := false
		for i, seat := range seats {
			if seat.ID == currTicketId && seat.userId == userId{
				cacelledCurrentTicket = true
				seats[i].booked = false;
				seats[i].userId = 0;
			}
		}
		if cacelledCurrentTicket {
			seats[newSeatIndex].booked = true;
			seats[newSeatIndex].userId = userId;
			return &seats[newSeatIndex]
		} else {
			log.Printf("Current ticket booking is not booked by requesting user")
			return nil;
		}
	}
	return nil;
}