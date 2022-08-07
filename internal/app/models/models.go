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
	DeletedAt     bool `db:"deleted_at"`
}

type DeleteUserURLs struct {
	UserID string
	Short  string
}
