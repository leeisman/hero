package hero

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"hero/controllers"
	"hero/database/ent"
	"hero/enums"
	userRepository "hero/repositories/user"
	userActiveRecordRepository "hero/repositories/user_active_record"
	"time"
)

type PlayRequest struct {
	FbUserID    string `json:"fb_user_id" form:"fb_user_id" query:"fb_user_id"`
	FBAvatarUrl string `json:"fb_avatar_url" form:"fb_avatar_url" query:"fb_avatar_url"`
	FbEmail     string `json:"fb_email" form:"fb_email" query:"fb_email"`
	FbName      string `json:"fb_name" form:"fb_name" query:"fb_name"`
}

type RecordRequest struct {
	RecordID string `json:"record_id"`
	Score    int    `json:"score"`
}

type TrackingRequest struct {
	FbUserID string `json:"fb_user_id" form:"fb_user_id" query:"fb_user_id"`
	FbName   string `json:"fb_name" form:"fb_name" query:"fb_name"`
	Type     string `json:"type"`
}

func Play(c echo.Context) error {
	ctx := context.Background()
	request := &PlayRequest{}
	now := time.Now().UTC()
	err := c.Bind(request)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	//找不到user也會當err,故不處理err
	user, _ := userRepository.FindBySocialUserID(ctx, request.FbUserID)
	if user == nil {
		newXID := xid.New()
		id := "UR_" + newXID.String()
		user, err = userRepository.Create(ctx, &ent.User{
			ID:              id,
			SocialUserID:    &request.FbUserID,
			SocialAvatarURL: &request.FBAvatarUrl,
			SocialEmail:     &request.FbEmail,
			SocialName:      &request.FbName,
			CreatedAt:       &now,
			UpdatedAt:       &now,
		})
		if err != nil {
			return controllers.ResponseFail(err, c)
		}
	}

	if user == nil {
		return controllers.ResponseFail(fmt.Errorf("Create user fail: %s", request.FbUserID), c)
	}

	userActiveRecord, err := userActiveRecordRepository.Create(ctx, &ent.UserActiveRecord{
		UserID:     &user.ID,
		ActiveType: enums.ActiveTypeHeroGame,
		StartedAt:  &now,
		CreatedAt:  &now,
		UpdatedAt:  &now,
	})
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	if userActiveRecord == nil {
		return controllers.ResponseFail(fmt.Errorf("Create user active record fail: %s", request.FbUserID), c)
	}
	return controllers.ResponseSuccess(userActiveRecord, c)
}

func Record(c echo.Context) error {
	ctx := context.Background()
	request := &RecordRequest{}
	err := c.Bind(request)
	userActiveRecord, err := userActiveRecordRepository.FindByID(ctx, request.RecordID)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	if userActiveRecord == nil {
		return controllers.ResponseFail(fmt.Errorf("userActiveRecord nil"), c)
	}
	if userActiveRecord.EndedAt != nil {
		return controllers.ResponseFail(fmt.Errorf("already finished"), c)
	}
	now := time.Now().UTC()
	userActiveRecord, err = userActiveRecord.Update().
		SetScore(request.Score).
		SetIsFinished(1).
		SetEndedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	return controllers.ResponseSuccess(userActiveRecord, c)
}

func Tracking(c echo.Context) error {
	ctx := context.Background()
	request := &TrackingRequest{}
	now := time.Now().UTC()
	err := c.Bind(request)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	user, err := userRepository.FindBySocialUserID(ctx, request.FbUserID)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	if request.Type != enums.ActiveTypeHeroShare && request.Type != enums.ActiveTypeHeroDownload {
		return controllers.ResponseFail(fmt.Errorf("type not invalid"), c)
	}

	userActiveRecord, err := userActiveRecordRepository.Create(ctx, &ent.UserActiveRecord{
		UserID:     &user.ID,
		ActiveType: request.Type,
		StartedAt:  &now,
		EndedAt:    &now,
		CreatedAt:  &now,
		UpdatedAt:  &now,
	})
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	return controllers.ResponseSuccess(userActiveRecord, c)
}
