package main

import (
	"fmt"
	"proyecto/fbconection"
	"proyecto/router"
)

func main() {
	go fbconection.UpDate()
	r := router.NewRoute()
	fmt.Println(r)

}
