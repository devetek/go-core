## Description

Scan file .graphql and used it as graphql schema


## How To Use

```golang
package internal

import (
	"embed"
	"os"
	"log"

	"github.com/devetek/go-core/mdfs"
	"github.com/devetek/go-core/gql"
)

var (
	// Update this embed when create .graphql in different structure

    //go:embed models/**/*.graphql
	mdl embed.FS

	Model = mdfs.New(mdl, "", "")
)

// init all graphql model from folder pattern models/**/*.graphql
gqlModel := gql.NewModel(Model)

// init gql HTTP client with host
gqlClient:= gql.NewClient("https://gql.terpusat.com")

// use specific gql model
userQuery, err := gql.Query("models/common/user.graphql")
if err != nil {
	log.Println(err)
}

// crete new request
reqGql := gqlclient.NewRequest(userQuery)

// then run it and capture the response
if err := gql.Run(ctx, reqGql, &generalSetting); err != nil {
	log.Println(err)
}
```

