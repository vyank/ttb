package main

import (
	"context"
	"log"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
)

func viewSeatAllocation(c pb.TicketBookingClient) {
	log.Println("-----view seat allocation started-------")
	r, err := c.ViewSeatAllocation(context.Background(), &pb.Section{
		Section: "A",
	})
	if err != nil {
		log.Fatalf("Could not get seat allocation details: %v\n", err)
	}
	log.Printf("Seat allocation : %v", r)
	log.Println("-----view seat allocation completed-------")
}
