package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/orayew2002/go_motto/internal/domains"
	"github.com/orayew2002/go_motto/pkg/validator"
)

type handler struct {
	service ServiceInterface
}

type Handler interface {
	CreateQuote(w http.ResponseWriter, r *http.Request)
	GetQuotes(w http.ResponseWriter, r *http.Request)
	GetRandomQuote(w http.ResponseWriter, r *http.Request)
	DeleteQuote(w http.ResponseWriter, r *http.Request)
}

type ServiceInterface interface {
	GetQuotes(ctx context.Context, author string) ([]domains.QuoteResponse, error)
	GetRandomQuote(ctx context.Context) (domains.QuoteResponse, error)
	GetQuotesByAuthor(ctx context.Context, author string) ([]domains.QuoteResponse, error)
	DeleteQuote(ctx context.Context, id int) error
	CreateQuote(ctx context.Context, qoute domains.CreateQuoteRequest) error
}

func NewHandler(service ServiceInterface) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var req domains.CreateQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseBadRequest(w, err)
		return
	}

	if err := validator.ValidateRequiredFields(req); err != nil {
		responseBadRequest(w, err)
		return
	}

	if err := h.service.CreateQuote(r.Context(), req); err != nil {
		responseInternalServerError(w, err)
		return
	}

	responseSuccess(w, "quote successfully created")
}

func (h *handler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	quoteAuthor := r.URL.Query().Get("author")
	quotes, err := h.service.GetQuotes(r.Context(), quoteAuthor)
	if err != nil {
		responseInternalServerError(w, err)
		return
	}

	responseSuccess(w, quotes)
}

func (h *handler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	qoutes, err := h.service.GetRandomQuote(r.Context())
	if err != nil {
		responseInternalServerError(w, err)
		return
	}

	responseSuccess(w, qoutes)
}

func (h *handler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[1] != "quotes" || parts[2] == "" {
		responseBadRequest(w, errors.New("invalide url"))
		return
	}

	quoteID, err := strconv.Atoi(parts[2])
	if err != nil {
		responseBadRequest(w, err)
		return
	}

	if err := h.service.DeleteQuote(r.Context(), quoteID); err != nil {
		responseInternalServerError(w, err)
		return
	}

	responseSuccess(w, "quote successfully delted")
}
