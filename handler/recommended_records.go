package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DeNA/aelog"
	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/viewmodel"
	"github.com/nonotakujet/memote-server/registry"
	"github.com/nonotakujet/memote-server/usecase"
	"github.com/thoas/go-funk"
)

type RecommendedRecordsHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type recommendedRecordsHandler struct {
	usecase usecase.RecommendedRecordsUseCase
}

func NewRecommendedRecordsHandler(repo registry.Repository) RecommendedRecordsHandler {
	userFixedRecordRepo := repo.NewUserFixedRecordRepository()
	userLocationRepo := repo.NewUserLocationRepository()
	uc := usecase.NewRecommendedRecordUseCase(userFixedRecordRepo, userLocationRepo)
	return &recommendedRecordsHandler{
		usecase: uc,
	}
}

func (p *recommendedRecordsHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	latitudeStr := r.URL.Query().Get("latitude")
	longitudeStr := r.URL.Query().Get("longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		aelog.Errorf(ctx, "%s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		aelog.Errorf(ctx, "%s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userFixedRecordModels, err := p.usecase.Get(ctx, latitude, longitude)
	if err != nil {
		aelog.Errorf(ctx, "post reocrd failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fixedRecordViewModels := make([]*viewmodel.FixedRecordViewModel, len(userFixedRecordModels))

	for i := range fixedRecordViewModels {
		userFixedRecordModel := userFixedRecordModels[i]
		fixedRecordViewModels[i] = &viewmodel.FixedRecordViewModel{
			Id:               userFixedRecordModel.Id,
			MainTitle:        userFixedRecordModel.MainTitle,
			MainPicture:      userFixedRecordModel.MainPicture,
			IsPictureFetched: userFixedRecordModel.IsPictureFetched,
			Message:          userFixedRecordModel.Message,
			EmotionType:      userFixedRecordModel.EmotionType,
			EmotionLevel:     userFixedRecordModel.EmotionLevel,
			Locations: funk.Map(userFixedRecordModel.Locations, func(location model.UserFixedRecordLocation) viewmodel.StayedLocationViewModel {
				return viewmodel.StayedLocationViewModel{
					Name:         location.Name,
					Latitude:     location.Latitude,
					Longitude:    location.Longitude,
					Pictures:     location.Pictures,
					StartTime:    location.StartTime,
					EndTime:      location.EndTime,
					Message:      location.Message,
					EmotionType:  0,
					EmotionLevel: 0,
				}
			}).([]viewmodel.StayedLocationViewModel),
			LastRecommendedAt: userFixedRecordModel.LastRecommendedAt,
			CreatedAt:         userFixedRecordModel.CreatedAt,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fixedRecordViewModels); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
