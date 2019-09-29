package blockmedic

import (
	"proyecto/controller"
	"proyecto/estructura"
	"time"
)

// GenerateBlockmedic genera un bloque del tipo Medic
func GenerateBlockmedic(medicLast estructura.Medic, code string, FirstName string, LastName string, CI string, Age int) estructura.Medic {
	var newBlock estructura.Medic
	t := time.Now()
	newBlock.Index = medicLast.Index + 1
	newBlock.Code = code
	newBlock.User.FirstName = FirstName
	newBlock.User.LastName = LastName
	newBlock.User.CI = CI
	newBlock.User.Age = Age
	newBlock.JoinedAt = t
	newBlock.PrevHash = medicLast.Hash
	newBlock.Hash = controller.CalculateHashMed(newBlock)

	return newBlock
}


