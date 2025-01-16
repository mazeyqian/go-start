package agent

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Response struct {
	StatusCode int          `json:"status_code"`
	Data       ResponseData `json:"data"`
}

type RecordRequestOrResponse struct {
	MethodOrStatusCode string `json:"method"`
	URL                string `json:"url"`
	Data               string `json:"data"`
}
