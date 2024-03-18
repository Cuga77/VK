// models/actor.go
package models

type Actor struct {
	// ID represents the unique identifier of an actor.
	ID   int    `json:"id`
	Name string `json:"name"`
	Age  int    `json:"age"`
	sex  string `json:"sex"`
}

func CreateActor(actor Actor) error {
	_, err := db.Exec("INSERT INTO actors (name) VALUES ($1)", actor.Name)
	return err
}

func GetActor(id int) (Actor, error) {
	var actor Actor
	err := db.QueryRow("SELECT id, name FROM actors WHERE id = $1", id).Scan(&actor.ID, &actor.Name)
	return actor, err
}

func UpdateActor(actor Actor) error {
	_, err := db.Exec("UPDATE actors SET name = $1 WHERE id = $2", actor.Name, actor.ID)
	return err
}

func DeleteActor(id int) error {
	_, err := db.Exec("DELETE FROM actors WHERE id = $1", id)
	return err
}