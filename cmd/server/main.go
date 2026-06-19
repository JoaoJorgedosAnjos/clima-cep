package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joaojorgedosanjos/clima-cep/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /clima/{cep}", handler.ClimaHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
