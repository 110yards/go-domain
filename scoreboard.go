package domain

import "time"

type Scoreboard struct {
	Year        int       `json:"year"`
	Week        int       `json:"week"`
	LastUpdated time.Time `json:"last_updated"`
	Games       []ScoreboardGame
}

func CreateScoreboard(year, week int, games []ScoreboardGame) Scoreboard {
	var lastUpdated time.Time

	for _, game := range games {
		if game.LastUpdated.After(lastUpdated) {
			lastUpdated = game.LastUpdated
		}
	}

	return Scoreboard{
		Year:        year,
		Week:        week,
		LastUpdated: lastUpdated,
		Games:       games,
	}
}

type ScoreboardGame struct {
	Year        int
	Week        int
	GameId      string     `json:"game_id"`
	LastUpdated time.Time  `json:"last_updated"`
	Away        Team       `json:"away"`
	Home        Team       `json:"home"`
	AwayScore   int        `json:"away_score"`
	HomeScore   int        `json:"home_score"`
	Winner      GameWinner `json:"winner"`
	GameStatus  GameStatus `json:"game_status"`
}
