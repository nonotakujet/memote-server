package persistence

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/DeNA/aelog"
	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/domain/repository"
)

// UserLocationRepository holds user position inteface
type UserLocationRepository struct {
	db *DB
}

// NewUserRecordRepository new user position
func NewUserLocationRepository(db *DB) repository.UserLocation {
	newRepo := &UserLocationRepository{
		db: db,
	}
	return newRepo
}

// get All UserRecord
func (r *UserLocationRepository) GetNearBy(ctx context.Context, uid *model.UID, geohash string) ([]*model.UserLocation, error) {
	/*
		桁数	南北の距離	東西の距離
		1	4,989,600.00m	4,050,000.00m
		2	623,700.00m	1,012,500.00m
		3	155,925.00m	126,562.50m
		4	19,490.62m	31,640.62m
		5	4,872.66m	3,955.08m
		6	609.08m	988.77m
		7	152.27m	123.60m
		8	19.03m	30.90m
		9	4.76m	3.86m
		10	0.59m	0.97m
	*/

	// まずは、8桁まで一致しているlocationを抽出.
	startAt := geohash[:8] + "00"
	endAt := geohash[:8] + "zz"
	dss, err := r.db.client.Collection("users").Doc(uid.ID).Collection("locations").OrderBy("geohash", firestore.Asc).StartAt(startAt).EndAt(endAt).Documents(ctx).GetAll()

	if err != nil {
		aelog.Errorf(ctx, "failed query locations: %v", err)
		return nil, err
	}

	// まずは、7桁まで一致しているlocationを抽出.
	if len(dss) == 0 {
		startAt = geohash[:7] + "000"
		endAt = geohash[:7] + "zzz"
		dss, err = r.db.client.Collection("users").Doc(uid.ID).Collection("locations").OrderBy("geohash", firestore.Asc).StartAt(startAt).EndAt(endAt).Documents(ctx).GetAll()

		if err != nil {
			aelog.Errorf(ctx, "failed query locations: %v", err)
			return nil, err
		}
	}

	// 7桁の結果がなければ、6桁まで一致しているものを抽出.
	if len(dss) == 0 {
		startAt = geohash[:6] + "0000"
		endAt = geohash[:6] + "zzzz"
		dss, err = r.db.client.Collection("users").Doc(uid.ID).Collection("locations").OrderBy("geohash", firestore.Asc).StartAt(startAt).EndAt(endAt).Documents(ctx).GetAll()

		if err != nil {
			aelog.Errorf(ctx, "failed query locations: %v", err)
			return nil, err
		}
	}

	// 6桁の結果がなければ、5桁まで一致しているものも抽出.
	if len(dss) == 0 {
		startAt = geohash[:5] + "00000"
		endAt = geohash[:5] + "zzzzz"
		dss, err = r.db.client.Collection("users").Doc(uid.ID).Collection("locations").OrderBy("geohash", firestore.Asc).StartAt(startAt).EndAt(endAt).Documents(ctx).GetAll()
		if err != nil {
			aelog.Errorf(ctx, "failed query locations: %v", err)
			return nil, err
		}
	}

	userLocaions := make([]*model.UserLocation, len(dss))
	for i, ss := range dss {
		var userLocation = model.UserLocation{}
		if err := ss.DataTo(&userLocation); err != nil {
			aelog.Errorf(ctx, "userLocation parse error : %v", err)
			return nil, err
		}
		userLocaions[i] = &userLocation
	}

	return userLocaions, nil
}
