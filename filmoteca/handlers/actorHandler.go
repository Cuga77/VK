// Package handlers provides HTTP request handlers.
package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/Cuga77/filmoteca/models"
)

// HandleCreateActor handles the creation of an actor.
func HandleCreateActor(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.CreateActor(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func HandleGetActor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	actor, err := models.GetActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(actor)
}

func HandleUpdateActor(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.UpdateActor(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func HandleDeleteActor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := models.DeleteActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}