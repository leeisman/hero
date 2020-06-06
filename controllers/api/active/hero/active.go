package hero

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"hero/controllers"
	"hero/database/ent"
	"hero/enums"
	"hero/pkg/logger"
	userRepository "hero/repositories/user"
	userActiveRecordRepository "hero/repositories/user_active_record"
	"net/http"
	"time"
)

// User
type HeroPlayRequest struct {
	FbUserID    string `json:"fb_user_id" form:"fb_user_id" query:"fb_user_id"`
	FBAvatarUrl string `json:"fb_avatar_url" form:"fb_avatar_url" query:"fb_avatar_url"`
	FbEmail     string `json:"fb_email" form:"fb_email" query:"fb_email"`
	FbName      string `json:"fb_name" form:"fb_name" query:"fb_name"`
}

func Play(c echo.Context) error {
	endPoint := "hero/play"
	ctx := context.Background()
	request := &HeroPlayRequest{}
	now := time.Now().UTC()
	err := c.Bind(request)
	logger.Print(endPoint, request)
	if err != nil {
		return controllers.ResponseFail(err, endPoint, c)
	}
	//找不到user也會當err,故不處理err
	user, _ := userRepository.FindUserBySocialUserID(ctx, request.FbUserID)
	if user == nil {
		newXID := xid.New()
		id := "UR_" + newXID.String()
		user, err = userRepository.CreateUser(ctx, &ent.User{
			ID:              id,
			SocialUserID:    &request.FbUserID,
			SocialAvatarURL: &request.FBAvatarUrl,
			SocialEmail:     &request.FbEmail,
			SocialName:      &request.FbName,
			CreatedAt:       &now,
			UpdatedAt:       &now,
		})
		if err != nil {
			return controllers.ResponseFail(err, endPoint, c)
		}
	}

	if user == nil {
		return controllers.ResponseFail(fmt.Errorf("Create user fail: %s", request.FbUserID), endPoint, c)
	}

	userActiveRecord, err := userActiveRecordRepository.Create(ctx, &ent.UserActiveRecord{
		UserID:     &user.ID,
		ActiveType: enums.ActiveTypeHeroGame,
		StartedAt:  &now,
		CreatedAt:  &now,
		UpdatedAt:  &now,
	})
	if err != nil {
		return controllers.ResponseFail(err, endPoint, c)
	}
	if userActiveRecord == nil {
		return controllers.ResponseFail(fmt.Errorf("Create user active record fail: %s", request.FbUserID), endPoint, c)
	}
	return controllers.ResponseSuccess(userActiveRecord, c)
}

func Record(c echo.Context) error {
	return c.String(http.StatusOK, "/Record")
}

func Share(c echo.Context) error {
	return c.String(http.StatusOK, "/share")
}

func Download(c echo.Context) error {
	return c.String(http.StatusOK, "/Download")
}
