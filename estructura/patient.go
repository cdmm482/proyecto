package estructura

type Patient struct {
	Index    int    `form:"index,omitempty" json:"index,omitempty" bson:"index,omitempty"`
	Code     string `form:"code,omitempty" json:"code,omitempty" bson:"code,omitempty"`
	User     *Userg `form:"user,omitempty" json:"user,omitempty" bson:"user,omitempty"`
	Hash     string `form:"hash,omitempty" json:"hash,omitempty" bson:"hash,omitempty"`
	PrevHash string `form:"prevhash,omitempty" json:"prevhash,omitempty" bson:"prevhash,omitempty"`
}
