package gqlmodel

import (
	"context"
)

func Context(ctx context.Context) *Graphql {
	c := ctx.Value("model").(*Graphql)

	return c
}
