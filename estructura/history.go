package estructura

import "time"

// History es una estructura para guardar los datos de consulta por paciente
type History struct {
	Index    int       `form:"index,omitempty" json:"index,omitempty" bson:"index,omitempty"`
	Patien   Patient   `form:"patien,omitempty" json:"patien,omitempty" bson:"patien,omitempty"`
	Medi     Medic     `form:"medi,omitempty" json:"medi,omitempty" bson:"medi,omitempty"`
	QueryAt  time.Time `form:"queryAt,omitempty" json:"queryAt,omitempty" bson:"queryAt,omitempty"`
	Hash     string    `form:"hash,omitempty" json:"hash,omitempty" bson:"hash,omitempty"`
	PrevHash string    `form:"prevHash,omitempty" json:"prevHash,omitempty" bson:"prevHash,omitempty"`
}
