package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculationHandler(t *testing.T) {
	data := `{
"tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "3", "left": null, "right": null, "value": 3},
      {"id": "2", "left": null, "right": null, "value": 2}
],
"root": "1" }
}`
	req, _ := http.NewRequest("POST", "/calculate", bytes.NewBufferString(data))
	rr := httptest.NewRecorder()
	CalculationHandler(rr, req)

	expected := `{"sum":6}`
	response := strings.TrimSpace(rr.Body.String())
	if response != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", response, expected)
	}
}
