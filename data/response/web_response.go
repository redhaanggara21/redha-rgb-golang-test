package response

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Token  string      `json:"token"`
	Data   interface{} `json:"data,omitempty"`
}
