package helper

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
