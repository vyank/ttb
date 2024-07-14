package main

import (
	"context"
	"testing"
	pb "github.com/vyank/train-ticket-app/ttb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegister2Users(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	user := &pb.UserRequest{
		FirstName: "Vyankatesh",LastName: "Inamdar", Email: "a@b.com",
	}
	exptectUserId  := int32(1);
	res, err := c.RegisterUser(context.Background(), user)
	if err != nil {
		t.Errorf("Register user (%v) got unexpected error", err)
	}
	if res.Id != exptectUserId {
		t.Errorf("User register, created Id = %v, expected %v", res.Id, exptectUserId)
	}
	user = &pb.UserRequest{
		FirstName: "Sam",LastName: "Kalk", Email: "c@b.com",
	}
	exptectUserId  = int32(2);
	res, err = c.RegisterUser(context.Background(), user)
	if err != nil {
		t.Errorf("Register user (%v) got unexpected error", err)
	}
	if res.Id != exptectUserId {
		t.Errorf("User register, created Id = %v, expected %v", res.Id, exptectUserId)
	}
	c.RegisterUser(context.Background(), &pb.UserRequest{
		FirstName: "Dan",LastName: "Kalk", Email: "d@b.com",
	})

	c.RegisterUser(context.Background(), &pb.UserRequest{
		FirstName: "Van",LastName: "Kalk", Email: "e@b.com",
	})
}

func TestAddSameUserError(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	user := &pb.UserRequest{
		FirstName: "Vyankatesh",LastName: "Inamdar", Email: "a@b.com",
	}
	_, err = c.RegisterUser(context.Background(), user)
	if err == nil {
		t.Error("Expected Error")
	}
	e, ok := status.FromError(err)
	if !ok {
		t.Error("Expected error")
	}
	if e.Code() != codes.Internal {
		t.Errorf("Expected Internal, got %v", e.Code().String())
	}
}

func TestPurchaseTicket(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	r, err := c.PurchaseTicket(context.Background(), &pb.TicketPurchaseRequest{
		From: "London",
		To: "Paris", 
		UserId: 1,
		Price: 20,
	})
	if err != nil {
		t.Errorf("Purhase ticket got unexpected error %v", err)
	}
	expectedId := int32(1)
	expectedSection := "A"
	expectedSeatNum := int32(1)
	if r.Id != expectedId || r.Section != expectedSection || r.SeatNum != expectedSeatNum {
		t.Errorf("Purchase ticket failed, result %v, expected %v %v %v", r, expectedId, expectedSection, expectedSeatNum)
	}
}

func TestPurchaseTicketInvalidUser(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	_, err = c.PurchaseTicket(context.Background(), &pb.TicketPurchaseRequest{
		From: "London",
		To: "Paris", 
		UserId: 5,
		Price: 20,
	})
	if err == nil {
		t.Error("Expected Error")
	}
	e, ok := status.FromError(err)
	if !ok {
		t.Error("Expected error")
	}
	if e.Code() != codes.Internal {
		t.Errorf("Expected Internal, got %v", e.Code().String())
	}
}

func TestNoMoreTicketsLeft(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	dummySlice := make([]struct{}, 8)
	for range dummySlice {
		c.PurchaseTicket(context.Background(), &pb.TicketPurchaseRequest{
			From: "London",
			To: "Paris", 
			UserId: 2,
			Price: 20,
		})
	}
	_, err = c.PurchaseTicket(context.Background(), &pb.TicketPurchaseRequest{
		From: "London",
		To: "Paris", 
		UserId: 3,
		Price: 20,
	})
	_, err = c.PurchaseTicket(context.Background(), &pb.TicketPurchaseRequest{
		From: "London",
		To: "Paris", 
		UserId: 3,
		Price: 20,
	})
	if err == nil {
		t.Error("Expected Error")
	}
	e, ok := status.FromError(err)
	if !ok {
		t.Error("Expected error")
	}
	if e.Code() != codes.Internal {
		t.Errorf("Expected Internal, got %v", e.Code().String())
	}
}

func TestGetTicketDetails(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)

	r, err := c.GetTicketDetails(context.Background(), &pb.UserId{
		Id: 1,
	})
	if err != nil {
		t.Errorf("Get ticket got unexpected error %v", err)
	}
	expectedId := int32(1)
	if r.Id != expectedId {
		t.Errorf("Get ticket failed, result %v, expected %v", r, expectedId)
	}
}

func TestGetTicketInvalidUser(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	_, err = c.GetTicketDetails(context.Background(), &pb.UserId{
		Id: 5,
	})
	if err == nil {
		t.Error("Expected Error")
	}
	e, ok := status.FromError(err)
	if !ok {
		t.Error("Expected error")
	}
	if e.Code() != codes.Internal {
		t.Errorf("Expected Internal, got %v", e.Code().String())
	}
}

func TestGetTicketNoSeatsOfUser(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	_, err = c.GetTicketDetails(context.Background(), &pb.UserId{
		Id: 4,
	})
	if err == nil {
		t.Error("Expected Error")
	}
	e, ok := status.FromError(err)
	if !ok {
		t.Error("Expected error")
	}
	if e.Code() != codes.Internal {
		t.Errorf("Expected Internal, got %v", e.Code().String())
	}
}

func TestViewSeatAllocation(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)

	r, err := c.ViewSeatAllocation(context.Background(), &pb.Section{
		Section: "A",
	})
	if err != nil {
		t.Errorf("View seat allocation got unexpected error %v", err)
	}
	expectedLen := 5
	if len(r.Seats) != expectedLen {
		t.Errorf("View allocation failed, result %v, expected %v", len(r.Seats), expectedLen)
	}
}

func TestRemoveUser(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	_, err = c.RemoveUser(context.Background(), &pb.UserId{
		Id: 3,
	})
	if err != nil {
		t.Errorf("Remove user got unexpected error %v", err)
	}
}

func TestModifySeatSeatNotAvailable(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	_, err = c.ModifySeat(context.Background(), &pb.ModifySeatRequest{
		CurrTicketId: 1,
		NewSeatNum: 4,
		NewSection: "B",
		UserId: 1,
	})
	if err == nil {
		t.Error("Expected Error")
	}
	e, ok := status.FromError(err)
	if !ok {
		t.Error("Expected error")
	}
	if e.Code() != codes.Internal {
		t.Errorf("Expected Internal, got %v", e.Code().String())
	}
}

func TestModifySeat(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketBookingClient(conn)
	_, err = c.ModifySeat(context.Background(), &pb.ModifySeatRequest{
		CurrTicketId: 1,
		NewSeatNum: 5,
		NewSection: "B",
		UserId: 1,
	})
	if err != nil {
		t.Errorf("Modify got unexpected error %v", err)
	}
}

