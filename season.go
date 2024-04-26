package domain

type SeasonMapping struct {
	Year       int    `json:"year"`
	SourceId   string `json:"sourceId"`
	SourceName string `json:"sourceName"`
}
