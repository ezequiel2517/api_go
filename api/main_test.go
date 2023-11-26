package main

import (
	"api/db"
	"api/test"
	"fmt"
	"log"
	"os"
	"testing"
)

func truncateTables() {
	db := db.GetConnection()
	tables := []string{"users", "vaccinations", "drugs"}

	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table)
		_, err := db.Exec(query)
		if err != nil {
			log.Printf(fmt.Sprintf("Error al truncar la tabla %s: %v", table, err))
		}
	}
}

func TestMain(m *testing.M) {
	cargarEnv(".env.qa")
	truncateTables()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestSuite(t *testing.T) {
	t.Run("Test Signup", test.TestSignupHandler)
	t.Run("Test Login", test.TestLoginHandler)
}
