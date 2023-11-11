package dto

type CreateUserRequest struct {
	RoleID    int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Login     string
	Password  string
}
