package render

import "html/template"

type Option func(*Engine)

type MinifierHTMLConfig struct {
	KeepComments            bool
	KeepConditionalComments bool
	KeepDefaultAttrVals     bool
	KeepDocumentTags        bool
	KeepEndTags             bool
	KeepQuotes              bool
	KeepWhitespace          bool
}
type MinifierOption struct {
	Enable bool
	HTML   MinifierHTMLConfig
}

// WithDefaultLayout sets the default layout for the engine
// Default to components/index.html
func WithDefaultLayout(layout string) Option {
	return func(e *Engine) {
		e.defaultLayout = layout
	}
}

// WithHelpers sets the helpers for the engine these helpers will be
// Available globally. eg: 	func(param string) string { return param }
func WithHelpers(hps template.FuncMap) Option {
	return func(e *Engine) {
		e.helpers = hps
	}
}

// WithValues sets key and value to plush context
// Available globally. e.g: map[string]any{ "key": "value" })
func WithValues(values map[string]any) Option {
	return func(e *Engine) {
		e.values = values
	}
}

// WithMinifier sets minifier configuration
// Available globally. e.g: map[string]any{ "key": "value" })
func WithMinifier(values MinifierOption) Option {
	return func(e *Engine) {
		e.minifier.Enable = values.Enable
		e.minifier.HTML.KeepComments = values.HTML.KeepComments
		e.minifier.HTML.KeepConditionalComments = values.HTML.KeepConditionalComments
		e.minifier.HTML.KeepDefaultAttrVals = values.HTML.KeepDefaultAttrVals
		e.minifier.HTML.KeepDocumentTags = values.HTML.KeepDocumentTags
		e.minifier.HTML.KeepEndTags = values.HTML.KeepEndTags
		e.minifier.HTML.KeepQuotes = values.HTML.KeepQuotes
		e.minifier.HTML.KeepWhitespace = values.HTML.KeepWhitespace
	}
}
