package github

// ErrorResponse describes the error response from the Github API
type ErrorResponse struct {
	Message          string  `json:"message"`
	Errors           []error `json:"errors"`
	DocumentationURL string  `json:"documentation_url"`
}

type error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
