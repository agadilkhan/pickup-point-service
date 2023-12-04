package auth

type GenerateTokenRequest struct {
	Login    string
	Password string
}

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

type CreateUserRequest struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Login     string
	Password  string
}

type ConfirmUserRequest struct {
	Email string
	Code  string
}

type UpdateUserRequest struct {
	ID        int
	RoleID    int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Login     string
	Password  string
}
