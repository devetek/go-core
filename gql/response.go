package gql

type gqlResponse struct {
	Data   any `json:"data"`
	Errors []gqlError
}

type gqlError struct {
	Message string
}

func (e gqlError) Error() string {
	return "graphql: " + e.Message
}
