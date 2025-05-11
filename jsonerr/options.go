package jsonerr

// Payload is a struct to store message and error
type NewError struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

// Options is a function type to set options
type Options func(err *instance)

// WithCaller set caller for error
func WithCaller(enable bool) Options {
	return func(err *instance) {
		err.caller = enable
	}
}

// WithAlwaysPrint set always print for error when invoke instance.Unwrap()
func WithAlwaysPrint(enable bool) Options {
	return func(err *instance) {
		err.alwaysPrint = enable
	}
}
