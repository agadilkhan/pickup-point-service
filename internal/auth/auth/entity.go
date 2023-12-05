package auth

//swagger:model GenerateToken
type GenerateTokenRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//swagger:model UserToken
type JWTUserToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTTokenContent struct {
	UserID    string
	Phone     string
	FirstName string
	LastName  string
}

//swagger:model RegisterUser
type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type ConfirmUserRequest struct {
	Email string
	Code  string
}

//swagger:model UpdateUser
type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

// swagger:model UserCode
type UserCodeRequest struct {
	Code string `json:"code"`
}

//swagger:model RefreshToken
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
