package blockpatient

import (
	"proyecto/controller"
	"proyecto/estructura"
)

// GenerateBlockPatient genera un bloque del tipo Patient
func GenerateBlockPatient(patientLast estructura.Patient, code string, FirstName string, LastName string, CI string, Age int) estructura.Patient {
	var newBlock estructura.Patient
	newBlock.Index = patientLast.Index + 1
	newBlock.Code = code
	newBlock.User.FirstName = FirstName
	newBlock.User.LastName = LastName
	newBlock.User.CI = CI
	newBlock.User.Age = Age
	newBlock.PrevHash = patientLast.Hash
	newBlock.Hash = controller.CalculateHashPac(newBlock)

	return newBlock
}
