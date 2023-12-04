package swagger_model

// swagger:model ResponseOK
type responseOK struct {
	// example: any
	Data interface{} `json:"data"`
}

// swagger:model ResponseMessage
type response struct {
	Message string `json:"message"`
}

// swagger:model ResponseCreated
type responseCreated struct {
	ID int `json:"id"`
}

// swagger:model ConfirmUserRequest
type ConfirmUserRequest struct {
	Code string `json:"code"`
}

//swagger:model RegisterUser
type CreateUserRequest struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Login     string
	Password  string
}

//swagger:model RefreshToken
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

//swagger:model UserToken
type JWTUserToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

//swagger:model GenerateToken
type GenerateTokenRequest struct {
	Login    string
	Password string
}

//swagger:model UpdateUserRequest
type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}
