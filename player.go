package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"110yards.ca/libs/go/core/logger"
)

type Player struct {
	PlayerId             string         `json:"player_id"`
	DateUpdated          time.Time      `json:"date_updated"`
	FirstName            string         `json:"first_name"`
	LastName             string         `json:"last_name"`
	FullName             string         `json:"full_name"`
	BirthDate            time.Time      `json:"birth_date"`
	BirthPlace           string         `json:"birth_place"`
	Height               string         `json:"height"`
	Weight               int            `json:"weight"`
	CanadianPlayer       bool           `json:"canadian_player"`
	Position             Position       `json:"position"`
	Team                 Team           `json:"team"`
	IsFreeAgent          bool           `json:"is_free_agent"`
	AlternateComputedIds []string       `json:"alternate_computed_ids"`
	Uniform              string         `json:"uniform"`
	School               string         `json:"school"`
	Seasons              []int          `json:"seasons"`
	InjuryStatus         *InjuryDetails `json:"injury_status"`
	SourceName           string         `json:"source_name"`
	SourceId             string         `json:"source_id"`
	Hash                 string         `json:"hash"`
}

type PlayerStub struct {
	PlayerId             string    `json:"player_id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	BirthDate            time.Time `json:"birth_date"`
	IsFreeAgent          bool      `json:"is_free_agent"`
	LastUpdated          time.Time `json:"last_updated"`
	SourceName           string    `json:"source_name"`
	SourceId             string    `json:"source_id"`
	AlternateComputedIds []string  `json:"alternate_computed_ids"`
	Hash                 string    `json:"hash"`
}

func CreatePlayer(
	firstName, lastName string,
	birthDate time.Time, birthplace string,
	height string, weight int,
	canadianPlayer bool,
	position Position,
	team Team,
	uniform string,
	school string,
	sourceName string,
	sourceId string,
	season int,
	injuryDetails *InjuryDetails,
) (Player, error) {

	playerId, err := computeId(firstName, lastName, birthDate)

	if err != nil {
		return Player{}, err
	}

	player := Player{
		PlayerId:             playerId,
		DateUpdated:          time.Now(),
		FirstName:            firstName,
		LastName:             lastName,
		FullName:             fmt.Sprintf("%s %s", firstName, lastName),
		BirthDate:            birthDate,
		BirthPlace:           birthplace,
		Height:               height,
		Weight:               weight,
		CanadianPlayer:       canadianPlayer,
		Position:             position,
		Team:                 team,
		IsFreeAgent:          team.IsFreeAgent(),
		AlternateComputedIds: []string{},
		Uniform:              uniform,
		School:               school,
		Seasons:              []int{season},
		InjuryStatus:         injuryDetails,
		SourceName:           sourceName,
		SourceId:             sourceId,
	}

	player.UpdateHash()

	return player, nil
}

func computeId(firstName, lastName string, birthDate time.Time) (string, error) {
	if firstName == "" {
		return "", errors.New("first name cannot be empty")
	}

	if lastName == "" {
		return "", errors.New("last name cannot be empty")
	}

	if birthDate.IsZero() {
		return "", errors.New("birth date cannot be empty")
	}

	birthDateUnix := int(birthDate.Unix())
	nameSlug := nameSlug(firstName, lastName)

	idClear := fmt.Sprintf("%s-%d", nameSlug, birthDateUnix)

	return calculateHash(idClear), nil
}

func (p *Player) UpdateHash() error {
	// don't include last update time in hash calculation
	copy := *p
	copy.DateUpdated = time.Time{}

	// marshal p to json
	j, err := json.Marshal(copy)

	if err != nil {
		logger.Errorf("Error while marshalling player: %s", err.Error())
		return err
	}

	// calculate hash
	p.Hash = calculateHash(string(j))
	return nil
}
