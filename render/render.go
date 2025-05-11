package render

import (
	"html/template"
	"io/fs"
	"net/http"
	"sync"

	"github.com/gobuffalo/plush"
	"github.com/tdewolff/minify/v2"
	htmlminify "github.com/tdewolff/minify/v2/html"
	svgminify "github.com/tdewolff/minify/v2/svg"
)

type key string

const CTX_KEY key = "renderer"

type Engine struct {
	templates     fs.FS
	defaultLayout string

	moot     sync.Mutex
	helpers  template.FuncMap
	values   map[string]any
	minifier MinifierOption
}

func NewEngine(fs fs.FS, options ...Option) *Engine {
	e := &Engine{
		templates: fs,

		values:  make(map[string]any),
		helpers: make(template.FuncMap),

		defaultLayout: "components/container/index.html",
	}

	for _, option := range options {
		option(e)
	}

	return e
}

func (e *Engine) Set(key string, value any) {
	e.moot.Lock()
	defer e.moot.Unlock()

	e.values[key] = value
}

func (e *Engine) SetHelper(key string, value any) {
	e.moot.Lock()
	defer e.moot.Unlock()

	e.helpers[key] = value
}

func (e *Engine) HTML(w http.ResponseWriter) *Page {
	p := &Page{
		fs:            e.templates,
		writer:        w,
		minify:        minify.New(),
		minifyEnable:  e.minifier.Enable,
		defaultLayout: e.defaultLayout,
	}

	if e.minifier.Enable {
		// minifier config
		p.minify.Add("text/html", &htmlminify.Minifier{
			KeepWhitespace:      e.minifier.HTML.KeepWhitespace,
			KeepDefaultAttrVals: e.minifier.HTML.KeepDefaultAttrVals,
			KeepDocumentTags:    e.minifier.HTML.KeepDocumentTags,
			KeepEndTags:         e.minifier.HTML.KeepEndTags,
			KeepQuotes:          e.minifier.HTML.KeepQuotes,
		})
		p.minify.Add("image/svg+xml", &svgminify.Minifier{
			KeepComments: false,
		})
	}

	ctx := plush.NewContext()
	for k, v := range e.values {
		ctx.Set(k, v)
	}

	// common helper
	ctx.Set("dangerouslySetInnerHTML", func(strHtml string) template.HTML {

		rawHTML := template.HTML(strHtml)

		return rawHTML
	})

	for k, v := range e.helpers {
		ctx.Set(k, v)
	}

	ctx.Set("partialFeeder", func(name string) (string, error) {
		return p.open(name)
	})

	p.context = ctx

	return p
}
