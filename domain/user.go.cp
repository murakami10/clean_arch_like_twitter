package domain

type User struct {
	ID        uint
	FirstName string
	LastName  string
	AvatarURL string
}

type Users []User