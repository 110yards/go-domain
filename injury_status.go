package domain

type InjuryStatus string

const (
	InjuryStatusProbable     InjuryStatus = "probable"
	InjuryStatusQuestionable InjuryStatus = "questionable"
	InjuryStatusOut          InjuryStatus = "out"
	InjuryStatusInjuredSix   InjuryStatus = "six-game"
)
