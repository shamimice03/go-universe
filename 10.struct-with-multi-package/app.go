package main

import (
	"cloudterms.net/struct-package/order"
	"cloudterms.net/struct-package/user"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	userId := randomdata.Digits(5)
	userFirstName := randomdata.FirstName(0)
	userLastName := randomdata.LastName()

	// var newUser *User = New(userId, userFirstName, userLastName)
	newUser, err := user.New(userId, userFirstName, userLastName)

	if err != nil {
		panic(err)
	}

	//calling method of the struct
	newUser.DisplayUserDetails()
	newUser.ChangeId()
	newUser.DisplayUserDetails()

	orderId := randomdata.Digits(2)
	orderAmount := randomdata.Decimal(22)

	//println(orderId, orderAmount)
	newOrder, _ := order.PlaceNewOrder(orderId, orderAmount, newUser)

	newOrder.DisplayOrder()

}
