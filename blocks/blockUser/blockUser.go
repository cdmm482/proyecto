package blockUser

import (
	"proyecto/estructura"
)

// GenerateBlockPatient genera un bloque del tipo Patient
func GenerateBlockUser(userLast estructura.Userg, Name string, LastN string, ci string, age int) estructura.Userg {
	var newBlock estructura.Userg
	newBlock.FirstName = Name
	newBlock.LastName = LastN
	newBlock.CI = ci
	newBlock.Age = age
	return newBlock
}
