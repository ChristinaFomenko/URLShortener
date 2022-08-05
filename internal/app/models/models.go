package models

type OriginalURL struct {
	CorrelationID string
	URL           string
}

type UserURL struct {
	UserID        string
	CorrelationID string
	ShortURL      string
	OriginalURL   string
	IsDeleted     bool `db:"is_deleted"`
}

type DeleteUserURLs struct {
	UserID string
	Short  string
}
