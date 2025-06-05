package service

import (
	"context"
	"fmt"

	"github.com/orayew2002/go_motto/internal/domains"
	"github.com/orayew2002/go_motto/pkg/log"
)

type Service struct {
	repo   RepositoryInterface
	logger log.Logger
}

type RepositoryInterface interface {
	GetQuotes(ctx context.Context, author string) ([]domains.Quote, error)
	GetRandomQuote(ctx context.Context) (domains.Quote, error)
	GetQuotesByAuthor(ctx context.Context, author string) ([]domains.Quote, error)

	CreateQuote(ctx context.Context, quote domains.Quote) error
	DeleteQuote(ctx context.Context, id int) error
}

func NewService(repo RepositoryInterface, logger log.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) CreateQuote(ctx context.Context, quote domains.CreateQuoteRequest) error {
	err := s.repo.CreateQuote(ctx, domains.Quote{
		Author: quote.Author,
		Quote:  quote.Quote,
	})

	if err != nil {
		s.logger.Error("failed create quote", "details", fmt.Sprintf("request:%-v error:%-v", quote, err))
		return err
	}

	return nil
}

func (s *Service) DeleteQuote(ctx context.Context, id int) error {
	if err := s.repo.DeleteQuote(ctx, id); err != nil {
		s.logger.Error("failed delete quote", "details", fmt.Sprintf("id:  %d  error:%-v", id, err))
		return err
	}

	return nil
}

func (s *Service) GetQuotes(ctx context.Context, author string) ([]domains.QuoteResponse, error) {
	quotes, err := s.repo.GetQuotes(ctx, author)
	if err != nil {
		s.logger.Error("failed get quotes", "details", fmt.Sprintf("error:%-v", err))
		return nil, err
	}

	var quotesResponse = make([]domains.QuoteResponse, len(quotes))
	for i, quote := range quotes {
		quotesResponse[i] = domains.QuoteResponse{
			Id:     quote.Id,
			Author: quote.Author,
			Quote:  quote.Quote,
		}
	}

	return quotesResponse, nil
}

func (s *Service) GetQuotesByAuthor(ctx context.Context, author string) ([]domains.QuoteResponse, error) {
	quotes, err := s.repo.GetQuotesByAuthor(ctx, author)
	if err != nil {
		s.logger.Error("failed get quotes by author", "details", fmt.Sprintf("error:%-v", err))
		return nil, err
	}

	var quotesResponse = make([]domains.QuoteResponse, len(quotes))
	for i, quote := range quotes {
		quotesResponse[i] = domains.QuoteResponse{
			Author: quote.Author,
			Quote:  quote.Quote,
		}
	}

	return quotesResponse, nil
}

func (s *Service) GetRandomQuote(ctx context.Context) (domains.QuoteResponse, error) {
	quote, err := s.repo.GetRandomQuote(ctx)
	if err != nil {
		s.logger.Error("failed get random quotes", "details", fmt.Sprintf("error:%-v", err))
		return domains.QuoteResponse{}, err
	}

	return domains.QuoteResponse{
		Id:     quote.Id,
		Author: quote.Author,
		Quote:  quote.Quote,
	}, nil
}
