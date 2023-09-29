package models

import "time"

type Book struct {
	//book table
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	//category table
	Name string `json:"name"`
}

type Author struct {
	//author table
	Name string `json:"name"`
}
