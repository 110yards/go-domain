package domain

import (
	"time"
)

type Schedule struct {
	Year        int            `json:"year"`
	DateUpdated time.Time      `json:"date_updated"`
	SourceName  string         `json:"source_name"`
	Games       []ScheduleGame `json:"games"`
}

type ScheduleGame struct {
	GameId     string    `json:"game_id"`
	SourceName string    `json:"source_name"`
	SourceId   string    `json:"source_id"`
	DateStart  time.Time `json:"date_start"`
	GameNumber int       `json:"game_number"`
	WeekNumber int       `json:"week_number"`
	GameType   GameType  `json:"game_type"`
	Away       Team      `json:"away"`
	Home       Team      `json:"home"`
}

func CreateScheduleGame(
	dateStart time.Time,
	gameNumber int,
	weekNumber int,
	gameType GameType,
	away Team,
	home Team,
	sourceName string,
	sourceId string,
) ScheduleGame {
	game := ScheduleGame{
		DateStart:  dateStart,
		GameNumber: gameNumber,
		WeekNumber: weekNumber,
		GameType:   gameType,
		Away:       away,
		Home:       home,
		SourceName: sourceName,
		SourceId:   sourceId,
	}

	game.GameId = generateGameId(dateStart.Year(), gameNumber)

	return game
}

type GameType string

const (
	GameTypePreseason         GameType = "preseason"
	GameTypeRegular           GameType = "regular"
	GameTypeDivisionSemiFinal GameType = "division-semifinal"
	GameTypeDivisionFinal     GameType = "division-final"
	GameTypeGreyCup           GameType = "grey-cup"
)
