package schemas

type SchemaError struct {
	Error   error
	Code    uint
	Message string
}

type ValidationEror struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
