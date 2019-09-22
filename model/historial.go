package model

import "time"

type history struct {
	Index    int       `form:"index,omitempty" json:"index,omitempty" bson:"index,omitempty"`
	Patien   patient   `form:"patien,omitempty" json:"patien,omitempty" bson:"patien,omitempty"`
	Medi     medic     `form:"medi,omitempty" json:"medi,omitempty" bson:"medi,omitempty"`
	QueryAt  time.Time `form:"queryAt,omitempty" json:"queryAt,omitempty" bson:"queryAt,omitempty"`
	Hash     string    `form:"hash,omitempty" json:"hash,omitempty" bson:"hash,omitempty"`
	PrevHash string    `form:"prevHash,omitempty" json:"prevHash,omitempty" bson:"prevHash,omitempty"`
}
