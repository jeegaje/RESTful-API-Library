package models

import "time"

type Book struct {
	//book table
	ID              uint
	Title           string `json:"title"`
	Description     string `json:"description"`
	BookLanguage    string `json:"book_language"`
	PublicationYear uint16 `json:"publication_year"`
	Publisher       string `json:"publisher"`
	Isbn            string `json:"isbn"`
	PageCount       uint16 `json:"page_count"`
	Price           uint   `json:"price"`
	Availability    string `json:"availability"`
	AverageRating   uint8  `json:"average_rating"`
	AuthorId        uint   `json:"author_id"`
	GenreId         uint   `json:"genre_id"`
	Author          Author
	Genre           Genre
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
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
	ID          uint
	FirstName   string    `json:"first_name" gorm:"not null"`
	LastName    string    `json:"last_name" gorm:"not null"`
	BirthDate   time.Time `json:"birth_date"`
	Nationality string    `json:"nationality"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
