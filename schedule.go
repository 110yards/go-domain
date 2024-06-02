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

func CreateSchedule(year int, sourceName string) Schedule {
	return Schedule{
		Year:        year,
		DateUpdated: time.Now(),
		SourceName:  sourceName,
		Games:       []ScheduleGame{},
	}
}

type ScheduleGame struct {
	Year       int       `json:"year"`
	Week       int       `json:"week"`
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
	year int,
	week int,
	gameId string,
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
		Year:       year,
		Week:       week,
		GameId:     gameId,
		DateStart:  dateStart,
		GameNumber: gameNumber,
		WeekNumber: weekNumber,
		GameType:   gameType,
		Away:       away,
		Home:       home,
		SourceName: sourceName,
		SourceId:   sourceId,
	}

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
