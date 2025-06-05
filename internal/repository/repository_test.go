package repository

import (
	"context"
	"testing"

	"github.com/orayew2002/go_motto/internal/domains"
)

func TestRepositoryLayer(t *testing.T) {
	repository := NewRepository()
	ctx := context.Background()

	t.Run("create quote", func(t *testing.T) {
		for _, mockQuote := range domains.MockQuoteRequests {
			err := repository.CreateQuote(ctx, domains.Quote{
				Author: mockQuote.Author,
				Quote:  mockQuote.Quote,
			})

			if err != nil {
				t.Errorf("can't create quote %+v", err)
			}
		}

		t.Log("all quotes success created")
	})

	t.Run("get random quote", func(t *testing.T) {
		quote, err := repository.GetRandomQuote(ctx)
		if err != nil {
			t.Errorf("can't get random quote %+v", err)
		}

		t.Logf("random quote %+v", quote)
	})

	t.Run("get quote by Author", func(t *testing.T) {
		quote := domains.MockQuoteRequests[0]

		quotes, err := repository.GetQuotesByAuthor(ctx, quote.Author)
		if err != nil {
			t.Errorf("can't get quote by author %+v", err)
		}

		t.Logf("get quotes by auhtor %+v", quotes)
	})

	t.Run("delete quote by id", func(t *testing.T) {
		quoteId := 1

		if err := repository.DeleteQuote(ctx, quoteId); err != nil {
			t.Errorf("error deleting error %+v", err)
		}

		t.Logf("quote success deleted")
	})

	t.Run("get quotes", func(t *testing.T) {
		quotes, err := repository.GetQuotes(ctx, "")
		if err != nil {
			t.Errorf("error gettting quotes %+v", err)
		}

		t.Logf("quotes %+v", quotes)
	})
}
