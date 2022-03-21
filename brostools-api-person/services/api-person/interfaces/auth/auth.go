package auth

type ResponseData struct {
	Errors []*ResponseErrors `json:"errors"`
	Data   []string          `json:"data"`
}

type ResponseErrors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
