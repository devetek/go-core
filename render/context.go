package render

import (
	"context"
)

func Context(ctx context.Context) *Page {
	page := ctx.Value("renderer").(*Page)

	return page
}
