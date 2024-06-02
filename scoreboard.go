package domain

import "time"

type Scoreboard struct {
	Year        int       `json:"year"`
	Week        int       `json:"week"`
	DateUpdated time.Time `json:"date_updated"`
	Games       []ScoreboardGame
}

func CreateScoreboard(year, week int, games []ScoreboardGame) Scoreboard {
	var dateUpdated time.Time

	for _, game := range games {
		if game.DateUpdated.After(dateUpdated) {
			dateUpdated = game.DateUpdated
		}
	}

	return Scoreboard{
		Year:        year,
		Week:        week,
		DateUpdated: dateUpdated,
		Games:       games,
	}
}

type ScoreboardGame struct {
	Year        int
	Week        int
	GameId      string     `json:"game_id"`
	DateUpdated time.Time  `json:"date_updated"`
	Away        Team       `json:"away"`
	Home        Team       `json:"home"`
	AwayScore   int        `json:"away_score"`
	HomeScore   int        `json:"home_score"`
	Winner      GameWinner `json:"winner"`
	GameStatus  GameStatus `json:"game_status"`
}
