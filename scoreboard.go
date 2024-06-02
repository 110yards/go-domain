package domain

import "time"

type Scoreboard struct {
	Year        int       `json:"year"`
	Week        int       `json:"week"`
	DateUpdated time.Time `json:"date_updated"`
}

type ScoreboardGame struct {
	Year       int
	Week       int
	GameId     string     `json:"game_id"`
	Away       Team       `json:"away"`
	Home       Team       `json:"home"`
	AwayScore  int        `json:"away_score"`
	HomeScore  int        `json:"home_score"`
	Winner     GameWinner `json:"winner"`
	GameStatus GameStatus `json:"game_status"`
}
