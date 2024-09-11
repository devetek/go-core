package main

import (
	"context"
	"log"
	"os"

	"github.com/devetek/go-core/gql"
)

type Hello struct {
	Data struct {
		Name string `json:"name"`
	} `json:"data"`
}

func main() {
	var hello Hello

	schema := os.DirFS("schema")

	gqlSchema := gql.NewSchema(schema)
	gqlHttpClient := gql.NewHttpClient(
		"https://gql.terpusat.com",
		gql.Debug(),
		gql.ImmediatelyCloseReqBody(),
	)
	gqlHttpClient.Log = func(s string) {
		log.Println(s)
	}

	useQuery, err := gqlSchema.Query("hello.graphql")
	if err != nil {
		log.Panicln(err)
	}

	gqlCall := gql.NewRequest(useQuery)

	err = gqlHttpClient.Run(context.Background(), gqlCall, &hello)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Gql Response:", hello)
}
