package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/router"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Não foi possível encontrar o arquiv .env")
	}

	r := router.GetRouter()

	handler := cors.Default().Handler(r)

	port := os.Getenv("WEB_API_PORT")
	log.Println("Estou escutando na porta ", port)
	err = http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
