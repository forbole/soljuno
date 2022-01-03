package types

type GraphQLError struct {
	Message string `json:"message"`
}

func NewError(err error) GraphQLError {
	return GraphQLError{
		Message: err.Error(),
	}
}
