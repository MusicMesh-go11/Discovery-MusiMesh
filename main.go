package main

import (
	"MusicMesh/Discovery-MusicMesh/generate/user_interactions"
	"MusicMesh/Discovery-MusicMesh/storage/postgres"
	"MusicMesh/Discovery-MusicMesh/storage/uInteraction"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func main() {
	userConn, err := grpc.NewClient(":5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer userConn.Close()

	CompositionConn, err := grpc.NewClient(":5052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer CompositionConn.Close()

	db, err := postgres.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	listner, err := net.Listen("tcp", ":5054")
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	grpcServer := grpc.NewServer()
	UI := uInteraction.NewUserInteraction(db, userConn, CompositionConn)
	user_interactions.RegisterUserInteractionsServiceServer(grpcServer, UI)
	log.Println("Listening on :5054")

	panic(grpcServer.Serve(listner))
}
