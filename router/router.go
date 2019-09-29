package router

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"proyecto/blocks/blockhisto"
	"proyecto/blocks/blockmedic"
	"proyecto/blocks/blockpatient"
	"proyecto/controller"
	"proyecto/estructura"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Medicos cadena de bloques de medicos
var Medicos []estructura.Medic

// Pacientes cadena de bloques de pacientes
var Pacientes []estructura.Patient

// Historia cadena de bloques del historial medico general
var Historia []estructura.History

var Usuarios []estructura.Userg

// MsgUser estructura que obtiene valores de datos en un POST de un medico y/o medico
type MsgUser struct {
	FirstName string
	LastName  string
	CI        string
	Age       int
}
type messageMedUser struct {
	Code string
	CI   string
}

// MsgHis estructura que obtiene valores de datos en un POST de un historial
type MsgHis struct {
	CodePat string
	CodeMed string
}

// HisMessage estructura que obtiene valores de datos en un POST de historial
type HisMessage struct {
	Paciente estructura.Patient
	Medico   estructura.Medic
}

var mutex = &sync.Mutex{}

// NewRoute makes a new route
func NewRoute() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	UserDefault := estructura.Userg{}
	UserDefault = estructura.Userg{"Imhotep", "el que viene en paz", " 26902610LP", 4000}
	go func() {
		t := time.Now()
		// genesisBlock := Block{}
		// genesisBlock = Block{0, t.String(), 1, "genesis", calculatHash(genesisBlock), calculatHash(genesisBlock)}
		// spew.Dump(genesisBlock)

		genPac := estructura.Patient{}
		genPac = estructura.Patient{0, "12345679", UserDefault, controller.CalculateHashPac(genPac), "genesis"}

		UserDefault.CI = "123321456CBB" // Para hacer controles de que un paciente no puede ser su propio medico

		genMed := estructura.Medic{}
		genMed = estructura.Medic{0, "233234423", UserDefault, t, controller.CalculateHashMed(genMed), "genesis"}

		genHist := estructura.History{}
		genHist = estructura.History{0, genPac, genMed, t, controller.CalculateHashHis(genHist), "genesis"}

		mutex.Lock()

		// Blockchain = append(Blockchain, genesisBlock)
		Medicos = append(Medicos, genMed)
		Pacientes = append(Pacientes, genPac)
		Historia = append(Historia, genHist)

		mutex.Unlock()
	}()
	log.Fatal(run())

}

