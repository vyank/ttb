package main

import (
	"context"
	"log"

	pb "github.com/vyank/train-ticket-app/ttb/proto"
)

func modifyTicket(c pb.TicketBookingClient) {
	log.Println("-------Modify ticket started-------")
	r, err := c.ModifySeat(context.Background(), &pb.ModifySeatRequest{
		CurrTicketId: 1,
		NewSeatNum: 2,
		NewSection: "B",
		UserId: 1,      
	})
	if err != nil {
		log.Fatalf("Could not Modify ticket: %v\n", err)
	}
	log.Printf("Modify success, Id : %v, Section :  %v, SeatNum :  %v\n", r.Id, r.Section, r.SeatNum)
	log.Println("-------Modify ticket completed-------")
}
