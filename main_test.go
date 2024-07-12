package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandleRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleRequest)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Nieprawidłowy kod statusu: otrzymano %v, oczekiwano %v",
			status, http.StatusOK)
	}

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Nie można zdekodować odpowiedzi JSON: %v", err)
	}

	now := time.Now()
	expectedDate := now.Format("2006-01-02")
	if response.Date != expectedDate {
		t.Errorf("Nieprawidłowa data: otrzymano %v, oczekiwano %v",
			response.Date, expectedDate)
	}

	_, err = time.Parse("15:04:05", response.Time)
	if err != nil {
		t.Errorf("Nieprawidłowy format czasu: %v", err)
	}
}
