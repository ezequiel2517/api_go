package api_test

import (
	"api/api/controllers/authController"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func cargarEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR AL CARGAR ARCHIVO .env")
	}
}

func TestSumHandler(t *testing.T) {
	cargarEnv()
	server := httptest.NewServer(http.HandlerFunc(authController.SignupHandler))
	jsonData := map[string]string{"Name": "value", "Email": "value", "Password": "value"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(server.URL+"/signup", "application/json", bytes.NewReader(jsonValue))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Esperaba código de estado %d, pero obtuve %d", http.StatusOK, response.StatusCode)
	}

	// Verificar el cuerpo de la respuesta
	expectedBody := "Sum: 5"
	actualBody := readBody(response.Body)
	if actualBody != expectedBody {
		t.Errorf("Esperaba cuerpo %s, pero obtuve %s", expectedBody, actualBody)
	}
}

func readBody(body io.ReadCloser) string {
	// Leer el cuerpo de la respuesta
	buf := new(strings.Builder)
	_, err := io.Copy(buf, body)
	if err != nil {
		panic(err)
	}
	defer body.Close() // ¡No olvides cerrar el cuerpo después de leerlo!
	return buf.String()
}
