package http

type responseOK struct {
	Data interface{} `json:"data"`
}

type responseMessage struct {
	Message string `json:"message"`
}

type responseCreated struct {
	ID int `json:"id"`
}
