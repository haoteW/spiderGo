package model

const Domain = "https://movie.douban.com"

type MovieInfo struct {
	ID           string `json:"m_id"`
	URL          string `json:"m_url"`
	Director     string `json:"m_director"`
	Screenwriter string `json:"m_screenwriter"`
	Starring     string `json:"m_starring"`
	Type         string `json:"m_type"`
	Region       string `json:"m_region"`
	Language     string `json:"m_language"`
	ReleaseDate  string `json:"m_releaseDate"`
	Long         string `json:"m_long"`
	Sname        string `json:"m_sname"`
	Title        string `json:"m_title"`
}

type RankType struct {
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
