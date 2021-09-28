package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DeNA/aelog"
	"github.com/go-chi/chi"
	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/viewmodel"
	"github.com/nonotakujet/memote-server/registry"
	"github.com/nonotakujet/memote-server/usecase"
	"github.com/thoas/go-funk"
)

type FixedRecordsHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type fixedRecordsHandler struct {
	usecase usecase.FixedRecordsUseCase
}

func NewFixedRecordsHandler(repo registry.Repository) FixedRecordsHandler {
	userFixedRecordRepo := repo.NewUserFixedRecordRepository()
	uc := usecase.NewFixedRecordUseCase(userFixedRecordRepo)
	return &fixedRecordsHandler{
		usecase: uc,
	}
}

func (p *fixedRecordsHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	recordId := chi.URLParam(r, "recordId")

	aelog.Infof(ctx, "%s\n", recordId)
	userFixedRecordModel, err := p.usecase.GetByRecordId(ctx, recordId)
	if err != nil {
		aelog.Errorf(ctx, "get fixed record failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fixedRecordViewModel := &viewmodel.FixedRecordViewModel{
		Id:               userFixedRecordModel.Id,
		MainTitle:        userFixedRecordModel.MainTitle,
		MainPicture:      userFixedRecordModel.MainPicture,
		IsPictureFetched: userFixedRecordModel.IsPictureFetched,
		Message:          userFixedRecordModel.Message,
		EmotionType:      userFixedRecordModel.EmotionType,
		EmotionLevel:     userFixedRecordModel.EmotionLevel,
		Locations: funk.Map(userFixedRecordModel.Locations, func(location model.UserFixedRecordLocation) viewmodel.StayedLocationViewModel {
			return viewmodel.StayedLocationViewModel{
				Name:      location.Name,
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
				Pictures:  location.Pictures,
				StartTime: location.StartTime,
				EndTime:   location.EndTime,
			}
		}).([]viewmodel.StayedLocationViewModel),
		LastRecommendedAt: userFixedRecordModel.LastRecommendedAt,
		CreatedAt:         userFixedRecordModel.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fixedRecordViewModel); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (p *fixedRecordsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	isFixedStr := r.URL.Query().Get("is_fixed")

	isPictureFetched, err := strconv.ParseBool(isFixedStr)
	if err != nil {
		aelog.Errorf(ctx, "%s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userFixedRecordModels, err := p.usecase.GetAllByPictureFecthedFlag(ctx, isPictureFetched)
	if err != nil {
		aelog.Errorf(ctx, "get all fixed records failed")
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
					Name:      location.Name,
					Latitude:  location.Latitude,
					Longitude: location.Longitude,
					Pictures:  location.Pictures,
					StartTime: location.StartTime,
					EndTime:   location.EndTime,
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

func (p *fixedRecordsHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	recordId := chi.URLParam(r, "recordId")

	aelog.Infof(ctx, "%s\n", recordId)

	fixedRecordViewModel := &viewmodel.FixedRecordViewModel{}
	if err := json.NewDecoder(r.Body).Decode(&fixedRecordViewModel); err != nil {
		aelog.Errorf(ctx, "parse request body failed : %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userFixedRecordModel := &model.UserFixedRecord{
		Id:               fixedRecordViewModel.Id,
		MainTitle:        fixedRecordViewModel.MainTitle,
		MainPicture:      fixedRecordViewModel.MainPicture,
		IsPictureFetched: fixedRecordViewModel.IsPictureFetched,
		Message:          fixedRecordViewModel.Message,
		EmotionType:      fixedRecordViewModel.EmotionType,
		EmotionLevel:     fixedRecordViewModel.EmotionLevel,
		Locations: funk.Map(fixedRecordViewModel.Locations, func(location viewmodel.StayedLocationViewModel) model.UserFixedRecordLocation {
			return model.UserFixedRecordLocation{
				Name:      location.Name,
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
				Pictures:  location.Pictures,
				StartTime: location.StartTime,
				EndTime:   location.EndTime,
			}
		}).([]model.UserFixedRecordLocation),
		LastRecommendedAt: fixedRecordViewModel.LastRecommendedAt,
		CreatedAt:         fixedRecordViewModel.CreatedAt,
	}

	_, err := p.usecase.Update(ctx, recordId, userFixedRecordModel)
	if err != nil {
		aelog.Errorf(ctx, "update fixed reocrd failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fixedRecordViewModel); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
