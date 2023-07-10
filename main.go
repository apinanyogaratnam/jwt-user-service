package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/apinanyogaratnam/jwt-user-service/jwt-protobuf/jwt"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("host.docker.internal:9000", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalln("Could not connect to gRPC server: ", err);
	}

	defer conn.Close()

	c := jwt.NewJWTServiceClient(conn)

	getTokenRequestMessage := jwt.GetTokenRequest{
		Id: 1,
	}

	var getTokenResponse *jwt.GetTokenResponse;

	getTokenResponse, err = c.GetToken(context.Background(), &getTokenRequestMessage)
	if err != nil {
		log.Fatalln("Error when calling GetToken: ", err)
	}

	log.Println("Received Token:", getTokenResponse.Token);

	validateTokenRequestMessage := jwt.ValidateTokenRequest{
		Token: getTokenResponse.Token,
	}

	var validateTokenResponse *jwt.ValidateTokenResponse;

	validateTokenResponse, err = c.ValidateToken(context.Background(), &validateTokenRequestMessage)
	if err != nil {
		log.Fatalln("Error when calling ValidateToken: ", err)
	}

	log.Println("Token validated:", validateTokenResponse.Valid);
}
