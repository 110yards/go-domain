package domain

type PlayerStats struct {
	PassAttempts                         int `json:"pass_attempts"`
	PassCompletions                      int `json:"pass_completions"`
	PassNetYards                         int `json:"pass_net_yards"`
	PassLong                             int `json:"pass_long"`
	PassTouchdowns                       int `json:"pass_touchdowns"`
	PassLongTouchdowns                   int `json:"pass_long_touchdowns"`
	PassInterceptions                    int `json:"pass_interceptions"`
	RushAttempts                         int `json:"rush_attempts"`
	RushNetYards                         int `json:"rush_net_yards"`
	RushLong                             int `json:"rush_long"`
	RushTouchdowns                       int `json:"rush_touchdowns"`
	RushLongTouchdowns                   int `json:"rush_long_touchdowns"`
	ReceiveAttempts                      int `json:"receive_attempts"`
	ReceiveCaught                        int `json:"receive_caught"`
	ReceiveYards                         int `json:"receive_yards"`
	ReceiveLong                          int `json:"receive_long"`
	ReceiveTouchdowns                    int `json:"receive_touchdowns"`
	ReceiveLongTouchdowns                int `json:"receive_long_touchdowns"`
	ReceiveYardsAfterCatch               int `json:"receive_yards_after_catch"`
	Fumbles                              int `json:"fumbles"`
	FumblesLost                          int `json:"fumbles_lost"`
	Punts                                int `json:"punts"`
	PuntGrossYards                       int `json:"punt_gross_yards"`
	PuntNetYards                         int `json:"punt_net_yards"`
	PuntLong                             int `json:"punt_long"`
	PuntSingles                          int `json:"punt_singles"`
	PuntsBlocked                         int `json:"punts_blocked"`
	PuntsIn10                            int `json:"punts_in_10"`
	PuntsIn20                            int `json:"punts_in_20"`
	PuntsReturned                        int `json:"punts_returned"`
	KickReturns                          int `json:"kick_returns"`
	KickReturnsYards                     int `json:"kick_returns_yards"`
	KickReturnsTouchdowns                int `json:"kick_returns_touchdowns"`
	KickReturnsLong                      int `json:"kick_returns_long"`
	KickReturnsTouchdownsLong            int `json:"kick_returns_touchdowns_long"`
	FieldGoalAttempts                    int `json:"field_goal_attempts"`
	FieldGoalMade                        int `json:"field_goal_made"`
	FieldGoalMisses                      int `json:"field_goal_misses"`
	FieldGoalYards                       int `json:"field_goal_yards"`
	FieldGoalSingles                     int `json:"field_goal_singles"`
	FieldGoalLong                        int `json:"field_goal_long"`
	FieldGoalPoints                      int `json:"field_goal_points"`
	MissedFieldGoalReturns               int `json:"missed_field_goal_returns"`
	MissedFieldGoalReturnsYards          int `json:"missed_field_goal_returns_yards"`
	MissedFieldGoalReturnsTouchdowns     int `json:"missed_field_goal_returns_touchdowns"`
	MissedFieldGoalReturnsLong           int `json:"missed_field_goal_returns_long"`
	MissedFieldGoalReturnsTouchdownsLong int `json:"missed_field_goal_returns_touchdowns_long"`
	PuntReturns                          int `json:"punt_returns"`
	PuntReturnsYards                     int `json:"punt_returns_yards"`
	PuntReturnsTouchdowns                int `json:"punt_returns_touchdowns"`
	PuntReturnsLong                      int `json:"punt_returns_long"`
	PuntReturnsTouchdownsLong            int `json:"punt_returns_touchdowns_long"`
	Kicks                                int `json:"kicks"`
	KickYards                            int `json:"kick_yards"`
	KicksNetYards                        int `json:"kicks_net_yards"`
	KicksLong                            int `json:"kicks_long"`
	KicksSingles                         int `json:"kicks_singles"`
	KicksOutOfEndZone                    int `json:"kicks_out_of_end_zone"`
	KicksOnside                          int `json:"kicks_onside"`
	OnePointConvertsAttempts             int `json:"one_point_converts_attempts"`
	OnePointConvertsMade                 int `json:"one_point_converts_made"`
	TwoPointConvertsMade                 int `json:"two_point_converts_made"`
	TacklesTotal                         int `json:"tackles_total"`
	TacklesDefensive                     int `json:"tackles_defensive"`
	TacklesSpecialTeams                  int `json:"tackles_special_teams"`
	SacksQbMade                          int `json:"sacks_qb_made"`
	Interceptions                        int `json:"interceptions"`
	InterceptionYards                    int `json:"interceptions_yards"`
	InterceptionTouchdowns               int `json:"interception_touchdowns"`
	InterceptionLong                     int `json:"interception_long"`
	InterceptionTouchdownsLong           int `json:"interception_touchdowns_long"`
	FumblesForced                        int `json:"fumbles_forced"`
	FumblesRecovered                     int `json:"fumbles_recovered"`
	PassesKnockedDown                    int `json:"passes_knocked_down"`
}

