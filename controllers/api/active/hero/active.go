package hero

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"hero/controllers"
	"hero/database/ent"
	"hero/enums"
	"hero/pkg/db/mysql"
	"hero/pkg/db/redis"
	"hero/pkg/logger"
	prizeRepository "hero/repositories/prize"
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
	FbUserID string `json:"fb_user_id" form:"fb_user_id" query:"fb_user_id"`
	RecordID string `json:"record_id"`
	Score    int    `json:"score"`
}

type TrackingRequest struct {
	FbUserID string `json:"fb_user_id" form:"fb_user_id" query:"fb_user_id"`
	FbName   string `json:"fb_name" form:"fb_name" query:"fb_name"`
	Type     string `json:"type"`
}

type PrizeRequest struct {
	FbUserID string `json:"fb_user_id" form:"fb_user_id" query:"fb_user_id"`
}

func Play(c echo.Context) error {
	ctx := context.Background()
	client, err := mysql.Client().Tx(ctx)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	request := &PlayRequest{}
	now := time.Now().UTC()
	err = c.Bind(request)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	//找不到user也會當err,故不處理err
	user, err := userRepository.FindBySocialUserID(ctx, request.FbUserID)
	if user == nil {
		logger.Printf("not found user id ", request.FbUserID)
		if err != nil {
			logger.Printf("err ", err.Error())
		}
		newXID := xid.New()
		id := "UR_" + newXID.String()
		user, err = userRepository.Create(ctx, client, &ent.User{
			ID:              id,
			SocialUserID:    request.FbUserID,
			SocialAvatarURL: request.FBAvatarUrl,
			SocialType:      enums.SocialTypeFacebook,
			SocialEmail:     request.FbEmail,
			SocialName:      request.FbName,
			CreatedAt:       now,
			UpdatedAt:       now,
		})
		if err != nil {
			mysql.Rollback(client, err)
			return controllers.ResponseFail(err, c)
		}
	}

	if user == nil {
		return controllers.ResponseFail(fmt.Errorf("Create user fail: %s", request.FbUserID), c)
	}
	userActiveRecord, err := userActiveRecordRepository.Create(ctx, &ent.UserActiveRecord{
		UserID:     user.ID,
		ActiveType: enums.ActiveTypeHeroGame,
		StartedAt:  &now,
		CreatedAt:  &now,
		UpdatedAt:  &now,
	})
	if err != nil {
		mysql.Rollback(client, err)
		return controllers.ResponseFail(err, c)
	}
	if userActiveRecord == nil {
		mysql.Rollback(client, err)
		return controllers.ResponseFail(fmt.Errorf("Create user active record fail: %s", request.FbUserID), c)
	}
	countFinishedOrUnfinishedUserTotal(ctx, 0)
	err = client.Commit()
	if err != nil {
		logger.Print("play commit err ", err.Error())
	}
	return controllers.ResponseSuccess(userActiveRecord, c)
}

func Record(c echo.Context) error {
	ctx := context.Background()
	client, err := mysql.Client().Tx(ctx)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	request := &RecordRequest{}
	err = c.Bind(request)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
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
	user, _ := userRepository.FindBySocialUserID(ctx, request.FbUserID)
	if user == nil {
		return controllers.ResponseFail(fmt.Errorf("User no found: %s", request.FbUserID), c)
	}
	repeatStatus := user.HeroRepeat
	if userActiveRecord.UserID != user.ID {
		return controllers.ResponseFail(fmt.Errorf("Record no match user: %s", request.FbUserID), c)
	}
	queryBuilder := user.Update().SetHeroPlayed(1)
	if user.HeroPlayed == 1 {
		queryBuilder.SetHeroRepeat(1)
	}
	_, err = queryBuilder.Save(ctx)
	if err != nil {
		mysql.Rollback(client, err)
		return controllers.ResponseFail(fmt.Errorf("Change user repeat type err: %s", request.FbUserID), c)
	}
	now := time.Now().UTC()
	userActiveRecord, err = userActiveRecord.Update().
		SetScore(request.Score).
		SetIsFinished(1).
		SetEndedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		mysql.Rollback(client, err)
		return controllers.ResponseFail(err, c)
	}

	//更新user score
	queryBuilder = user.Update().SetLatestHeroScore(request.Score)
	if request.Score > user.BetterHeroScore {
		queryBuilder = queryBuilder.SetBetterHeroScore(request.Score)
	}
	_, err = queryBuilder.Save(ctx)
	if err != nil {
		mysql.Rollback(client, err)
		return controllers.ResponseFail(err, c)
	}

	countFinishedOrUnfinishedUserTotal(ctx, 1)
	countRepeatOrNotRepeatUserTotal(ctx, repeatStatus, user.ID)
	err = client.Commit()
	if err != nil {
		logger.Print("play commit err ", err.Error())
	}
	return controllers.ResponseSuccess(userActiveRecord, c)
}

func Tracking(c echo.Context) error {
	ctx := context.Background()
	client, err := mysql.Client().Tx(ctx)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	request := &TrackingRequest{}
	now := time.Now().UTC()
	err = c.Bind(request)
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
		UserID:     user.ID,
		ActiveType: request.Type,
		StartedAt:  &now,
		EndedAt:    &now,
		CreatedAt:  &now,
		UpdatedAt:  &now,
	})
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	err = client.Commit()
	if err != nil {
		logger.Print("play commit err ", err.Error())
	}
	return controllers.ResponseSuccess(userActiveRecord, c)
}

func Prize(c echo.Context) error {
	ctx := context.Background()
	request := &PrizeRequest{}
	now := time.Now().UTC().Format("2006-01-02")
	logger.Print("now", now)
	err := c.Bind(request)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	prize, _ := prizeRepository.FindBySocialUserIDAndDate(ctx, request.FbUserID, now)
	if prize != nil {
		return controllers.ResponseFail(fmt.Errorf("User has recorded"), c)
	}
	prize, err = prizeRepository.Create(ctx, request.FbUserID, now)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	return controllers.ResponseSuccess(prize, c)
}

func countFinishedOrUnfinishedUserTotal(ctx context.Context, isFinished uint) {
	if isFinished == 1 {
		finished := redis.Client().Incr(ctx, enums.RedisFinishedGameCount)
		unfinished := redis.Client().Decr(ctx, enums.RedisUnfinishedGameCount)
		logger.Print(enums.RedisFinishedGameCount, finished.String())
		logger.Print(enums.RedisUnfinishedGameCount, unfinished.String())
	} else {
		result := redis.Client().Incr(ctx, enums.RedisUnfinishedGameCount)
		logger.Print(enums.RedisFinishedGameCount, result.String())
	}
}

func countRepeatOrNotRepeatUserTotal(ctx context.Context, repeatStatus uint, userID string) {
	if repeatStatus == 1 {
		redis.Client().HSet(ctx, enums.RedisRepeatUserCount, userID, 1)
		redis.Client().HDel(ctx, enums.RedisNotRepeatUserCount, userID)
	} else {
		redis.Client().HSet(ctx, enums.RedisNotRepeatUserCount, userID, 1)
	}
	logger.Print(enums.RedisRepeatUserCount, redis.Client().HLen(ctx, enums.RedisRepeatUserCount).String())
	logger.Print(enums.RedisNotRepeatUserCount, redis.Client().HLen(ctx, enums.RedisNotRepeatUserCount).String())
}
