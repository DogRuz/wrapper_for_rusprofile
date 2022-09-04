package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"wrapper_for_rusprofile/internal/api"
)

func main() {
	flag.Parse()
	inn := flag.Arg(0)
	conn, err := grpc.Dial(":8070", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewGetInformationClient(conn)

	res, err := c.GetByInn(context.Background(), &api.GetInformationRequest{Inn: inn})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
