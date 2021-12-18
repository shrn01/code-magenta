package models

type Movies struct {
	MovieIdList []int `json:"movie_id_list,omitempty"`
}

type Movie struct {
	Id              int     `json:"id"`
	Movie           string  `json:"movie"`
	Year            int     `json:"year"`
	Imdb            float64 `json:"imdb"`
	Image           int64   `json:"image"`
	Genre           string  `json:"genre"`
	Actors          string  `json:"actors"`
	Likes           int     `json:"likes"`
	Dislikes        int     `json:"dislikes"`
	Summary         string  `json:"summary"`
	AddedBy         string  `json:"added_by"`
	Movie_or_series string  `json:"movie_or_series"`
	Trailer         string  `json:"trailer"`
}
