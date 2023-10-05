package handlers

import (
	"binary-tree/services"
	"binary-tree/types"
	"encoding/json"
	"fmt"
	"net/http"
)

// CalculationHandler @Summary Calculate maximum path sum in a binary tree
// @Description Accepts an array of nodes and returns the sum of max path.
// @Accept  json
// @Produce  json
// @Param items body []types.CalculationRequest true
// @Success 200 {object} types.Response
// @Failure 400 {object} types.Response "Invalid input"
// @Router /calculate [post]
func CalculationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var statusCode int
	var response types.Response
	defer func() {
		if err != nil {
			response.Error = err.Error()
		} else {
			statusCode = http.StatusOK
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	}()
	if r.Method != http.MethodPost {
		err = fmt.Errorf("only POST requests are allowed")
		statusCode = http.StatusMethodNotAllowed
		return
	}
	var req types.CalculationRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		statusCode = http.StatusBadRequest
		return
	}
	sum := services.MaxPathSum(req)
	response.Sum = sum
}
