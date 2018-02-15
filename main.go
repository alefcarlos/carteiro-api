package main

import (
	"log"
	"net/http"

	"github.com/alefcarlos/carteiro-api/repo"

	"github.com/alefcarlos/carteiro-api/router"
	"github.com/rs/cors"
)

func main() {

	err := repo.NewMongoSession()
	if err != nil {
		log.Fatal("Parando a carga do servidor. Erro ao abrir a sessao com o MongoDB: ", err.Error())
		return
	}

	r := router.GetRouter()

	log.Println("Estou escutando na porta 8080")
	handler := cors.Default().Handler(r)
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
