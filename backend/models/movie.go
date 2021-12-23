package models

type Movies struct {
	MovieList []map[string]interface{} `json:"movie_list"`
}

type Movie struct {
	Id              int     `json:"id" gorm:"primaryKey"`
	Movie           string  `json:"movie"`
	Year            int     `json:"year"`
	Imdb            float64 `json:"imdb"`
	Image           []byte  `json:"image"`
	Genre           string  `json:"genre"`
	Actors          string  `json:"actors"`
	Likes           int     `json:"likes"`
	Dislikes        int     `json:"dislikes"`
	Summary         string  `json:"summary"`
	AddedBy         string  `json:"added_by" gorm:"column:addedBy"`
	Movie_or_series string  `json:"movie_or_series"`
	Trailer         string  `json:"trailer"`
}

func (Movie) TableName() string {
	return "movies"
}
