package helper

func CodeToMessage(code int) string {
	var message string

	switch code {
	case 401:
		message = "Unauthorized"
	case 403:
		message = "Forbidden request"
	case 404:
		message = "Resource not found"
	default:
		message = "Internal server error"
	}

	return message
}