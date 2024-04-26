package domain

import (
	"time"
)

type InjuryDetails struct {
	StatusId    InjuryStatus `json:"status_id"`
	Text        string       `json:"text"`
	LastUpdated time.Time    `json:"last_updated"`
	Injury      string       `json:"injury"`
}
