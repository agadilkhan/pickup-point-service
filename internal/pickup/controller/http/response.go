package http

// swagger:model ResponseOK
type responseOK struct {
	// example: any
	Data interface{} `json:"data"`
}

// swagger:model ResponseMessage
type responseMessage struct {
	Message string `json:"message"`
}

// swagger:model ResponseCreated
type responseCreated struct {
	ID int `json:"id"`
}
