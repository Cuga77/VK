package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/Cuga77/filmoteca/handlers"
	"github.com/Cuga77/filmoteca/models"
)

func TestHandleCreateMovie(t *testing.T) {
	movie := &models.Movie{
		// Заполните поля структуры Movie
	}

	b, err := json.Marshal(movie)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/movie", bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleCreateMovie)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestHandleGetMovie(t *testing.T) {
	req, err := http.NewRequest("GET", "/movie?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleGetMovie)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}