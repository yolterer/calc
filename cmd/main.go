package main

import (
	"encoding/json"
	"net/http"

	"github.com/yolterer/calc/internal/calc"
)

func main() {
	http.HandleFunc("/api/v1/calculate", calcHandler)
	http.ListenAndServe(":8080", nil)
}

type MyHandler struct{}

type CalculateRequest struct {
	Expression string `json:"expression"`
}

func calcHandler(w http.ResponseWriter, r *http.Request) {

	var response map[string]interface{}

	errorResponse, result, code := calcRun(w, r)

	if errorResponse != "" {
		response = map[string]interface{}{
			"error": errorResponse,
		}
	} else {
		response = map[string]interface{}{
			"result": result,
		}
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func calcRun(w http.ResponseWriter, r *http.Request) (string, float64, int) {
	if r.Method != http.MethodPost {
		return "Internal server error", 0, http.StatusInternalServerError
	}
	var req CalculateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return "Internal server error", 0, http.StatusInternalServerError
	}
	result, err := calc.Calc(req.Expression)
	if err != nil && err.Error() != "invalid parentheses" {
		return "Internal server error", 0, http.StatusInternalServerError
	}
	if err != nil && err.Error() == "invalid parentheses" {
		return "Expression is not valid", 0, http.StatusUnprocessableEntity
	}

	return "", result, http.StatusOK
}
