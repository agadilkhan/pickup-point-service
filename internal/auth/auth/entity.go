package auth

type GenerateTokenRequest struct {
	Login    string
	Password string
}

type JWTUserToken struct {
	Token        string
	RefreshToken string
}

type JWTTokenContent struct {
	UserID    string
	Phone     string
	FirstName string
	LastName  string
}

type CreateUserRequest struct {
	Login    string
	Password string
}
