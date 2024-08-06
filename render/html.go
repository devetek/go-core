package render

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gobuffalo/plush"
	"github.com/tdewolff/minify/v2"
)

type Page struct {
	context      *plush.Context
	writer       http.ResponseWriter
	fs           fs.FS
	minify       *minify.M
	minifyEnable bool

	defaultLayout string
}

// common context, you can set any data type
func (p *Page) Set(key string, value any) {
	p.context.Set(key, value)
}

func (p *Page) Value(key string) any {
	return p.context.Value(key)
}

func (p *Page) Render(page string) error {
	// find the template from the fs
	html, err := p.open(page)
	if err != nil {
		return fmt.Errorf("[render.render] - error on find the template from the fs: %w", err)
	}

	html, err = plush.Render(html, p.context)
	if err != nil {
		return err
	}

	layout, err := p.open(p.defaultLayout)
	if err != nil {
		return fmt.Errorf("[render.render] - error on find default layout from the fs: %w", err)
	}

	layout = strings.Replace(layout, "<%= Tyield %>", html, 1)
	html, err = plush.Render(layout, p.context)
	if err != nil {
		return err
	}

	if p.minifyEnable {
		shtml, err := p.minify.String("text/html", html)
		if err != nil {
			return fmt.Errorf("failed to minify html: %w", err)
		}

		_, err = p.writer.Write([]byte(shtml))
		if err != nil {
			return fmt.Errorf("could not write to response minify: %w", err)
		}
	}

	_, err = p.writer.Write([]byte(html))
	if err != nil {
		return fmt.Errorf("could not write to response non minify: %w", err)
	}

	return nil
}

func (p *Page) RenderWithLayout(page, layout string) error {
	html, err := p.open(page)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	html, err = plush.Render(html, p.context)
	if err != nil {
		return err
	}

	layout, err = p.open(layout)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	layout = strings.Replace(layout, "<%= Tyield %>", html, 1)
	html, err = plush.Render(layout, p.context)
	if err != nil {
		return err
	}

	if p.minifyEnable {
		shtml, err := p.minify.String("text/html", html)
		if err != nil {
			return fmt.Errorf("failed to minify html: %w", err)
		}

		_, err = p.writer.Write([]byte(shtml))
		if err != nil {
			return fmt.Errorf("could not write to response minify: %w", err)
		}
	}

	_, err = p.writer.Write([]byte(html))
	if err != nil {
		return fmt.Errorf("could not write to response non minify: %w", err)
	}

	return nil
}

func (p *Page) RenderClean(name string) error {
	// find the template from the fs
	html, err := p.open(name)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	html, err = plush.Render(html, p.context)
	if err != nil {
		return err
	}

	if p.minifyEnable {
		shtml, err := p.minify.String("text/html", html)
		if err != nil {
			return fmt.Errorf("failed to minify html: %w", err)
		}

		_, err = p.writer.Write([]byte(shtml))
		if err != nil {
			return fmt.Errorf("could not write to response minify: %w", err)
		}
	}

	_, err = p.writer.Write([]byte(html))
	if err != nil {
		return fmt.Errorf("could not write to response non minify: %w", err)
	}

	return nil
}

func (p *Page) open(name string) (string, error) {
	px, err := p.fs.Open(name)
	if err != nil {
		return "", fmt.Errorf("[render.open] - error on p.fs.Open: %w", err)
	}

	html, err := io.ReadAll(px)
	if err != nil {
		return "", fmt.Errorf("[render.open] - error on io.ReadAll: %w", err)
	}

	return string(html), err
}
