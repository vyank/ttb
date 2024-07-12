package main

import (
	"context"
	"log"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
)

func registerUser(c pb.TicketBookingClient) {
	log.Println("-----Register User started--------")
	r, err := c.RegisterUser(context.Background(), &pb.UserRequest{FirstName: "Vyankatesh",LastName: "Inamdar", Email: "a@b.com"})
	if err != nil {
		log.Fatalf("Could not register user: %v\n", err)
	}
	log.Printf("User Id : %v\n", r.Id)
	log.Println("-----Register User ended--------")
}
