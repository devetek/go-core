package render

import "html/template"

type Option func(*Engine)

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
