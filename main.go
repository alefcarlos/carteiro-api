package main

import (
	"log"
	"net/http"

	"github.com/alefcarlos/carteiro-api/router"
	"github.com/rs/cors"
)

func main() {

	r := router.GetRouter()

	log.Println("Estou escutando na porta 8080")
	handler := cors.Default().Handler(r)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
