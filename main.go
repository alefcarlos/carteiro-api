package main

import (
	"log"
	"net/http"

	"github.com/alefcarlos/carteiro-api/router"
)

func main() {

	r := router.GetRouter()
	log.Println("Estou escutando na porta 8080")

	// Use the HostSwitch to listen and serve on port 12345
	log.Fatal(http.ListenAndServe(":8080", r))
}
