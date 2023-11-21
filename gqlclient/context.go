package gqlclient

import (
	"context"
)

func Context(ctx context.Context) *Client {
	c := ctx.Value("gql").(*Client)

	return c
}
