package main

import (
	"context"
	"log"
	"fmt"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*Server) RemoveUser(ctx context.Context, in *pb.UserId) (*emptypb.Empty, error) {
	log.Printf("***** Remove user invoked with %v\n", in)
	cancelBookedSeat(in.Id)
	deleteUser(in.Id)
	log.Printf("---users-----\n")
	for _, user := range users {
        fmt.Printf("ID: %d, FirstName: %s, LastName: %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
    }
	log.Printf("---users-----\n")
	log.Printf("***** Remove user completed\n")
	return &emptypb.Empty{}, nil
}

func (*Server) RemoveUserBySeat(ctx context.Context, in *pb.BookedSeat) (*emptypb.Empty, error) {
	log.Printf("***** Remove user by seat invoked with %v\n", in)
	
	seat := getUserBySeat(in.SeatNum, in.Section);
	log.Println("got the seat")
	log.Println(seat)
	cancelBookedSeat(seat.userId)
	deleteUser(seat.userId)
	log.Printf("---users-----\n")
	for _, user := range users {
        fmt.Printf("ID: %d, FirstName: %s, LastName: %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
    }
	log.Printf("---users-----\n")
	log.Printf("***** Remove user completed\n")
	return &emptypb.Empty{}, nil
}

func getUserBySeat(seatNum int32, section string) (*Seat) {
	for _, seat := range seats {
		if seat.seatNum == seatNum && seat.section == section {
			return &seat
		}
	}
	return nil;
}

func cancelBookedSeat(userId int32) {
	for i, seat := range seats {
		if seat.userId == userId {
			seats[i].userId = 0
			seats[i].booked = false
		}
	}
}

func deleteUser(userId int32) {
	index := 0
	for i, user := range users {
		if user.ID == userId {
			index = i
		}
	}
	users = append(users[:index], users[index+1:]...)
}