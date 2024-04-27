package domain

import "time"

type Boxscore struct {
	GameId       string                `json:"game_id"`
	SourceGameId string                `json:"source_game_id"`
	SourceName   string                `json:"source_name"`
	DateStart    time.Time             `json:"date_start"`
	Status       GameStatus            `json:"status"`
	Score        LineScores            `json:"score"`
	Attendance   int                   `json:"attendance"`
	Teams        BoxScoreTeams         `json:"teams"`
	PlayerStats  []BoxscorePlayerStats `json:"player_stats"`
	TeamStats    []struct {
		Away TeamStats `json:"away"`
		Home TeamStats `json:"home"`
	} `json:"team_stats"`
}

type BoxScoreTeams struct {
	Away Team `json:"away"`
	Home Team `json:"home"`
}

type LineScores struct {
	Away LineScore `json:"away"`
	Home LineScore `json:"home"`
}

type LineScore struct {
	QuarterScores []int `json:"quarter_scores"`
	TotalScore    int   `json:"total_score"`
}

type Score struct {
	Away int `json:"away"`
	Home int `json:"home"`
}

type BoxscorePlayerStats struct {
	SourcePlayerId string      `json:"source_player_id"`
	Name           string      `json:"name"`
	Stats          PlayerStats `json:"stats"`
}
