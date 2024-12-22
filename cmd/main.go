package main

import (
	"encoding/json"
	"net/http"

	"github.com/yolterer/calc/internal/calc"
)

func main() {
	http.HandleFunc("/api/v1/calculate", helloHandler)
	http.ListenAndServe(":8080", nil)
}

type MyHandler struct{}

type CalculateRequest struct {
	Expression string `json:"expression"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var errorResponse string
	if r.Method != http.MethodPost {
		errorResponse = "Internal server error"
		http.Error(w, errorResponse, http.StatusInternalServerError)
	}

	var req CalculateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errorResponse = "Internal server error"
		http.Error(w, errorResponse, http.StatusInternalServerError)
	}

	result, err := calc.Calc(req.Expression)
	if err != nil && err.Error() != "invalid parentheses" {
		errorResponse = "Internal server error"
		http.Error(w, errorResponse, http.StatusInternalServerError)
	}
	if err != nil && err.Error() == "invalid parentheses" {
		errorResponse = "Expression is not valid"
		http.Error(w, errorResponse, http.StatusUnprocessableEntity)
	}

	var response map[string]interface{}

	if errorResponse != "" {
		response = map[string]interface{}{
			"error": errorResponse,
		}
	} else {
		response = map[string]interface{}{
			"result": result,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
