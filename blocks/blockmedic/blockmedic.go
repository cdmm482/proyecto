package blockmedic

import (
	"proyecto/controller"
	"proyecto/estructura"
	"time"
)

// GenerateBlockmedic genera un bloque del tipo Medic
func GenerateBlockmedic(medicLast estructura.Medic, code string, ci string, Usuarios []estructura.Userg) estructura.Medic {
	var newBlock estructura.Medic
	t := time.Now()
	newBlock.Index = medicLast.Index + 1
	newBlock.Code = code
	newBlock.User = controller.FindUserByCI(ci, Usuarios)
	newBlock.JoinedAt = t
	newBlock.PrevHash = medicLast.Hash
	newBlock.Hash = controller.CalculateHashMed(newBlock)

	return newBlock
}
