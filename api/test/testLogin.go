package test

import (
	"api/api/controllers/authController"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func readBody(body io.ReadCloser) string {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, body)
	if err != nil {
		panic(err)
	}
	defer body.Close()
	return buf.String()
}

func TestLoginHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(authController.LoginHandler))
	jsonData := map[string]string{"Email": "value", "Password": "value"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(server.URL+"/login", "application/json", bytes.NewReader(jsonValue))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("CODIGO ESPERADO: %d, CODIGO RECIBIDO: %d", http.StatusOK, response.StatusCode)
	}

	var responseBody map[string]interface{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&responseBody); err != nil {
		t.Fatalf("ERROR AL DECODIFICAR RESPUESTA: %v", err)
	}

	if _, ok := responseBody["token"]; !ok {
		t.Error("NO SE ENCUENTRA EL TOKEN EN LA RESPUESTA")
	}
}
