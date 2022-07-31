package store

import (
	"errors"
	"fmt"

	"github.com/abdulloh76/serverless-aurora/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user with given ID not found")
)

type PostgresDBStore struct {
	db *gorm.DB
}

var _ types.Store = (*PostgresDBStore)(nil)

func NewPostgresDBStore(dsn string) *PostgresDBStore {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&types.User{})

	return &PostgresDBStore{
		db,
	}
}

func (d *PostgresDBStore) All() ([]types.User, error) {
	var users []types.User
	err := d.db.Model(&types.User{}).Find(&users).Error

	return users, err
}

func (d *PostgresDBStore) Get(id string) (*types.User, error) {
	var user types.User
	err := d.db.First(&user, "id = ?", id).Error

	if user == (types.User{}) {
		return nil, fmt.Errorf("%w", ErrUserNotFound)
	}

	return &user, err
}

func (d *PostgresDBStore) Create(user *types.User) error {
	err := d.db.Create(&user).Error

	return err
}

func (d *PostgresDBStore) Modify(id string, userDto types.CreateUser) (*types.User, error) {
	var user types.User
	err := d.db.First(&user, "id = ?", id).Error

	if user == (types.User{}) {
		return nil, fmt.Errorf("%w", ErrUserNotFound)
	}
	if err != nil {
		return nil, err
	}

	user.Firstname = userDto.Firstname
	user.Lastname = userDto.Lastname
	err = d.db.Save(&user).Error

	return &user, err
}

func (d *PostgresDBStore) Delete(id string) error {
	var user types.User
	err := d.db.First(&user, "id = ?", id).Error

	if user == (types.User{}) {
		return fmt.Errorf("%w", ErrUserNotFound)
	}
	if err != nil {
		return err
	}

	err = d.db.Delete(&user).Error

	return err
}
