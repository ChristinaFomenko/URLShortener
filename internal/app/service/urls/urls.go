package urls

import (
	"context"
	"errors"
	"fmt"
	"github.com/ChristinaFomenko/shortener/internal/app/models"
	errs "github.com/ChristinaFomenko/shortener/pkg/errors"
	_ "github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -source=urls.go -destination=mocks/mocks.go

const (
	idLength int64 = 5
)

type urlRepository interface {
	Add(ctx context.Context, urlID, url, userID string) error
	Get(ctx context.Context, urlID string) (string, error)
	FetchURLs(ctx context.Context, userID string) ([]models.UserURL, error)
	AddBatch(ctx context.Context, urls []models.UserURL, userID string) error
	DeleteUserURLs(ctx context.Context, userID string, toDelete []string) error
}

type generator interface {
	Letters(n int64) (string, error)
}

type service struct {
	repository urlRepository
	generator  generator
	host       string
	//deletionChan chan models.DeleteUserURLs
	//buf          []models.DeleteUserURLs
	//timer        *time.Timer
	//isTimeout    bool
}

func NewService(repository urlRepository, generator generator, host string) *service {
	s := service{
		repository: repository,
		generator:  generator,
		host:       host,
		//deletionChan: make(chan models.DeleteUserURLs),
		//buf:          make([]models.DeleteUserURLs, 0, utils.BufLen),
		//isTimeout:    true,
		//timer:        time.NewTimer(0),
	}

	//go s.worker()

	return &s
}

func (s *service) Shorten(ctx context.Context, url, userID string) (string, error) {
	urlID, err := s.generator.Letters(idLength)
	if err != nil {
		log.WithError(err).
			WithField("userID", userID).
			WithField("url", url).Error("add url error")
		return "", err
	}

	if err = s.repository.Add(ctx, urlID, url, userID); err != nil {
		var uniqueErr *errs.NotUniqueURLErr
		if errors.As(err, &uniqueErr) {
			return s.buildShortURL(uniqueErr.URLID), errs.ErrNotUniqueURL
		}

		log.WithError(err).
			WithField("userID", userID).
			WithField("urlID", urlID).
			WithField("url", url).
			Error("add url error")
		return "", err

	}

	return s.buildShortURL(urlID), nil
}

// Return by id

func (s *service) Expand(ctx context.Context, urlID string) (string, error) {
	url, err := s.repository.Get(ctx, urlID)
	if err != nil {
		if errors.Is(err, errs.ErrURLNotFound) {
			return "", errs.ErrURLNotFound
		}
		log.WithError(err).WithField("urlID", urlID).Error("get url error")
		return "", err
	}

	return url, nil
}

func (s *service) FetchURLs(ctx context.Context, userID string) ([]models.UserURL, error) {
	urls, err := s.repository.FetchURLs(ctx, userID)
	if err != nil {
		log.WithError(err).WithField("urlID", userID).Error("get url list error")
		return nil, err
	}

	for idx := range urls {
		urls[idx].ShortURL = s.buildShortURL(urls[idx].ShortURL)
	}

	return urls, nil
}

func (s *service) ShortenBatch(ctx context.Context, originalURLs []models.OriginalURL, userID string) ([]models.UserURL, error) {
	urls := make([]models.UserURL, len(originalURLs))
	for idx := range urls {
		urlID, err := s.generator.Letters(idLength)
		if err != nil {
			log.WithError(err).
				WithField("userID", userID).
				WithField("originalURLs", originalURLs).
				Error("generate urlID error")
			return nil, err
		}
		urls[idx] = models.UserURL{
			CorrelationID: originalURLs[idx].CorrelationID,
			ShortURL:      urlID,
			OriginalURL:   originalURLs[idx].URL,
		}
	}

	err := s.repository.AddBatch(ctx, urls, userID)
	if err != nil {
		log.WithError(err).
			WithField("userID", userID).
			WithField("originalURLs", originalURLs).
			WithField("urls", urls).
			Error("add urls batch error")
		return nil, err
	}

	for idx := range urls {
		urls[idx].ShortURL = s.buildShortURL(urls[idx].ShortURL)
	}

	return urls, nil
}

func (s *service) buildShortURL(id string) string {
	return fmt.Sprintf("%s/%s", s.host, id)
}

func (s *service) DeleteUserURLs(ctx context.Context, userID string, toDelete []string) error {
	return s.repository.DeleteUserURLs(ctx, userID, toDelete)
}

//func (s *service) flush(ctx context.Context) {
//	del := make([]models.DeleteUserURLs, len(s.buf))
//	copy(del, s.buf)
//	s.buf = make([]models.DeleteUserURLs, 0)
//	go func() {
//		err := s.repository.DeleteUserURLs(ctx, del)
//		if err != nil {
//			log.Printf("error deleting: " + err.Error())
//		}
//	}()
//}
//
//func (s *service) worker() {
//	ctx := context.Background()
//
//	for {
//		select {
//		case delRequest := <-s.deletionChan:
//			if s.isTimeout {
//				s.timer.Reset(time.Second * utils.Timeout)
//				s.isTimeout = false
//			}
//			s.buf = append(s.buf, delRequest)
//			if len(s.buf) >= utils.BufLen {
//				s.flush(ctx)
//				s.timer.Stop()
//				s.isTimeout = true
//			}
//		case <-s.timer.C:
//			if len(s.buf) > 0 {
//				s.flush(ctx)
//			}
//			s.isTimeout = true
//		}
//	}
//}
