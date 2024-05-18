package domain

import "time"

type ScoreboardGame struct {
	GameId     string     `json:"game_id"`
	DateStart  time.Time  `json:"date_start"`
	GameNumber int        `json:"game_number"`
	GameType   GameType   `json:"game_type"`
	GameStatus GameStatus `json:"game_status"`
	Away       Team       `json:"away"`
	Home       Team       `json:"home"`
	AwayScore  int        `json:"away_score"`
	HomeScore  int        `json:"home_score"`
}

type Scoreboard struct {
	Games      []ScoreboardGame `json:"games"`
	WeekNumber int              `json:"week_number"`
}
