package model

const Domain = "https://movie.douban.com"

type MovieInfo struct {
	ID           string `json:"id"`
	URL          string `json:"url"`
	Director     string `json:"director"`
	Screenwriter string `json:"screenwriter"`
	Starring     string `json:"starring"`
	Type         string `json:"type"`
	Region       string `json:"region"`
	Language     string `json:"language"`
	ReleaseDate  string `json:"releaseDate"`
	Long         string `json:"long"`
	Sname        string `json:"sname"`
	Title        string `json:"title"`
}

type RankType []struct {
	ActorCount  int64    `json:"actor_count"`
	Actors      []string `json:"actors"`
	CoverURL    string   `json:"cover_url"`
	ID          string   `json:"id"`
	IsPlayable  bool     `json:"is_playable"`
	IsWatched   bool     `json:"is_watched"`
	Rank        int64    `json:"rank"`
	Rating      []string `json:"rating"`
	Regions     []string `json:"regions"`
	ReleaseDate string   `json:"release_date"`
	Score       string   `json:"score"`
	Title       string   `json:"title"`
	Types       []string `json:"types"`
	URL         string   `json:"url"`
	VoteCount   int64    `json:"vote_count"`
}
