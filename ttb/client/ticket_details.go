package main

import (
	"context"
	"log"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
)

func ticketDetails(c pb.TicketBookingClient) {
	log.Println("------Get ticket details started-------")
	r, err := c.GetTicketDetails(context.Background(), &pb.UserId{
		Id: 1,
	})
	if err != nil {
		log.Fatalf("Could not get ticket details: %v\n", err)
	}
	log.Printf("Ticket details : %v", r)
	log.Println("------Get ticket details completed-------")
}
