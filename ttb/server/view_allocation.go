package main

import (
	"context"
	"log"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
)

func (*Server) ViewSeatAllocation(ctx context.Context, in *pb.Section) (*pb.SeatAllocationResponse, error) {
	log.Printf("***** view seat allocation invoked with %v\n", in)
	seatsResp := getSeats(in.Section)
	seatsResp1 := []*pb.Seat{}
	for _, seatItem := range seatsResp {
		user := getUser(seatItem.userId)
		if user != nil {
			seatsResp1 = append(seatsResp1, &pb.Seat{
				SeatNum: seatItem.seatNum,
				User: &pb.UserRequest{
					FirstName: user.FirstName,
					LastName: user.LastName,
					Email: user.Email,
				},
				Available: "No",
			})
		} else {
			seatsResp1 = append(seatsResp1, &pb.Seat{
				SeatNum: seatItem.seatNum,
				Available: "Yes",
			})
		}
		
	}
	log.Printf("***** view seat allocation completed\n")
	return &pb.SeatAllocationResponse{
		Seats: seatsResp1,
	}, nil
}

func getSeats(section string) ([]Seat) {
	seatsResp := []Seat{}
	for _, seat := range seats {
		if seat.section == section {
			seatsResp = append(seatsResp, seat)
		}
	}
	return seatsResp;
}