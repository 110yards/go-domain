package domain

type Team struct {
	City string
	Name string
	Abbr string
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
		City: "Free agents",
		Name: "Free agents",
		Abbr: "FA",
	}
}

func TeamBC() Team {
	return Team{
		City: "BC",
		Name: "Lions",
		Abbr: "BC",
	}
}

func TeamCGY() Team {
	return Team{
		City: "Calgary",
		Name: "Stampeders",
		Abbr: "CGY",
	}
}

func TeamEDM() Team {
	return Team{
		City: "Edmonton",
		Name: "Elks",
		Abbr: "EDM",
	}
}

func TeamSSK() Team {
	return Team{
		City: "Saskatchewan",
		Name: "Roughriders",
		Abbr: "SSK",
	}
}

func TeamWPG() Team {
	return Team{
		City: "Winnipeg",
		Name: "Blue Bombers",
		Abbr: "WPG",
	}
}

func TeamHAM() Team {
	return Team{
		City: "Hamilton",
		Name: "Tiger-Cats",
		Abbr: "HAM",
	}
}

func TeamTOR() Team {
	return Team{
		City: "Toronto",
		Name: "Argonauts",
		Abbr: "TOR",
	}
}

func TeamOTT() Team {
	return Team{
		City: "Ottawa",
		Name: "REDBLACKS",
		Abbr: "OTT",
	}
}

func TeamMTL() Team {
	return Team{
		City: "Montreal",
		Name: "Alouettes",
		Abbr: "MTL",
	}
}
