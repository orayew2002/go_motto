package service

import (
	"context"
	"testing"

	"github.com/orayew2002/go_motto/internal/domains"
	"github.com/orayew2002/go_motto/internal/repository"
)

func TestServiceLayer(t *testing.T) {
	service := NewService(repository.NewRepository(), domains.MockLogger{})
	ctx := context.Background()

	t.Run("create quotes", func(t *testing.T) {
		for _, quote := range domains.MockQuoteRequests {
			if err := service.CreateQuote(ctx, quote); err != nil {
				t.Fatalf("failed to create quote by %s: %v", quote.Author, err)
			}
		}

		t.Log("All quotes created successfully")
	})

	t.Run("get all quotes", func(t *testing.T) {
		quotes, err := service.GetQuotes(ctx, "")
		if err != nil {
			t.Fatalf("failed to get quotes: %v", err)
		}

		if len(quotes) == 0 {
			t.Error("expected quotes, got none")
		}

		t.Logf("Retrieved %d quotes", len(quotes))
	})

	t.Run("get random quote", func(t *testing.T) {
		quote, err := service.GetRandomQuote(ctx)
		if err != nil {
			t.Fatalf("failed to get random quote: %v", err)
		}

		t.Logf("Random quote retrieved: %+v", quote)
	})

	t.Run("get quotes by author", func(t *testing.T) {
		author := domains.MockQuoteRequests[0].Author
		quotes, err := service.GetQuotesByAuthor(ctx, author)
		if err != nil {
			t.Fatalf("failed to get quotes by author %s: %v", author, err)
		}

		if len(quotes) == 0 {
			t.Errorf("expected quotes for author %s, got none", author)
		}

		t.Logf("Retrieved %d quotes by author %s", len(quotes), author)
	})

	t.Run("delete quote", func(t *testing.T) {
		quoteID := 1
		if err := service.DeleteQuote(ctx, quoteID); err != nil {
			t.Fatalf("failed to delete quote with ID %d: %v", quoteID, err)
		}

		t.Logf("Quote with ID %d deleted successfully", quoteID)
	})
}
