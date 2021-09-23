package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nonotakujet/memote-server/domain/viewmodel"
	"github.com/nonotakujet/memote-server/registry"
	"github.com/nonotakujet/memote-server/usecase"
)

type PositionHandler interface {
	Post(w http.ResponseWriter, r *http.Request)
}

type positionHandler struct {
	usecase usecase.PositionUseCase
}

func NewPositionHandler(repo registry.Repository) PositionHandler {
	userPositionRepo := repo.NewUserPositionRepository()
	uc := usecase.NewPositionUseCase(userPositionRepo)
	return &positionHandler{
		usecase: uc,
	}
}

func (p *positionHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	locationViewModel := &[]viewmodel.LocationViewModel{}
	if err := json.NewDecoder(r.Body).Decode(&locationViewModel); err != nil {
		log.Fatalf("parse request body failed : %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userPositionModel := p.usecase.Post(ctx, *locationViewModel)
	if userPositionModel != nil {

	}

	successViewModel := viewmodel.SuccessViewModel{
		Success: "success",
	}

	//クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(successViewModel); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
