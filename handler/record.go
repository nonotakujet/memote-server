package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DeNA/aelog"
	"github.com/nonotakujet/memote-server/domain/viewmodel"
	"github.com/nonotakujet/memote-server/registry"
	"github.com/nonotakujet/memote-server/usecase"
)

type RecordHandler interface {
	Post(w http.ResponseWriter, r *http.Request)
}

type recordHandler struct {
	usecase usecase.RecordUseCase
}

func NewRecordHandler(repo registry.Repository) RecordHandler {
	userRecordRepo := repo.NewUserRecordRepository()
	uc := usecase.NewRecordUseCase(userRecordRepo)
	return &recordHandler{
		usecase: uc,
	}
}

func (p *recordHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	recordViewModel := &viewmodel.RecordViewModel{}
	if err := json.NewDecoder(r.Body).Decode(&recordViewModel); err != nil {
		aelog.Errorf(ctx, "parse request body failed : %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	_, err := p.usecase.Post(ctx, recordViewModel)
	if err != nil {
		aelog.Errorf(ctx, "post reocrd failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// success.
	successViewModel := viewmodel.SuccessViewModel{
		Success: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(successViewModel); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
