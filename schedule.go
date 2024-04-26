package domain

import (
	"fmt"
	"time"
)

type Schedule struct {
	Year       int            `json:"year"`
	SourceName string         `json:"source_name"`
	Games      []ScheduleGame `json:"games"`
	// ByeWeeks   ByeWeeks       `json:"bye_weeks"`
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

func generateGameId(year, gameNumber int) string {
	// yeargameNumber(00-padded)
	return fmt.Sprintf("%d%03d", year, gameNumber)
}

type GameType string

const (
	GameTypePreseason         GameType = "preseason"
	GameTypeRegular           GameType = "regular"
	GameTypeDivisionSemiFinal GameType = "division-semifinal"
	GameTypeDivisionFinal     GameType = "division-final"
	GameTypeGreyCup           GameType = "grey-cup"
)

// type ByeWeeks struct {
// 	Bc  []int `json:"bc"`
// 	Cgy []int `json:"cgy"`
// 	Edm []int `json:"edm"`
// 	Ham []int `json:"ham"`
// 	Mtl []int `json:"mtl"`
// 	Ott []int `json:"ott"`
// 	Ssk []int `json:"ssk"`
// 	Tor []int `json:"tor"`
// 	Wpg []int `json:"wpg"`
// }

// func CreateByeWeeks() ByeWeeks {
// 	return ByeWeeks{
// 		Bc:  []int{},
// 		Cgy: []int{},
// 		Edm: []int{},
// 		Ham: []int{},
// 		Mtl: []int{},
// 		Ott: []int{},
// 		Ssk: []int{},
// 		Tor: []int{},
// 		Wpg: []int{},
// 	}
// }

// func (b ByeWeeks) AddByes(weekNumber int, teams []string) ByeWeeks {
// 	for _, team := range teams {
// 		switch team {
// 		case "bc":
// 			b.Bc = append(b.Bc, weekNumber)
// 		case "cgy":
// 			b.Cgy = append(b.Cgy, weekNumber)
// 		case "edm":
// 			b.Edm = append(b.Edm, weekNumber)
// 		case "ham":
// 			b.Ham = append(b.Ham, weekNumber)
// 		case "mtl":
// 			b.Mtl = append(b.Mtl, weekNumber)
// 		case "ott":
// 			b.Ott = append(b.Ott, weekNumber)
// 		case "ssk":
// 			b.Ssk = append(b.Ssk, weekNumber)
// 		case "tor":
// 			b.Tor = append(b.Tor, weekNumber)
// 		case "wpg":
// 			b.Wpg = append(b.Wpg, weekNumber)
// 		}
// 	}

// 	return b
// }
