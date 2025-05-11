package render

import (
	"context"
)

func Context(ctx context.Context) *Page {
	page := ctx.Value(CTX_KEY).(*Page)

	return page
}
