package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/orayew2002/go_motto/internal/domains"
)

func responseInternalServerError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(domains.Reponse{
		Status: "error",
		Error:  err.Error(),
	})
}

func responseBadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(domains.Reponse{
		Status: "error",
		Error:  err.Error(),
	})
}

func responseSuccess(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(domains.Reponse{
		Status: "success",
		Data:   response,
	})
}
