package domain

type Team struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	Abbr     string `json:"abbr"`
}

func TeamFromAbbr(abbr string) Team {
	switch abbr {
	case "BC":
		return TeamBC()
	case "CGY":
		return TeamCGY()
	case "EDM":
		return TeamEDM()
	case "SSK":
		return TeamSSK()
	case "WPG":
		return TeamWPG()
	case "HAM":
		return TeamHAM()
	case "TOR":
		return TeamTOR()
	case "OTT":
		return TeamOTT()
	case "MTL":
		return TeamMTL()
	default:
		return TeamFreeAgents()
	}
}

func (t Team) IsFreeAgent() bool {
	return t.Abbr == "FA"
}

func TeamFreeAgents() Team {
	return Team{
		Location: "Free agents",
		Name:     "Free agents",
		Abbr:     "FA",
	}
}

func TeamBC() Team {
	return Team{
		Location: "BC",
		Name:     "Lions",
		Abbr:     "BC",
	}
}

func TeamCGY() Team {
	return Team{
		Location: "Calgary",
		Name:     "Stampeders",
		Abbr:     "CGY",
	}
}

func TeamEDM() Team {
	return Team{
		Location: "Edmonton",
		Name:     "Elks",
		Abbr:     "EDM",
	}
}

func TeamSSK() Team {
	return Team{
		Location: "Saskatchewan",
		Name:     "Roughriders",
		Abbr:     "SSK",
	}
}

func TeamWPG() Team {
	return Team{
		Location: "Winnipeg",
		Name:     "Blue Bombers",
		Abbr:     "WPG",
	}
}

func TeamHAM() Team {
	return Team{
		Location: "Hamilton",
		Name:     "Tiger-Cats",
		Abbr:     "HAM",
	}
}

func TeamTOR() Team {
	return Team{
		Location: "Toronto",
		Name:     "Argonauts",
		Abbr:     "TOR",
	}
}

func TeamOTT() Team {
	return Team{
		Location: "Ottawa",
		Name:     "REDBLACKS",
		Abbr:     "OTT",
	}
}

func TeamMTL() Team {
	return Team{
		Location: "Montreal",
		Name:     "Alouettes",
		Abbr:     "MTL",
	}
}
