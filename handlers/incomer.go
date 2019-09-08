package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/scukonick/eta/logger"
	"github.com/scukonick/eta/repos"
	dbStructs "github.com/scukonick/eta/repos/structs"
	"github.com/scukonick/eta/structs"
)

type Incomer struct {
	tasksRepo repos.Tasks
}

func NewIncomer(tasksRepo repos.Tasks) *Incomer {
	return &Incomer{
		tasksRepo: tasksRepo,
	}
}

func (s *Incomer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := &structs.ETARequest{}

	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	taskID := uuid.New().String()

	err = s.tasksRepo.Store(r.Context(), dbStructs.Task{
		ID:  taskID,
		Lat: data.Lng,
		Lng: data.Lng,
	})
	if err != nil {
		logger.FromContext(r.Context()).
			WithError(err).Error("failed to store task")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &structs.ETAResponse{
		ID: taskID,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(resp)
	if err != nil {
		logger.FromContext(r.Context()).
			WithError(err).Error("failed to encode and write response")
		return
	}
}