func (ps PlayerStats) Add(other PlayerStats) PlayerStats {
	return PlayerStats{
		PassAttempts:                         ps.PassAttempts + other.PassAttempts,
		PassCompletions:                      ps.PassCompletions + other.PassCompletions,
		PassNetYards:                         ps.PassNetYards + other.PassNetYards,
		PassLong:                             ps.PassLong + other.PassLong,
		PassTouchdowns:                       ps.PassTouchdowns + other.PassTouchdowns,
		PassLongTouchdowns:                   ps.PassLongTouchdowns + other.PassLongTouchdowns,
		PassInterceptions:                    ps.PassInterceptions + other.PassInterceptions,
		RushAttempts:                         ps.RushAttempts + other.RushAttempts,
		RushNetYards:                         ps.RushNetYards + other.RushNetYards,
		RushLong:                             ps.RushLong + other.RushLong,
		RushTouchdowns:                       ps.RushTouchdowns + other.RushTouchdowns,
		RushLongTouchdowns:                   ps.RushLongTouchdowns + other.RushLongTouchdowns,
		ReceiveAttempts:                      ps.ReceiveAttempts + other.ReceiveAttempts,
		ReceiveCaught:                        ps.ReceiveCaught + other.ReceiveCaught,
		ReceiveYards:                         ps.ReceiveYards + other.ReceiveYards,
		ReceiveLong:                          ps.ReceiveLong + other.ReceiveLong,
		ReceiveTouchdowns:                    ps.ReceiveTouchdowns + other.ReceiveTouchdowns,
		ReceiveLongTouchdowns:                ps.ReceiveLongTouchdowns + other.ReceiveLongTouchdowns,
		ReceiveYardsAfterCatch:               ps.ReceiveYardsAfterCatch + other.ReceiveYardsAfterCatch,
		Fumbles:                              ps.Fumbles + other.Fumbles,
		FumblesLost:                          ps.FumblesLost + other.FumblesLost,
		Punts:                                ps.Punts + other.Punts,
		PuntGrossYards:                       ps.PuntGrossYards + other.PuntGrossYards,
		PuntNetYards:                         ps.PuntNetYards + other.PuntNetYards,
		PuntLong:                             ps.PuntLong + other.PuntLong,
		PuntSingles:                          ps.PuntSingles + other.PuntSingles,
		PuntsBlocked:                         ps.PuntsBlocked + other.PuntsBlocked,
		PuntsIn10:                            ps.PuntsIn10 + other.PuntsIn10,
		PuntsIn20:                            ps.PuntsIn20 + other.PuntsIn20,
		PuntsReturned:                        ps.PuntsReturned + other.PuntsReturned,
		KickReturns:                          ps.KickReturns + other.KickReturns,
		KickReturnsYards:                     ps.KickReturnsYards + other.KickReturnsYards,
		KickReturnsTouchdowns:                ps.KickReturnsTouchdowns + other.KickReturnsTouchdowns,
		KickReturnsLong:                      ps.KickReturnsLong + other.KickReturnsLong,
		KickReturnsTouchdownsLong:            ps.KickReturnsTouchdownsLong + other.KickReturnsTouchdownsLong,
		FieldGoalAttempts:                    ps.FieldGoalAttempts + other.FieldGoalAttempts,
		FieldGoalMade:                        ps.FieldGoalMade + other.FieldGoalMade,
		FieldGoalMisses:                      ps.FieldGoalMisses + other.FieldGoalMisses,
		FieldGoalYards:                       ps.FieldGoalYards + other.FieldGoalYards,
		FieldGoalSingles:                     ps.FieldGoalSingles + other.FieldGoalSingles,
		FieldGoalLong:                        ps.FieldGoalLong + other.FieldGoalLong,
		FieldGoalPoints:                      ps.FieldGoalPoints + other.FieldGoalPoints,
		MissedFieldGoalReturns:               ps.MissedFieldGoalReturns + other.MissedFieldGoalReturns,
		MissedFieldGoalReturnsYards:          ps.MissedFieldGoalReturnsYards + other.MissedFieldGoalReturnsYards,
		MissedFieldGoalReturnsTouchdowns:     ps.MissedFieldGoalReturnsTouchdowns + other.MissedFieldGoalReturnsTouchdowns,
		MissedFieldGoalReturnsLong:           ps.MissedFieldGoalReturnsLong + other.MissedFieldGoalReturnsLong,
		MissedFieldGoalReturnsTouchdownsLong: ps.MissedFieldGoalReturnsTouchdownsLong + other.MissedFieldGoalReturnsTouchdownsLong,
		PuntReturns:                          ps.PuntReturns + other.PuntReturns,
		PuntReturnsYards:                     ps.PuntReturnsYards + other.PuntReturnsYards,
		PuntReturnsTouchdowns:                ps.PuntReturnsTouchdowns + other.PuntReturnsTouchdowns,
		PuntReturnsLong:                      ps.PuntReturnsLong + other.PuntReturnsLong,
		PuntReturnsTouchdownsLong:            ps.PuntReturnsTouchdownsLong + other.PuntReturnsTouchdownsLong,
		Kicks:                                ps.Kicks + other.Kicks,
		KickYards:                            ps.KickYards + other.KickYards,
		KicksNetYards:                        ps.KicksNetYards + other.KicksNetYards,
		KicksLong:                            ps.KicksLong + other.KicksLong,
		KicksSingles:                         ps.KicksSingles + other.KicksSingles,
		KicksOutOfEndZone:                    ps.KicksOutOfEndZone + other.KicksOutOfEndZone,
		KicksOnside:                          ps.KicksOnside + other.KicksOnside,
		OnePointConvertsAttempts:             ps.OnePointConvertsAttempts + other.OnePointConvertsAttempts,
		OnePointConvertsMade:                 ps.OnePointConvertsMade + other.OnePointConvertsMade,
		TwoPointConvertsMade:                 ps.TwoPointConvertsMade + other.TwoPointConvertsMade,
		TacklesTotal:                         ps.TacklesTotal + other.TacklesTotal,
		TacklesDefensive:                     ps.TacklesDefensive + other.TacklesDefensive,
		TacklesSpecialTeams:                  ps.TacklesSpecialTeams + other.TacklesSpecialTeams,
		SacksQbMade:                          ps.SacksQbMade + other.SacksQbMade,
		Interceptions:                        ps.Interceptions + other.Interceptions,
		InterceptionYards:                    ps.InterceptionYards + other.InterceptionYards,
		InterceptionTouchdowns:               ps.InterceptionTouchdowns + other.InterceptionTouchdowns,
		InterceptionLong:                     ps.InterceptionLong + other.InterceptionLong,
		InterceptionTouchdownsLong:           ps.InterceptionTouchdownsLong + other.InterceptionTouchdownsLong,
		FumblesForced:                        ps.FumblesForced + other.FumblesForced,
		FumblesRecovered:                     ps.FumblesRecovered + other.FumblesRecovered,
		PassesKnockedDown:                    ps.PassesKnockedDown + other.PassesKnockedDown,
	}
}
