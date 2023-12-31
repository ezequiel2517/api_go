package vaccionationRoutes

import (
	"api/api/controllers/vaccinationController"
	"api/api/middlewares"
	"encoding/json"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/vaccinations", middlewares.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			vaccinationController.InsertVaccination(w, r)
		case http.MethodGet:
			vaccinationController.GetVaccinations(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
			errorMessage := map[string]string{"error": "Método no permitido para la ruta"}
			json.NewEncoder(w).Encode(errorMessage)
		}
	}))

	http.HandleFunc("/vaccinations/", middlewares.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			vaccinationController.DeleteVaccination(w, r)
		case http.MethodPut:
			vaccinationController.UpdateVaccination(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
			errorMessage := map[string]string{"error": "Método no permitido para la ruta"}
			json.NewEncoder(w).Encode(errorMessage)
		}
	}))
}
