package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"proyecto/estructura"
	"strconv"
)

// IsMedicValid valida el bloque del Medico
func IsMedicValid(newBlock estructura.Medic, oldBlock estructura.Medic) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHashMed(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// IsPatientValid valida el bloque del paciente
func IsPatientValid(newBlock estructura.Patient, oldBlock estructura.Patient) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	// if calculateHash(newBlock) != newBlock.Hash {
	// 	return false
	// }

	return true
}

// IsHistoryValid valida el bloque de historial
func IsHistoryValid(newBlock estructura.History, oldBlock estructura.History) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	// if calculateHash(newBlock) != newBlock.Hash {
	// 	return false
	// }

	return true
}

// CalculateHashMed calcula el hash de un bloque de un medico
func CalculateHashMed(med estructura.Medic) string {
	record := strconv.Itoa(med.Index) + med.Code + med.User.FirstName + med.User.LastName + med.User.CI + strconv.Itoa(med.User.Age) + med.JoinedAt.String() + med.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// CalculateHashPac calcula el hash de un bloque paciente
func CalculateHashPac(pac estructura.Patient) string {
	record := strconv.Itoa(pac.Index) + pac.Code + pac.User.FirstName + pac.User.LastName + pac.User.CI + strconv.Itoa(pac.User.Age) + pac.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// CalculateHashHis calcula el hash de un bloque del historial
func CalculateHashHis(his estructura.History) string {
	record := strconv.Itoa(his.Index) + his.Patien.Code + his.Medi.Code + his.QueryAt.String() + his.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// FindUser encuentra al medico y paciente segun el codigo
func FindUser(codP string, codM string, Medicos []estructura.Medic, Pacientes []estructura.Patient) (estructura.Patient, estructura.Medic) {
	findM := estructura.Medic{}
	findP := estructura.Patient{}
	for i := 0; i < len(Medicos); i++ {
		if codM == Medicos[i].Code {
			findM = Medicos[i]
		}
	}
	for i := 0; i < len(Pacientes); i++ {
		if codP == Pacientes[i].Code {
			findP = Pacientes[i]
		}
	}
	return findP, findM
}
