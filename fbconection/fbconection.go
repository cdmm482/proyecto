package fbconection

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/lithammer/shortuuid"
	"google.golang.org/api/option"
)

// UpDate makes an update of the firebase
func UpDate() {
	var veces int = 2
	ctx := context.Background()
	// serviceKey.json son las credenciales de la base de datos FB
	sa := option.WithCredentialsFile("fbconection/serviceKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	//
	// Falta comprobar si es true el leido para que se actualice
	//
	//

	defer client.Close()
	for i := 1; i <= veces; {
		// tiempo dormido en segundos
		time.Sleep(30 * time.Second)
		// Texto a sobreescribir en formato JSON
		Lobby := map[string]interface{}{
			"CI":    "null",
			"Hash":  calculateHash(genShortUUID()),
			"leido": false,
		}
		// La llave del documento debe ser asignada a cada hospital
		bytesLeidos, err := ioutil.ReadFile("fbconection/IdDocumento.txt")
		if err != nil {
			fmt.Printf("Error leyendo archivo: %v", err)
		}
		// La variable contenido es el id del documento
		contenido := string(bytesLeidos)
		_, err = client.Collection("Consulta").Doc(contenido).Set(ctx, Lobby, firestore.MergeAll)
		fmt.Println(err)
	}
}

// Calcula el Hash del texto entrante y lo devuelve
func calculateHash(txt string) string {
	h := sha256.New()
	h.Write([]byte(txt))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Devuelve IDs unicos
func genShortUUID() string {
	id := shortuuid.New()
	// fmt.Printf("github.com/lithammer/shortuuid: %s\n", id)
	return id
}
