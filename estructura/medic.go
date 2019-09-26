package estructura

import "time"
// Medic es una estructura del medico que contiene una variable del tipo usuario
type Medic struct {
	Index    int       `form:"index,omitempty" json:"index,omitempty" bson:"index,omitempty"`
	Code     string    `form:"code,omitempty" json:"code,omitempty" bson:"code,omitempty"`
	User     Userg     `form:"user,omitempty" json:"user,omitempty" bson:"user,omitempty"`
	JoinedAt time.Time `form:"createdAt,omitempty" json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Hash     string    `form:"hash,omitempty" json:"hash,omitempty" bson:"hash,omitempty"`
	PrevHash string    `form:"prevHash,omitempty" json:"prevHash,omitempty" bson:"prevHash,omitempty"`
}
