package domain

import "strings"

type Position struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func PositionFromId(id string) Position {
	id = strings.ToLower(id)
	switch id {
	case "qb":
		return Position{Id: id, Name: "Quarterback"}
	case "rb":
		return Position{Id: id, Name: "Runningback"}
	case "fb":
		return Position{Id: id, Name: "Fullback"}
	case "wr":
		return Position{Id: id, Name: "WideReceiver"}
	case "db":
		return Position{Id: id, Name: "DefensiveBack"}
	case "de":
		return Position{Id: id, Name: "DefensiveEnd"}
	case "dl":
		return Position{Id: id, Name: "DefensiveLine"}
	case "k":
		return Position{Id: id, Name: "Kicker"}
	case "ol":
		return Position{Id: id, Name: "OffensiveLine"}
	case "p":
		return Position{Id: id, Name: "Punter"}
	case "lb":
		return Position{Id: id, Name: "Linebacker"}
	case "ls":
		return Position{Id: id, Name: "LongSnapper"}
	case "dt":
		return Position{Id: id, Name: "DefensiveTackle"}
	case "t":
		return Position{Id: id, Name: "Tackle"}
	case "te":
		return Position{Id: id, Name: "TightEnd"}
	case "s":
		return Position{Id: id, Name: "Safety"}
	case "c":
		return Position{Id: id, Name: "Center"}
	case "g":
		return Position{Id: id, Name: "Guard"}
	default:
		return Position{Id: "unknown", Name: "Unknown"}
	}
}
