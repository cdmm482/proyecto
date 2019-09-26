package blockmedic

import (
	"crypto/sha256"
	"encoding/hex"
	"proyecto/estructura"
	"strconv"
	"time"
)

func generateblockmedic(medicLast estructura.Medic, medicNew estructura.Medic, code string, FirstName string, LastName string, CI string, Age int) estructura.Medic {
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
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}
func calculateHash(med estructura.Medic) string {
	record := strconv.Itoa(med.Index) + med.Code + med.User.FirstName + med.User.LastName + med.User.CI + strconv.Itoa(med.User.Age) + med.JoinedAt.String() + med.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// func handleUpdateBlock(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var msg Message

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&msg); err != nil {
// 		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
// 		return
// 	}
// 	defer r.Body.Close()

// 	mutex.Lock()
// 	prevBlock := Blockchain[len(Blockchain)-1]
// 	newBlock := generateBlock(prevBlock, msg.BPM)

// 	if isBlockValid(newBlock, prevBlock) {
// 		Blockchain = append(Blockchain, newBlock)
// 		spew.Dump(Blockchain)
// 	}
// 	mutex.Unlock()

// 	respondWithJSON(w, r, http.StatusCreated, newBlock)

// }
