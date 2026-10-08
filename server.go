package main

import (
	"grpc-advcalc/advcalc"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("unable to listen to 8081. Port already in use")
	}

	grpcServer := grpc.NewServer()

	aServer := advcalc.AdvancedCalcServerLive{}
	advcalc.RegisterAdvancedCalcServer(grpcServer, &aServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
