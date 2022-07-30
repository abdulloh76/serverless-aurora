package domain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/abdulloh76/serverless-aurora/types"
)

var (
	ErrJsonUnmarshal = errors.New("failed to parse user from request body")
)

type Users struct {
	store types.Store
}

func NewUsersDomain(s types.Store) *Users {
	return &Users{
		store: s,
	}
}

func (u *Users) GetUser(id string) (*types.User, error) {
	product, err := u.store.Get(id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return product, nil
}

func (u *Users) AllUsers() ([]types.User, error) {
	allUsers, err := u.store.All()
	if err != nil {
		return allUsers, fmt.Errorf("%w", err)
	}

	return allUsers, nil
}

func (u *Users) Create(body []byte) (*types.User, error) {
	user := types.CreateUser{}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("%w", ErrJsonUnmarshal)
	}

	newUser, err := u.store.Create(user)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return newUser, nil
}

func (u *Users) ModifyUser(id string, body []byte) (*types.User, error) {
	user := types.CreateUser{}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("%w", ErrJsonUnmarshal)
	}

	updatedUser, err := u.store.Modify(id, user)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return updatedUser, nil
}

func (u *Users) DeleteUser(id string) error {
	err := u.store.Delete(id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
