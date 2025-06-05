package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/orayew2002/go_motto/internal/domains"
)

func TestApp(t *testing.T) {
	// Start the app in a goroutine
	go Run(domains.AppDependencies{
		Logger: domains.MockLogger{},
	})

	// Give the server a moment to start
	time.Sleep(500 * time.Millisecond)

	serverURL := "http://localhost:8080"

	// 1. POST /quotes
	t.Run("POST /quotes", func(t *testing.T) {
		postData := `{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}`
		_, err := http.Post(serverURL+"/quotes", "application/json", bytes.NewBufferString(postData))
		if err != nil {
			t.Fatalf("Failed to POST quote: %v", err)
		}
	})

	// 2. GET /quotes
	t.Run("GET /quotes", func(t *testing.T) {
		resp, err := http.Get(serverURL + "/quotes")
		if err != nil {
			t.Fatalf("Failed to GET /quotes: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		if !strings.Contains(string(body), "Confucius") {
			t.Errorf("Expected quote by Confucius in response")
		}
	})

	// 3. GET /quotes/random
	t.Run("GET /quotes/random", func(t *testing.T) {
		resp, err := http.Get(serverURL + "/quotes/random")
		if err != nil {
			t.Fatalf("Failed to GET /quotes/random: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
		}
	})

	// 4. GET /quotes?author=Confucius
	t.Run("GET /quotes?author=Confucius", func(t *testing.T) {
		resp, err := http.Get(serverURL + "/quotes?author=Confucius")
		if err != nil {
			t.Fatalf("Failed to GET /quotes?author=Confucius: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		if !strings.Contains(string(body), "Confucius") {
			t.Errorf("Expected filtered quote by Confucius")
		}
	})

	// 5. DELETE /quotes/{id}
	t.Run("DELETE /quotes/{id}", func(t *testing.T) {
		url := serverURL + fmt.Sprintf("/quotes/%d", 1)
		req, _ := http.NewRequest(http.MethodDelete, url, nil)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to DELETE quote: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
		}
	})
}
