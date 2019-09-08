package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/scukonick/eta/logger"
	"github.com/scukonick/eta/repos"
	dbStructs "github.com/scukonick/eta/repos/structs"
	"github.com/scukonick/eta/structs"
)

type Resulter struct {
	r repos.Results
}

func NewResulter(r repos.Results) *Resulter {
	return &Resulter{
		r: r,
	}
}

func (r *Resulter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := r.r.Get(req.Context(), id)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		logger.FromContext(req.Context()).WithError(err).
			Error("failed to lookup result")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var status string
	switch result.Status {
	case dbStructs.StatusProcessed:
		status = "done"
	case dbStructs.StatusNotProcessed:
		status = "pending"
	default:
		logger.FromContext(req.Context()).
			WithField("status", result.Status).Error("invalid status")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&structs.ETAResponse{
		Status: status,
		ID:     id,
		ETA:    result.ETA.Int64,
	})
	if err != nil {
		logger.FromContext(req.Context()).WithError(err).
			Error("failed to marshal json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		logger.FromContext(req.Context()).WithError(err).
			Error("failed to write response")
		return
	}
}
