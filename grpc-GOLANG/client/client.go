package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	user "grpc/proto/gen"
	"log"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8200", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	usr := user.NewUserServiceClient(conn)

	fmt.Print("DIGITE O USERNAME DO SEU GITHUB: ")
	var inputUser string
	fmt.Scanln(&inputUser)

	response, err := usr.GetUser(context.Background(), &user.UserRequest{Username: inputUser})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}

	log.Printf("RESPOSTA DO SERVIDOR: %v", response)
}
