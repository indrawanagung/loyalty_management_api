package web

type WebResponse struct {
	Header Header      `json:"header"`
	Data   interface{} `json:"data"`
}
type Header struct {
	Message interface{} `json:"message"`
	Error   bool        `json:"error"`
}
