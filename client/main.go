package main

import (
	"flag"
	"golang-grpc-examples/client/api"
	"golang-grpc-examples/messages/messenger"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

var (
	servicePort string
	apiPort     string
)

func init() {
	flag.StringVar(&servicePort, "sp", ":5000", "set service port")
	flag.StringVar(&apiPort, "p", ":8080", "set api port")
}

func main() {
	flag.Parse()

	log.Printf("Api listening on port %v\n", apiPort)
	log.Printf("Target service port %v\n", servicePort)

	conn, err := grpc.Dial(servicePort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	messengerClient := messenger.NewMessengerServiceClient(conn)

	api := api.NewMessengerAPI(messengerClient)

	http.HandleFunc("/messages", api.Handlers)

	log.Fatal(http.ListenAndServe(apiPort, nil))
}
