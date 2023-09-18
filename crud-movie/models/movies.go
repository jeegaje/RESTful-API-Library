package models

import "time"

type Movies struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(300)" json:"title"`
	Budget      int64     `gorm:"type:bigint" json:"budget"`
	Overview    string    `gorm:"type:text" json:"overview"`
	Popularity  uint8     `gorm:"type:tinyint" json:"popularity"`
	ReleaseDate time.Time `json:"release_date"`
	Revenue     int64     `gorm:"type:bigint" json:"revenue"`
	Runtime     uint16    `gorm:"type:smallint" json:"runtime"`
	MovieStatus string    `gorm:"varchar(100)" json:"movie_status"`
	Tagline     string    `gorm:"varchar(300)" json:"tagline"`
	VotesAvg    uint8     `gorm:"type:tinyint" json:"votes_avg"`
	VotesCount  uint32    `gorm:"type:int" json:"votes_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
