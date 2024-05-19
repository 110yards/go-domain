package domain

import "time"

type GameWinner string

const (
	NoResult       GameWinner = "no_result"
	AwayGameWinner GameWinner = "away"
	HomeGameWinner GameWinner = "home"
	TieGameWinner  GameWinner = "tie"
)

type GameResult struct {
	GameId      string     `json:"game_id"`
	Year        int        `json:"year"`
	DateStart   time.Time  `json:"date_start"`
	GameNumber  int        `json:"game_number"`
	WeekNumber  int        `json:"week_number"`
	GameType    GameType   `json:"game_type"`
	GameStatus  GameStatus `json:"game_status"`
	Away        Team       `json:"away"`
	Home        Team       `json:"home"`
	AwayScore   int        `json:"away_score"`
	HomeScore   int        `json:"home_score"`
	Winner      GameWinner `json:"winner"`
	DateUpdated time.Time  `json:"date_updated"`
}
