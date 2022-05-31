package models

type Sent struct {
	// From   string `json:"from" bson:"from"`
	To     string `json:"to" bson:"to"`
	Amount int    `json:"amount" bson:"amount"`
}

type Received struct {
	// From   string `json:"from" bson:"from"`
	From   string `json:"from" bson:"from"`
	Amount int    `json:"amount" bson:"amount"`
}

type User struct {
	Name        string     `json:"name" bson:"name"`
	IsTransform bool       `json:"isTransform" bson:"isTransform"`
	Balance     int        `json:"balance" bson:"balance"`
	Sent        []Sent     `json:"sent" bson:"sent"`
	Received    []Received `json:"received" bson:"received"`
}
