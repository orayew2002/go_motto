package handlers

import (
	"net/http"
)

func Routes(handler Handler) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(post("/quotes"), handler.CreateQuote)
	r.HandleFunc(get("/quotes"), handler.GetQuotes)
	r.HandleFunc(get("/quotes/random"), handler.GetRandomQuote)
	r.HandleFunc(delete("/quotes/{id}"), handler.DeleteQuote)

	return r
}

func post(path string) string {
	return "POST " + path
}

func get(path string) string {
	return "GET " + path
}

func delete(path string) string {
	return "DELETE " + path
}
