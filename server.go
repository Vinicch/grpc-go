package main

import (
	protos "grpc/protos/currency"
	"grpc/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	cs := server.NewCurrency(log)
	gs := grpc.NewServer()

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(lis)
}
