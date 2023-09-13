package main

import (
	"context"
	"log"
	"time"

	pb "github.com/withthemoonrabbit/lagoserv/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "34.64.50.248:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateNewUser(ctx, &pb.NewUser{
		Nickname: "a",
		Email:    "b",
		Pwhash:   "c",
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf(`User Details:
NicknameME: %s
Email: %s
PasswordHash: %s`, r.GetNickname(), r.GetEmail(), r.GetPwhash())
}
