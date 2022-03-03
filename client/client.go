package main

import (
	"context"
	"fmt"
	protos "grpc/protos/currency"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	cs := protos.NewCurrencyClient(conn)

	resp, err := cs.GetRate(context.Background(), &protos.RateRequest{
		Base:        "GBP",
		Destination: "USD",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Rate)
}
