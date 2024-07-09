package response

type Map map[string]any

// based on: https://github.com/omniti-labs/jsend
var (
	ErrInvalidProductID = Map{
		"status": "fail",
		"data": Map{
			"product": "invalid product ID",
		},
	}
	ErrNoProductNameProvided = Map{
		"status": "fail",
		"data": Map{
			"product": "name is required",
		},
	}
	ErrNotAuthorized = Map{
		"status": "fail",
	}
)
