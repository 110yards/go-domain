package domain

import "time"

type GameWinner string

const (
	NoResult       GameWinner = "no_result"
	AwayGameWinner GameWinner = "away"
	HomeGameWinner GameWinner = "home"
	TieGameWinner  GameWinner = "tie"
)

type GameResult struct {
	GameId      string     `json:"game_id"`
	Year        int        `json:"year"`
	Week        int        `json:"week"`
	DateStart   time.Time  `json:"date_start"`
	Status      GameStatus `json:"game_status"`
	Away        Team       `json:"away"`
	Home        Team       `json:"home"`
	AwayScore   int        `json:"away_score"`
	HomeScore   int        `json:"home_score"`
	Winner      GameWinner `json:"winner"`
	DateUpdated time.Time  `json:"date_updated"`
}

func GameResultFromBox(boxscore Boxscore) GameResult {
	return GameResult{
		GameId:      boxscore.GameId,
		Year:        boxscore.Year,
		Week:        boxscore.Week,
		DateStart:   boxscore.DateStart,
		Status:      boxscore.Status,
		Away:        boxscore.Teams.Away,
		Home:        boxscore.Teams.Home,
		AwayScore:   boxscore.Score.Away.TotalScore,
		HomeScore:   boxscore.Score.Home.TotalScore,
		Winner:      getWinner(boxscore),
		DateUpdated: boxscore.DateUpdated,
	}
}

func getWinner(boxscore Boxscore) GameWinner {
	if boxscore.Status.StatusId != GameStatusId_Final {
		return NoResult
	}

	if boxscore.Score.Away.TotalScore > boxscore.Score.Home.TotalScore {
		return AwayGameWinner
	}

	if boxscore.Score.Away.TotalScore < boxscore.Score.Home.TotalScore {
		return HomeGameWinner
	}

	return TieGameWinner
}
