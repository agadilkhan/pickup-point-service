package entity

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	Login       string
	Password    string
	IsConfirmed bool
}
