package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/nonotakujet/memote-server/registry"
	"github.com/nonotakujet/memote-server/usecase"
)

type PositionHandler interface {
	Post(w http.ResponseWriter, r *http.Request, pr httprouter.Params)
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

func (p *positionHandler) Post(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	type response struct {
		Latitude  int64     `json:"latitude"`
		Longitude int64     `json:"longitude"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	latitude, _ := strconv.Atoi(r.FormValue("la"))
	longitude, _ := strconv.Atoi(r.FormValue("lo"))
	userPositionModel := p.usecase.Post(int64(latitude), int64(longitude))

	//取得したドメインモデルをresponseに変換
	res := new(response)
	res.Latitude = userPositionModel.Latitude
	res.Longitude = userPositionModel.Longitude
	res.CreatedAt = userPositionModel.CreatedAt
	res.UpdatedAt = userPositionModel.UpdatedAt

	//クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	return
}
