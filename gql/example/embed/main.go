package main

import (
	"context"
	"embed"
	"log"

	"github.com/devetek/go-core/gql"
	"github.com/devetek/go-core/mdfs"
)

type Extensions struct {
	Debug []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"debug"`
}

type GeneralSetting struct {
	Data struct {
		GeneralSettings struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
		} `json:"generalSettings"`
	} `json:"data"`
	Extensions Extensions `json:"extensions"`
}

var (
	endpoint = "https://internal.okcarlomboktransport.com/graphql/"

	// Update embed location if files in the different structure
	//go:embed schema/*.graphql
	mdl embed.FS

	schema = mdfs.New(mdl, "", "")
)

func main() {
	var generalSetting GeneralSetting

	gqlSchema := gql.NewSchema(schema)
	gqlHttpClient := gql.NewHttpClient(
		endpoint,
		gql.Debug(),
		gql.ImmediatelyCloseReqBody(),
	)
	gqlHttpClient.Log = func(s string) {
		log.Println(s)
	}

	useQuery, err := gqlSchema.Query("schema/wp_general_setting.graphql")
	if err != nil {
		log.Panicln(err)
	}

	gqlCall := gql.NewRequest(useQuery)

	err = gqlHttpClient.Run(context.Background(), gqlCall, generalSetting)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Gql Response:", generalSetting)
}
