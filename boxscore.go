package domain

import (
	"time"
)

type Boxscore struct {
	GameId       string        `json:"game_id"`
	Year         int           `json:"year"`
	Week         int           `json:"week"`
	DateUpdated  time.Time     `json:"date_updated"`
	SourceGameId string        `json:"source_game_id"`
	SourceName   string        `json:"source_name"`
	DateStart    time.Time     `json:"date_start"`
	Status       GameStatus    `json:"status"`
	Score        LineScores    `json:"score"`
	Attendance   int           `json:"attendance"`
	Teams        BoxScoreTeams `json:"teams"`
	TeamStats    []struct {
		Away TeamStats `json:"away"`
		Home TeamStats `json:"home"`
	} `json:"team_stats"`
}

func CreateBoxScore(
	gameId string,
	sourceGameId string,
	sourceName string,
	dateStart time.Time,
	status GameStatus,
	score LineScores,
	teams BoxScoreTeams,
	teamStats []struct {
		Away TeamStats `json:"away"`
		Home TeamStats `json:"home"`
	},
	attendance int,
) Boxscore {
	boxscore := Boxscore{
		GameId:       gameId,
		DateUpdated:  time.Now(),
		SourceGameId: sourceGameId,
		SourceName:   sourceName,
		DateStart:    dateStart,
		Status:       status,
		Score:        score,
		Teams:        teams,
		TeamStats:    teamStats,
		Attendance:   attendance,
	}

	return boxscore
}

type BoxScoreTeams struct {
	Away      Team                  `json:"away"`
	AwayStats []BoxscorePlayerStats `json:"away_stats"`
	Home      Team                  `json:"home"`
	HomeStats []BoxscorePlayerStats `json:"home_stats"`
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
	PlayerId       string      `json:"player_id"`
	GameId         string      `json:"game_id"`
	Year           int         `json:"year"`
	Week           int         `json:"week"`
	SourceName     string      `json:"source_name"`
	SourcePlayerId string      `json:"source_player_id"`
	SourceGameId   string      `json:"source_game_id"`
	Team           Team        `json:"team"`
	Opponent       Team        `json:"opponent"`
	Name           string      `json:"name"`
	Stats          PlayerStats `json:"stats"`
	DateUpdated    time.Time   `json:"date_updated"`
}
