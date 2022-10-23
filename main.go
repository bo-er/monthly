package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/bo-er/monthly/proto/company"
	servers "github.com/bo-er/monthly/servers"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterDepartmentServiceServer(s, &servers.Server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:10001")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:10001",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	ctx := context.Background()
	err = pb.RegisterDepartmentServiceHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register department service gateway:", err.Error())
	}
	pb.RegisterEmployeeServiceHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register employee service gateway: ", err.Error())
	}

	pb.RegisterPetServiceHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register pet service gateway: ", err.Error())
	}

	gwServer := &http.Server{
		Addr:    ":10002",
		Handler: gwmux,
	}

	go func() {
		fs := http.FileServer(http.Dir("./swagger-ui"))
		http.Handle("/swagger", http.StripPrefix("/swagger", fs))
	}()
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:10002")
	log.Fatalln(gwServer.ListenAndServe())
}
