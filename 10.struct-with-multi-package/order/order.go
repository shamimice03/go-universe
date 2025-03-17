package order

import (
	"errors"
	"fmt"

	"cloudterms.net/struct-package/user"
)

type Order struct {
	orderId     string
	orderAmount float64
	userDetails user.User
}

// method
func (o *Order) DisplayOrder() {
	fmt.Println(o)
	fmt.Println(o.orderId, o.orderAmount, o.userDetails)

	// can access method of User package
	o.userDetails.DisplayUserDetails()
}

// constructor pattern
func PlaceNewOrder(orderId string, orderAmount float64, userDetails *user.User) (*Order, error) {
	if orderId == "" || orderAmount == 0.0 {
		return nil, errors.New("orderId and orderAmount is required")
	}

	newOrder := Order{
		orderId:     orderId,
		orderAmount: orderAmount,
		userDetails: *userDetails,
	}

	return &newOrder, nil
}
