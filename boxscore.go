package domain

import (
	"encoding/json"
	"time"

	"110yards.ca/libs/go/core/logger"
)

type Boxscore struct {
	GameId       string                `json:"game_id"`
	Hash         string                `json:"hash"`
	DateUpdated  time.Time             `json:"date_updated"`
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

func CreateBoxScore(
	gameId string,
	sourceGameId string,
	sourceName string,
	dateStart time.Time,
	status GameStatus,
	score LineScores,
	teams BoxScoreTeams,
	playerStats []BoxscorePlayerStats,
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
		PlayerStats:  playerStats,
		TeamStats:    teamStats,
		Attendance:   attendance,
	}

	boxscore.UpdateHash()

	return boxscore
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

func (b *Boxscore) UpdateHash() error {
	// don't include last update time in hash calculation
	copy := *b
	copy.DateUpdated = time.Time{}

	// marshal p to json
	j, err := json.Marshal(copy)

	if err != nil {
		logger.Errorf("Error while marshalling boxscore: %s", err.Error())
		return err
	}

	// calculate hash
	b.Hash = calculateHash(string(j))
	return nil
}
