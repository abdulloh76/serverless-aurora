package store

import (
	"errors"

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

	return &PostgresDBStore{
		db,
	}
}

func (d *PostgresDBStore) All() ([]types.User, error) {
	// todo
	// 	if err != nil {
	// 		return productRange, fmt.Errorf("failed to get items from db: %w", err)
	// 	}
	// if err != nil {
	// 		return productRange, fmt.Errorf("failed to unmarshal data from db: %w", err)
	// 	}

	return []types.User{}, nil
}

func (d *PostgresDBStore) Get(id string) (*types.User, error) {
	// todo
	// ErrUserNotFound
	return &types.User{}, nil
}

func (d *PostgresDBStore) Create(user types.CreateUser) (*types.User, error) {
	// marshall, create
	return nil, nil
}

func (d *PostgresDBStore) Modify(id string, user types.CreateUser) (*types.User, error) {
	// todo marshall, update
	// ErrUserNotFound
	// if err != nil {
	// 	return fmt.Errorf("unable to marshal product: %w", err)
	// }
	// if err != nil {
	// 	return fmt.Errorf("cannot put item: %w", err)
	// }

	return nil, nil
}

func (d *PostgresDBStore) Delete(id string) error {
	// todo
	// ErrUserNotFound
	// if err != nil {
	// 	return fmt.Errorf("can't delete item: %w", err)
	// }

	return nil
}

// The project is a test project to explore using Go for Lambdas and hopefully find a developer
// to convert over 100 Lambdas written in Javascript and TypeScript and rewrite them in GO.
