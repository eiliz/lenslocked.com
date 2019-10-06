package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"lenslocked.com/models"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "ella"
	dbname = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Email  string `gorm:"not null; unique_index"`
	Color  string
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	us, err := models.NewUserService(psqlInfo)

	if err != nil {
		panic(err)
	}

	defer us.Close()
	us.DestructiveReset()

	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com",
	}

	if err := us.Create(&user); err != nil {
		panic(err)
	}

	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}

	userByID, err := us.ByID(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(userByID)

}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error

	if err != nil {
		panic(err)
	}
}
