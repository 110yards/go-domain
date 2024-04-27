package domain

type GameStatusId string

const (
	GameStatusId_PreGame    GameStatusId = "pre_game"
	GameStatusId_InProgress GameStatusId = "in_progress"
	GameStatusId_Final      GameStatusId = "final"
	GameStatusId_Postponed  GameStatusId = "postponed"
	GameStatusId_Cancelled  GameStatusId = "cancelled"
)

type GameStatus struct {
	StatusId   GameStatusId `json:"status_id"`
	IsActive   bool         `json:"is_active"`
	HasStarted bool         `json:"has_started"`
	Quarter    int          `json:"quarter"`
	Minutes    int          `json:"minutes"`
	Seconds    int          `json:"seconds"`
	Down       int          `json:"down"`
	YardsToGo  int          `json:"yards_to_go"`
}

func CreateGameStatusInProgress(
	quarter int,
	minutes int,
	seconds int,
	down int,
	yardsToGo int,
) GameStatus {
	return GameStatus{
		StatusId:   GameStatusId_InProgress,
		IsActive:   true,
		HasStarted: true,
		Quarter:    quarter,
		Minutes:    minutes,
		Seconds:    seconds,
		Down:       down,
		YardsToGo:  yardsToGo,
	}
}

func CreateGameStatus(statusId GameStatusId) GameStatus {
	if statusId == GameStatusId_InProgress {
		return CreateGameStatusInProgress(1, 0, 0, 1, 10)
	}

	hasStarted := statusId == GameStatusId_Final

	return GameStatus{
		StatusId:   statusId,
		IsActive:   false,
		HasStarted: hasStarted,
	}

}
