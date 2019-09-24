package fbconection

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func repetir(veces int) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("serviceKey.json")
	// fmt.Println("ALGO leio el JSOn")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	for i := 1; i <= veces; i++ {
		// fmt.Printf("#%v > %v\n", i, texto)
		time.Sleep(5000 * time.Millisecond)

		Lobby := map[string]interface{}{
			"CI": "null",
			// "Hash": calculateHash(i, "yonose"),
		}
		_, err = client.Collection("Consulta").Doc("04c2Hc-tbSh8Wn3Hw4-nLjwhSHN_uVC2mcq65Yv4tcQ").Set(ctx, Lobby, firestore.MergeAll)
		fmt.Println(err)
	}
}
