package gqlmodel

import (
	"fmt"
	"io"
	"io/fs"
)

type Graphql struct {
	fs fs.FS
}

func New(fs fs.FS) *Graphql {
	e := &Graphql{
		fs: fs,
	}

	return e
}

func (p *Graphql) Query(name string) (string, error) {
	px, err := p.fs.Open(name)
	if err != nil {
		return "", fmt.Errorf("[models.Query] - error on p.fs.Open: %w", err)
	}

	query, err := io.ReadAll(px)
	if err != nil {
		return "", fmt.Errorf("[models.Query] - error on io.ReadAll: %w", err)
	}

	return string(query), err
}
