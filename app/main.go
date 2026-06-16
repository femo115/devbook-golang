package main

import (
	utils "app/src/Utils"
	"app/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	utils.CarregarTamplates()
	r := router.Gerar()

	fmt.Println("Escutando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}
