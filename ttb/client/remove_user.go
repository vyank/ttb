package main

import (
	"context"
	"log"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
)

func removeUser(c pb.TicketBookingClient) {
	log.Println("----Remove user started-----")
	_, err := c.RemoveUser(context.Background(), &pb.UserId{
		Id: 1,
	})
	if err != nil {
		log.Fatalf("Could not remove user: %v\n", err)
	}
	log.Printf("User and all his/her bookings are removed")
	log.Println("----Remove user completed-----")
}

func removeUserBySeat(c pb.TicketBookingClient) {
	log.Println("----Remove user by seat started-----")
	_, err := c.RemoveUserBySeat(context.Background(), &pb.BookedSeat{
		SeatNum: 1,
		Section: "A",
	})
	if err != nil {
		log.Fatalf("Could not remove user: %v\n", err)
	}
	log.Printf("User and all his/her bookings are removed")
	log.Println("----Remove user completed-----")
}
