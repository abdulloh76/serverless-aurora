package types

type Store interface {
	All() ([]User, error)
	Get(id string) (*User, error)
	Create(user CreateUser) error
	Modify(id string, user CreateUser) (*User, error)
	Delete(id string) error
}
