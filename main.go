package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/apinanyogaratnam/jwt-user-service/jwt"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalln("Could not connect to gRPC server: ", err);
	}

	defer conn.Close()

	c := jwt.NewJWTServiceClient(conn)

	message := jwt.JWTRequest{
		Id: 1,
	}

	var response *jwt.JWTResponse;

	response, err = c.GetToken(context.Background(), &message)
	if err != nil {
		log.Fatalln("Error when calling GetToken: ", err)
	}

	log.Printf("Response from server: %s", response.Token);
}
