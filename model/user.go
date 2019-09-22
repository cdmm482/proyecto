package model

type userg struct {
	FirstName string `form:"firstname,omitempty" json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `form:"lastname,omitempty" json:"lastname,omitempty" bson:"lastname,omitempty"`
	CI        string `form:"ci,omitempty" json:"ci,omitempty" bson:"ci,omitempty"`
	Age       int    `form:"age,omitempty" json:"age,omitempty" bson:"age,omitempty"`
}
