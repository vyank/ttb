package main

import (
	"context"
	"log"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
)

func purchaseTicket(c pb.TicketBookingClient) {
	log.Println("-----Purchase ticket started--------")
	r, err := c.PurchaseTicket(context.Background(), &pb.TicketPurchaseRequest{
		From: "London",
		To: "Paris", 
		UserId: 1,
		Price: 20,
	})
	if err != nil {
		log.Fatalf("Could not Purchase ticket: %v\n", err)
	}
	log.Printf("Purchase success, Id : %v, Section :  %v, SeatNum :  %v\n", r.Id, r.Section, r.SeatNum)
	log.Println("-----Purchase ticket completed--------")
}