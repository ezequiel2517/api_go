package test

import (
	"api/api/controllers/authController"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignupHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(authController.SignupHandler))
	jsonData := map[string]string{"Name": "value", "Email": "value", "Password": "value"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(server.URL+"/signup", "application/json", bytes.NewReader(jsonValue))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		t.Errorf("CODIGO ESPERADO: %d, CODIGO RECIBIDO: %d", http.StatusCreated, response.StatusCode)
	}

	expectedBodyMap := map[string]string{"ok": "Usuario registrado con exito."}
	expectedBody, _ := json.Marshal(expectedBodyMap)

	actualBody := readBody(response.Body)

	if strings.TrimSpace(actualBody) != strings.TrimSpace(string(expectedBody)) {
		t.Errorf("BODY ESPERADO: %s, BODY RECIBIDO: %s", expectedBody, actualBody)
	}
}
