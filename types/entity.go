package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string
	Firstname string `gorm:"varchar(100)"`
	Lastname  string `gorm:"varchar(100)"`
}

// func (invoice *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	invoice.ID, err = shortid.Generate()

// 	if err != nil {
// 		return errors.New("can't save invalid data")
// 	}
// 	return nil
// }
