package internal

type RequestData struct {
	QueryParameters map[string][]string
	Body            interface{}
}

type RequestDataOption = func(*RequestData)

func WithQuery(queryParameters map[string][]string) RequestDataOption {
	return func(p *RequestData) {
		p.QueryParameters = queryParameters
	}
}

func WithBody(body interface{}) RequestDataOption {
	return func(p *RequestData) {
		p.Body = body
	}
}
