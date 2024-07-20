package main

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
)

var addr string = "localhost:50051"

func main() {
	opts := []grpc.DialOption{}
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, creds)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	registerUser(c)
	purchaseTicket(c)
	ticketDetails(c)// section A: seat 1
	// viewSeatAllocation(c)
	// modifyTicket(c)
	// removeUser(c)
	removeUserBySeat(c)
}
