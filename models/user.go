package models

// Sent History
type Sent struct {
	To     string `json:"to" bson:"to"`
	Amount int    `json:"amount" bson:"amount"`
}

// Received History
type Received struct {
	From   string `json:"from" bson:"from"`
	Amount int    `json:"amount" bson:"amount"`
}

// User:
//		Name			: User's name
//		IsTransform	: Check that Are you in the other transfer? (To solve race condition)
//		Balance		: Check the balance (must greater than the amount) when you transfer
//		Sent			: Check the history sent to others user
//		Received		: Check the history received from others user
// Use sent and received history to validate the user's balance whenever you need
type User struct {
	Name       string     `json:"name" bson:"name"`
	IsTransfer bool       `json:"isTransform" bson:"isTransform"`
	Balance    int        `json:"balance" bson:"balance"`
	Sent       []Sent     `json:"sent" bson:"sent"`
	Received   []Received `json:"received" bson:"received"`
}
