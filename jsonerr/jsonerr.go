package jsonerr

// JSONError is an interface to create custom error, print error, and return error message
type Error interface {
	Create(err NewError)
	Print()
	Error() string
	Unwrap() error
}
