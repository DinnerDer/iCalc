package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/DinnerDer/iCalc/pkg/calc"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var inputData struct {
		Expression string `json:"expression"`
	}
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	result, err := calc.Calc(inputData.Expression)
	if err != nil {
		if err.Error() == "invalid expression" {
			http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	responseData := map[string]interface{}{
		"result": result,
	}
	responseJSON, _ := json.Marshal(responseData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
