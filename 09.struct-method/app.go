package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type User struct {
	Id        string
	FirstName string
	LastName  string
	CreatedAt time.Time
}

// method of the struct
func (u *User) dispalyUserDetails() {
	fmt.Println(u.Id, u.FirstName, u.LastName, u.CreatedAt)
}

// https://www.udemy.com/course/go-the-complete-guide/learn/lecture/40838178#notes
func (u *User) changeId() {
	u.Id = randomdata.Digits(5)
}

// constructor pattern
func New(userId, userFirstName, userLastName string) (*User, error) {

	// validation
	if userId == "" || userFirstName == "" || userLastName == "" {
		return nil, errors.New("userId, userFirstName, userLastName cannot be empty!")
	}

	user := User{
		Id:        userId,
		FirstName: userFirstName,
		LastName:  userLastName,
		CreatedAt: time.Now(),
	}

	return &user, nil
}

func main() {
	userId := randomdata.Digits(5)
	userFirstName := randomdata.FirstName(0)
	userLastName := randomdata.LastName()
	// user := User{
	// 	Id:        userId,
	// 	FirstName: userFirstName,
	// 	LastName:  userLastName,
	// 	CreatedAt: time.Now(),
	// }

	// var newUser *User = New(userId, userFirstName, userLastName)
	newUser, err := New(userId, userFirstName, userLastName)

	if err != nil {
		panic(err)
	}

	//calling method of the struct
	newUser.dispalyUserDetails()
	newUser.changeId()
	newUser.dispalyUserDetails()

}
