package blockmaker

import (
	"proyecto/estructura"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Name      string
	Hash      string
	PrevHash  string
}

func generateblockmedic(BPM int, Name string, cod string) estructura.Medic {
	var newBlock estructura.Medic
	t := time.Now()
	// newBlock.Index = oldBlock.Index + 1
	// newBlock.Code = cod
	newBlock.Code = t.String()
	// newBlock.Name = Name
	// newBlock.PrevHash = oldBlock.Hash
	// newBlock.Hash = calculateHash(newBlock)

	return newBlock
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
