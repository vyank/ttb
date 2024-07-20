package main

import (
	"context"
	"log"
	"fmt"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) GetTicketDetails(ctx context.Context, in *pb.UserId) (*pb.TicketDetailsResponse, error) {
	log.Printf("***** get ticket details invoked with %v\n", in)
	user := getUser(in.Id)
	if user == nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid user"),
		)
	}
	seat := getBookedSeat(in.Id)
	if seat == nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("User does not have any booking"),
		)
	}
	userResp := &pb.UserRequest{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
	}
	log.Printf("***** get ticket details completed\n")
	return &pb.TicketDetailsResponse{
		Id: seat.ID,
		Section: seat.section,
		SeatNum: seat.seatNum,
		From: trains[0].from,
		To: trains[0].to,
		Price: seat.price,
		User: userResp,
	}, nil
}

func getBookedSeat(userId int32) (*Seat) {
	for i, seat := range seats {
		if seat.userId == userId {
			return &seats[i];
		}
	}
	return nil;
}

func getUser(userId int32) (*User) {
	for i, user := range users {
		if user.ID == userId {
			return &users[i];
		}
	}
	return nil;
}