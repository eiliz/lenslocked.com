package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is returned when a resource was not found in the db
	ErrNotFound = errors.New("models: resource not found")

	// ErrInvalidID is returned when an invalid ID is provided
	// to a method like Delete
	ErrInvalidID = errors.New("models: ID provided was invalid")
)

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &UserService{
		db: db,
	}, nil
}

// Close closes the db connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// DestructiveReset drops the users and rebuilds it
func (us *UserService) DestructiveReset() error {
	if err := us.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return us.AutoMigrate()
}

// AutoMigrate will attempt to atuomatically migrate the Users table
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

// Create will create the provided user and backfill data
// like ID, CreateAt, and UpdatedAt
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// Update will updated the provided user with all of data
// in the provided user object
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

type UserService struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null; unique_index"`
}

// ByID will lookup a user a given ID
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// ByEmail looks up a user with the given email address
// and returns that user
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// Delete will remove the user with the provided ID
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}

	user := User{
		Model: gorm.Model{ID: id},
	}

	err := us.db.Delete(&user).Error
	return err
}

// first will query using the provided gorm.DB and it will get
// the first item returned and place it into dst.
// If nothing is found in the query, it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error

	if err == gorm.ErrRecordNotFound {
		return ErrNotFound

	}
	return err
}
