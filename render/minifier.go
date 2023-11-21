package render

import (
	"github.com/tdewolff/minify/v2"
	htmlminify "github.com/tdewolff/minify/v2/html"
	svgminify "github.com/tdewolff/minify/v2/svg"
)

var m = minify.New()

func minifyHtml(html string) (string, error) {
	m.Add("text/html", &htmlminify.Minifier{
		KeepWhitespace:      false,
		KeepDefaultAttrVals: false,
		KeepDocumentTags:    false,
		KeepEndTags:         false,
		KeepQuotes:          false,
	})
	m.Add("image/svg+xml", &svgminify.Minifier{
		KeepComments: false,
	})

	shtml, err := m.String("text/html", html)
	if err != nil {
		return "", err
	}

	// ssvg, err := m.String("image/svg+xml", shtml)
	// if err != nil {
	// 	return "", err
	// }

	return shtml, nil
}
