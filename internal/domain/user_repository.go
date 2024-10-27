package domain

type UserRepository interface {
	GetUser(username string) (*User, error)
}
