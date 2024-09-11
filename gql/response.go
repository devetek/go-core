package gql

type graphResponse struct {
	Data   interface{} `json:"data"`
	Errors []graphErr
}

type graphErr struct {
	Message string
}

func (e graphErr) Error() string {
	return "graphql: " + e.Message
}
