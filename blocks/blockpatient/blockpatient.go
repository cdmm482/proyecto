package blockpatient

import (
	"proyecto/controller"
	"proyecto/estructura"
)

// GenerateBlockPatient genera un bloque del tipo Patient
func GenerateBlockPatient(patientLast estructura.Patient, code string, ci string, Usuarios []estructura.Userg) estructura.Patient {
	var newBlock estructura.Patient
	newBlock.Index = patientLast.Index + 1
	newBlock.Code = code
	newBlock.User = controller.FindUserByCI(ci, Usuarios)
	newBlock.PrevHash = patientLast.Hash
	newBlock.Hash = controller.CalculateHashPac(newBlock)

	return newBlock
}
