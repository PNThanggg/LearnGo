package models

import (
	"github.com/google/uuid"
)

type Genre struct {
	GenreId   uuid.UUID `json:"genre_id" validate:"required"`
	GenreName string    `json:"genre_name" validate:"required,min=2,max=100"`
}

type Ranking struct {
	RankingValue int    `json:"ranking_value" validate:"required"`
	RakingName   string `json:"raking_name" validate:"required"`
}

type Movie struct {
	ID          uuid.UUID `json:"id" db:"id" validate:"required"`
	ImdbID      string    `json:"imdb_id" db:"imdb_id" validate:"required"`
	Title       string    `json:"title" db:"title" validate:"required,min=2,max=500"`
	PosterPath  string    `json:"poster_path" db:"poster_path"`
	YouTubeURL  string    `json:"youtube_url" db:"youtube_url"`
	Genres      []Genre   `json:"genres" db:"genres"`
	AdminReview string    `json:"admin_review" db:"admin_review"`
	Rating      Ranking   `json:"rating" db:"rating"`
}
