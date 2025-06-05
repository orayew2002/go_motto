package repository

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/orayew2002/go_motto/internal/domains"
)

var quotesLastId int

type Repository struct {
	mu     sync.RWMutex
	quotes []domains.Quote
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateQuote(ctx context.Context, q domains.Quote) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	quotesLastId++
	q.Id = quotesLastId

	r.quotes = append(r.quotes, q)
	return nil
}

func (r *Repository) DeleteQuote(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, q := range r.quotes {
		if q.Id == id {
			r.quotes = append(r.quotes[:i], r.quotes[i+1:]...)
			break
		}
	}

	return nil
}

func (r *Repository) GetQuotesByAuthor(ctx context.Context, author string) ([]domains.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var quotes []domains.Quote
	for _, q := range r.quotes {
		if q.Author == author {
			quotes = append(quotes, q)
		}
	}

	return quotes, nil
}

func (r *Repository) GetQuotes(ctx context.Context, author string) ([]domains.Quote, error) {
	if len(author) != 0 {
		return r.GetQuotesByAuthor(ctx, author)
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	quotesCopy := make([]domains.Quote, len(r.quotes))
	copy(quotesCopy, r.quotes)

	return quotesCopy, nil
}

func (r *Repository) GetRandomQuote(ctx context.Context) (domains.Quote, error) {
	if len(r.quotes) == 0 {
		return domains.Quote{}, nil
	}

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(r.quotes))

	return r.quotes[idx], nil
}
