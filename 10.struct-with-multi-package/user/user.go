package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type User struct {
	id        string
	firstName string
	lastName  string
	createdAt time.Time
}

// method of the struct
func (u *User) DisplayUserDetails() {
	fmt.Println(u.id, u.firstName, u.lastName, u.createdAt)
}

func (u *User) ChangeId() {
	u.id = randomdata.Digits(5)
}

// constructor pattern
func New(userId, userFirstName, userLastName string) (*User, error) {

	// validation
	if userId == "" || userFirstName == "" || userLastName == "" {
		return nil, errors.New("userId, userFirstName, userLastName cannot be empty")
	}

	user := User{
		id:        userId,
		firstName: userFirstName,
		lastName:  userLastName,
		createdAt: time.Now(),
	}

	return &user, nil
}
