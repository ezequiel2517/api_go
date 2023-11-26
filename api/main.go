package main

import (
	"api/api"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func cargarEnv(env string) {
	err := godotenv.Load(env)
	if err != nil {
		log.Fatalf("ERROR AL CARGAR .ENV: %v", err)
	}
}

func main() {
	cargarEnv(".env.dev")
	server := api.New(os.Getenv("PUERTO"))
	server.ListenAndServe()
}
