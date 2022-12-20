package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErroAPI struct {
	Erro string `json:"erro"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	fmt.Println(data)
	fmt.Println(json.NewEncoder(w).Encode(data))

	if data != nil {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}

}

func ErrorHandler(w http.ResponseWriter, r *http.Response) {

	var erro ErroAPI

	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)

}
