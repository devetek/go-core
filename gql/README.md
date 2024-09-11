## Description

Golang graphql client implementation, to simplify interact with graphql query using file .graphql. Check folder example for detail.

## Getting Started

### Installation
To start using gql, install Go and run go get:
```sh
go get -u github.com/devetek/go-core
```

### Basic
Import gql into your application to access its gql capabilities
```sh
"github.com/devetek/go-core/gql"
```

### Usage
```sh
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

	err = gqlHttpClient.Run(context.Background(), gqlCall, hello)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Gql Response:", hello)
}
```