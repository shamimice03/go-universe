package main

import (
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

func (u User) New() {

}

func main() {
	userId := randomdata.Digits(5)
	userFirstName := randomdata.FirstName(0)
	userLastName := randomdata.LastName()
	user := User{
		Id:        userId,
		FirstName: userFirstName,
		LastName:  userLastName,
		CreatedAt: time.Now(),
	}

	dispalyUserDetails(&user)

	fmt.Println(user)
	fmt.Println(user.CreatedAt.Date())
}

/* //without struct
func dispalyUserDetails(userId, userFirstName, userLastName string) {
	fmt.Println(userId, userFirstName, userLastName)
}
*/

// with struct
func dispalyUserDetails(u *User) {
	fmt.Println((*u).Id, u.FirstName, (*u).LastName, u.CreatedAt)
}
