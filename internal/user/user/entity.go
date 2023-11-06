package user

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}
