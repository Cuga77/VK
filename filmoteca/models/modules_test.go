package models_test

import (
	"testing"

	"github.com/Cuga77/filmoteca/models"
)

func TestCreateActor(t *testing.T) {
	actor := models.Actor{
		123,
		"John",
		"Doe",
		"1990-01-01",
	}

	err := models.CreateActor(actor)
	if err != nil {
		t.Errorf("CreateActor failed: %v", err)
	}
}



func TestGetActor(t *testing.T) {
	actor, err := models.GetActor(1) // предполагается, что актер с ID 1 существует
	if err != nil {
		t.Errorf("GetActor failed: %v", err)
	}
	if actor.ID != 1 {
		t.Errorf("GetActor failed: expected ID 1, got %d", actor.ID)
	} else if actor.Name != "John" {
		t.Errorf("GetActor failed: expected name John, got %s", actor.Name)
	} else if actor.Surname != "Doe" {
		t.Errorf("GetActor failed: expected surname Doe, got %s", actor.Surname)
	} else if actor.BirthDate != "1990-01-01" {
		t.Errorf("GetActor failed: expected birth date 1990-01-01, got %s", actor.BirthDate)
	}
}

func TestUpdateActor(t *testing.T) {
	actor := models.Actor{
		1,
		"Michael",
		"Bell",
		"1998-02-05",
	}

	err := models.UpdateActor(actor)
	if err != nil {
		t.Errorf("UpdateActor failed: %v", err)
	}
}

func TestDeleteActor(t *testing.T) {
	err := models.DeleteActor(1) // предполагается, что актер с ID 1 существует
	if err != nil {
		t.Errorf("DeleteActor failed: %v", err)
	} else {
		t.Log("DeleteActor passed")
	}
}

func TestCreateMovie(t *testing.T) {
    movie := models.Movie{
		"Drive",
		"2011",
		"Nicolas Winding Refn",
		"Crime",
		"Ryan Gosling",
		"Carey Mulligan",
    }

    err := models.CreateMovie(movie)
    if err != nil {
        t.Errorf("CreateMovie failed: %v", err)
    }
}

func TestGetMovie(t *testing.T) {
    movie, err := models.GetMovie(1) // предполагается, что фильм с ID 1 существует
    if err != nil {
        t.Errorf("GetMovie failed: %v", err)
    }
	if movie.ID != 1 {
		t.Errorf("GetMovie failed: expected ID 1, got %d", movie.ID)
	} else if movie.Title != "Drive" {
		t.Errorf("GetMovie failed: expected title Drive, got %s", movie.Title)
	} else if movie.Year != "2011" {
		t.Errorf("GetMovie failed: expected year 2011, got %s", movie.Year)
	} else if movie.Director != "Nicolas Winding Refn" {
		t.Errorf("GetMovie failed: expected director Nicolas Winding Refn, got %s", movie.Director)
	} else if movie.Genre != "Crime" {
		t.Errorf("GetMovie failed: expected genre Crime, got %s", movie.Genre)
	} else if movie.Actors != "Ryan Gosling" {
		t.Errorf("GetMovie failed: expected actor Ryan Gosling, got %s", movie.Actors)
	} else if movie.Actresses != "Carey Mulligan" {
		t.Errorf("GetMovie failed: expected actress Carey Mulligan, got %s", movie.Actresses)
	}
}

func TestUpdateMovie(t *testing.T) {
    movie := models.Movie{
        "Boys",
		"2014",
		"Kesha",
		"Comedy",
		"John Doe",
		"Jane Doe",
    }

    err := models.UpdateMovie(movie)
    if err != nil {
        t.Errorf("UpdateMovie failed: %v", err)
    }
}

func TestDeleteMovie(t *testing.T) {
    err := models.DeleteMovie(1)
    if err != nil {
        t.Errorf("DeleteMovie failed: %v", err)
    } 
}