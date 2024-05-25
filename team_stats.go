package domain

type TeamStats struct {
	TotalFirstDowns   int    `json:"total_first_downs"`
	Rushes            int    `json:"rushes"`
	RushYards         int    `json:"rush_yards"`
	RushTouchdowns    int    `json:"rush_touchdowns"`
	PassAttempts      int    `json:"pass_attempts"`
	PassCompletions   int    `json:"pass_completions"`
	PassGrossYards    int    `json:"pass_gross_yards"`
	PassTouchdowns    int    `json:"pass_touchdowns"`
	PassInterceptions int    `json:"pass_interceptions"`
	QbRating          string `json:"qb_rating"`
	SacksTaken        int    `json:"sacks_taken"`
	TotalGrossYards   int    `json:"total_gross_yards"`
	TimesFumbled      int    `json:"times_fumbled"`
	FumblesLost       int    `json:"fumbles_lost"`
	Turnovers         int    `json:"turnovers"`
	Penalties         int    `json:"penalties"`
	PenaltyYards      int    `json:"penalty_yards"`
}
