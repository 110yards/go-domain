package domain

import "time"

type Schedule struct {
	Year       int            `json:"year"`
	Hash       string         `json:"hash"`
	SourceName string         `json:"source_name"`
	Games      []ScheduleGame `json:"games"`
	ByeWeeks   ByeWeeks       `json:"bye_weeks"`
}

type ScheduleGame struct {
	GameId     string    `json:"gameId"`
	SourceName string    `json:"source_name"`
	SourceUrl  string    `json:"source_url"`
	DateStart  time.Time `json:"date_start"`
	GameNumber int       `json:"game_number"`
	WeekNumber int       `json:"week_number"`
	GameType   string    `json:"game_type"`
	AwayAbbr   string    `json:"away_abbr"`
	HomeAbbr   string    `json:"home_abbr"`
}

type ByeWeeks struct {
	Bc  []int `json:"bc"`
	Cgy []int `json:"cgy"`
	Edm []int `json:"edm"`
	Ham []int `json:"ham"`
	Mtl []int `json:"mtl"`
	Ott []int `json:"ott"`
	Ssk []int `json:"ssk"`
	Tor []int `json:"tor"`
	Wpg []int `json:"wpg"`
}

func CreateByeWeeks() ByeWeeks {
	return ByeWeeks{
		Bc:  []int{},
		Cgy: []int{},
		Edm: []int{},
		Ham: []int{},
		Mtl: []int{},
		Ott: []int{},
		Ssk: []int{},
		Tor: []int{},
		Wpg: []int{},
	}
}

func (b ByeWeeks) AddByes(weekNumber int, teams []string) ByeWeeks {
	for _, team := range teams {
		switch team {
		case "bc":
			b.Bc = append(b.Bc, weekNumber)
		case "cgy":
			b.Cgy = append(b.Cgy, weekNumber)
		case "edm":
			b.Edm = append(b.Edm, weekNumber)
		case "ham":
			b.Ham = append(b.Ham, weekNumber)
		case "mtl":
			b.Mtl = append(b.Mtl, weekNumber)
		case "ott":
			b.Ott = append(b.Ott, weekNumber)
		case "ssk":
			b.Ssk = append(b.Ssk, weekNumber)
		case "tor":
			b.Tor = append(b.Tor, weekNumber)
		case "wpg":
			b.Wpg = append(b.Wpg, weekNumber)
		}
	}

	return b
}
