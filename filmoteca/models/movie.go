package models


type Movie struct {
	ID          int		`bd:"id"`
	Title       string `bd:"title"`
	Description string `bd:"description"`
	Year        int    `bd:"year"`
}

func CreateMovie(movie Movie) error {
	_, err := db.Exec("INSERT INTO movies (title, year) VALUES ($1, $2)", movie.Title, movie.Year)
	return err
}

func GetMovie(id int) (Movie, error) {
	var movie Movie
	err := db.QueryRow("SELECT id, title, year FROM movies WHERE id = $1", id).Scan(&movie.ID, &movie.Title, &movie.Year)
	return movie, err
}

func UpdateMovie(movie Movie) error {
	_, err := db.Exec("UPDATE movies SET title = $1, year = $2 WHERE id = $3", movie.Title, movie.Year, movie.ID)
	return err
}

func DeleteMovie(id int) error {
	_, err := db.Exec("DELETE FROM movies WHERE id = $1", id)
	return err
}