// web server
func run() error {
	mux := makeMuxRouter()
	httpPort := os.Getenv("PORT")

	log.Println("HTTP Server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           "localhost:" + httpPort, // Debe ser IP del mismo PC, el puerto esta por defecto en 80, lo puedes cambiar en .env
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	return err
}

// create handlers
func makeMuxRouter() http.Handler {

	muxRouter := mux.NewRouter()
	// muxRouter.HandleFunc("/history", handleGetBlockchain).Methods("GET")
	// muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")

	muxRouter.HandleFunc("/user", handleWriteUser).Methods("POST")
	muxRouter.HandleFunc("/user", handleGetUser).Methods("GET")

	muxRouter.HandleFunc("/medic", handleWriteMedic).Methods("POST")
	muxRouter.HandleFunc("/medic", handleGetMedics).Methods("GET")

	muxRouter.HandleFunc("/patient", handleWritePat).Methods("POST")
	muxRouter.HandleFunc("/patient", handleGetPatients).Methods("GET")

	muxRouter.HandleFunc("/history", handleWriteHisto).Methods("POST")
	muxRouter.HandleFunc("/history", handleGetHistories).Methods("GET")

	// muxRouter.HandleFunc("/", handleUpdateBlock).Methods("PUT") // Hecho solo para pruebas
	return muxRouter
}

// POSTS  - -- - - - -- - - - - - - -  - - - - - -  - - - -  - -
func handleWriteMedic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg messageMedUser

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	prevBlock := Medicos[len(Medicos)-1]
	newBlock := blockmedic.GenerateBlockmedic(prevBlock, msg.Code, msg.CI, Usuarios)
	if controller.IsMedicValid(newBlock, prevBlock) {
		Medicos = append(Medicos, newBlock)
		spew.Dump(Medicos)
	}
	mutex.Unlock()

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}
func handleWritePat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg messageMedUser

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	prevBlock := Pacientes[len(Pacientes)-1]
	newBlock := blockpatient.GenerateBlockPatient(prevBlock, msg.Code, msg.CI, Usuarios)
	if controller.IsPatientValid(newBlock, prevBlock) {
		Pacientes = append(Pacientes, newBlock)
		spew.Dump(Pacientes)
	}
	mutex.Unlock()

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}
func handleWriteHisto(w http.ResponseWriter, r *http.Request) { //
	w.Header().Set("Content-Type", "application/json")
	var msg MsgHis

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	prevBlock := Historia[len(Historia)-1]
	newBlock := blockhisto.GenerateBlockHisto(prevBlock, msg.CodeMed, msg.CodePat, Medicos, Pacientes)
	if controller.IsHistoryValid(newBlock, prevBlock) && controller.ExistsGuys(Pacientes, Medicos, msg.CodePat, msg.CodeMed) {
		Historia = append(Historia, newBlock)
		spew.Dump(Historia)
	}
	mutex.Unlock()

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}
func handleWriteUser(w http.ResponseWriter, r *http.Request) { //
	w.Header().Set("Content-Type", "application/json")
	var msg MsgUser

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	newBlock := estructura.Userg{msg.FirstName, msg.LastName, msg.CI, msg.Age}
	if controller.NotExistsCI(Usuarios, msg.CI) {
		Usuarios = append(Usuarios, newBlock)
		spew.Dump(Usuarios)
	}
	mutex.Unlock()

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

// GETTERS - - - - - - - - - - - - -  - - - - - -  - - - - -  --
func handleGetMedics(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Medicos, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func handleGetPatients(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Pacientes, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func handleGetHistories(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Historia, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Usuarios, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

//
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

// // Block asfvesegfdfgwesrg
// type Block struct {
// 	Index     int
// 	Timestamp string
// 	BPM       int
// 	Name      string
// 	Hash      string
// 	PrevHash  string
// }
// // Blockchain asdfewfksjdfpaosefmpsekmdfasmlkf
// var Blockchain []Block

// // Message asdwfqwsaf qwsdc wef qwsfdc sd
// type Message struct {
// 	BPM  int
// 	Name string
// }

// // // create a new block using previous block's hash
// func generateBlock(oldBlock Block, BPM int, Name string) Block {
// 	var newBlock Block
// 	t := time.Now()
// 	newBlock.Index = oldBlock.Index + 1
// 	newBlock.Timestamp = t.String()
// 	newBlock.BPM = BPM
// 	newBlock.Name = Name
// 	newBlock.PrevHash = oldBlock.Hash
// 	newBlock.Hash = calculatHash(newBlock)

// 	return newBlock
// }
// // SHA256 hasing
// func calculatHash(block Block) string {
// 	// dao.Company := c
// 	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.Name + block.PrevHash
// 	h := sha256.New()
// 	h.Write([]byte(record))
// 	hashed := h.Sum(nil)
// 	return hex.EncodeToString(hashed)
// }

// // make sure block is valid by checking index, and comparing the hash of the previous block
// func isBlockValid(newBlock, oldBlock Block) bool {
// 	if oldBlock.Index+1 != newBlock.Index {
// 		return false
// 	}

// 	if oldBlock.Hash != newBlock.PrevHash {
// 		return false
// 	}

// 	if calculatHash(newBlock) != newBlock.Hash {
// 		return false
// 	}

// 	return true
// }
// // write blockchain when we receive an http request
// func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
// 	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	io.WriteString(w, string(bytes))
// }

// // takes JSON payload as an input for heart rate (BPM)
// func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
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
// 	newBlock := generateBlock(prevBlock, msg.BPM, msg.Name)
// 	// comp := isBlockValid(newBlock, prevBlock)
// 	if isBlockValid(newBlock, prevBlock) {
// 		Blockchain = append(Blockchain, newBlock)
// 		spew.Dump(Blockchain)
// 	}
// 	mutex.Unlock()

// 	respondWithJSON(w, r, http.StatusCreated, newBlock)

// }
