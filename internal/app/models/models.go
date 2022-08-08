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

func NewURL(origin, short string) *UserURL {
	return &UserURL{OriginalURL: origin,
		ShortURL: short,
	}
}

type DeleteUserURLs struct {
	UserID string
	Short  string
}
