package admin

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
