package main

import (
	"context"
	"log"
	"fmt"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) RegisterUser(ctx context.Context, in *pb.UserRequest) (*pb.UserId, error) {
	log.Printf("***** Register user invoked with %v\n", in)
	userIdInc = userIdInc + 1
	userExists := checkIfUserExists(in.Email)
	if userExists {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("User with requested email already exists!"),
		)
	}
	user := User{ID: userIdInc, FirstName: in.FirstName, LastName: in.LastName, Email: in.Email}
	users = append(users, user)
	log.Printf("---users start-----\n")
	for _, user := range users {
        fmt.Printf("ID: %d, FirstName: %s, LastName: %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
    }
	log.Printf("---users end-----\n")
	log.Printf("***** Register user completed")
	return &pb.UserId{Id: userIdInc}, nil
}

func checkIfUserExists(email string) bool {
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}
	return false
}