package models

import "time"

type Book struct {
	//book table
	ID          uint
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorId    uint   `json:"author_id"`
	GenreId     uint   `json:"genre_id"`
	Author      Author
	Genre       Genre
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Genre struct {
	//Genre table
	ID        uint
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Author struct {
	//author table
	ID        uint
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
