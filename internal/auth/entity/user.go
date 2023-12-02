package entity

type User struct {
	ID          int    `json:"id"`
	RoleID      int    `json:"role_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	IsConfirmed bool   `json:"is_confirmed"`
}
