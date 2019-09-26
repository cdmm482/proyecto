package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"proyecto/estructura"
	"strconv"
)

// make sure block is valid by checking index, and comparing the hash of the previous block
func isMedicValid(newBlock estructura.Medic, oldBlock estructura.Medic) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func isPatientValid(newBlock estructura.Patient, oldBlock estructura.Patient) bool {
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

func calculateHash(med estructura.Medic) string {
	record := strconv.Itoa(med.Index) + med.Code + med.User.FirstName + med.User.LastName + med.User.CI + strconv.Itoa(med.User.Age) + med.JoinedAt.String() + med.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
