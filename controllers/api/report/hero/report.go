package hero

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"hero/controllers"
	"hero/database/ent"
	"hero/enums"
	"hero/pkg/logger"
	"hero/repositories/user"
	"hero/repositories/user_active_record"
	"strconv"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

type UserCountResponse struct {
	FinishedUserCount   int               `json:"finished_user_count"`
	RepeatUserCount     int               `json:"repeat_user_count"`
	NotRepeatUserCount  int               `json:"not_repeat_user_count"`
	UnfinishedUserCount int               `json:"unfinished_user_count"`
	PlayGameCount       int               `json:"play_game_count"`
	ShareCount          int               `json:"share_count"`
	DownloadGameCount   int               `json:"download_game_count"`
	StartAt             time.Time         `json:"start_at"`
	EndAt               time.Time         `json:"end_at"`
	Node                map[string]string `json:"node"`
}

type RankResponse struct {
	Rank []*subRankResponse `json:"rank"`
	ME   *subRankResponse   `json:"me"`
}

type subRankResponse struct {
	LatestHeroScore int     `json:"latest_hero_score"`
	BetterHeroScore int     `json:"better_hero_score"`
	SocialName      string  `json:"social_name"`
	SocialEmail     string  `json:"social_email"`
	SocialAvatarURL string  `json:"social_avatar_url"`
	Percentage      float32 `json:"percentage"`
}

func UserCount(c echo.Context) error {

	ctx := context.Background()
	//轉換時間
	startT, err := time.Parse(TimeLayout, c.QueryParam("start_at"))
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	endT, err := time.Parse(TimeLayout, c.QueryParam("end_at"))
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	// 重複人數
	repeatUserCount, err := user.CountRepeatUser(ctx, 1, startT, endT)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	// 不重複人數
	notRepeatUserCount, err := user.CountRepeatUser(ctx, 0, startT, endT)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}

	// 沒玩完人數
	unfinishedUserCount, err := user.CountUnFinishedUser(ctx, startT, endT)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}

	// 開始遊戲次數
	playGameCount, err := user_active_record.CountRecord(ctx, enums.ActiveTypeHeroGame, startT, endT)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}

	// 分享次數
	shareCount, err := user_active_record.CountRecord(ctx, enums.ActiveTypeHeroShare, startT, endT)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}

	// 下載次數
	downloadGameCount, err := user_active_record.CountRecord(ctx, enums.ActiveTypeHeroDownload, startT, endT)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}

	node := map[string]string{
		"finished_user_count":   "總參與人數：完成遊戲的總人數，跳出不計算",
		"repeat_user_count":     "重複玩人數：同一玩家重複完整玩完的人數",
		"not_repeat_user_count": "不重複玩人數：玩家完整玩完的不重複玩人數",
		"unfinished_user_count": "跳出人數：未完成遊戲的跳出人數",
		"play_game_count":       "開始遊戲",
		"share_count":           "分享遊戲",
		"download_game_count":   "下載遊戲",
	}
	response := &UserCountResponse{
		FinishedUserCount:   repeatUserCount + notRepeatUserCount,
		RepeatUserCount:     repeatUserCount,
		NotRepeatUserCount:  notRepeatUserCount,
		UnfinishedUserCount: unfinishedUserCount,
		PlayGameCount:       playGameCount,
		ShareCount:          shareCount,
		DownloadGameCount:   downloadGameCount,
		StartAt:             startT,
		EndAt:               endT,
		Node:                node,
	}
	return controllers.ResponseSuccess(response, c)
}

func ScoreCount(c echo.Context) error {
	ctx := context.Background()
	score, err := strconv.ParseInt(c.QueryParam("score"), 10, 32)
	if err != nil {
		logger.Print("ScoreCount", "err", err.Error())
	}
	scoreCount := user_active_record.CountScore(ctx, int(score), c.QueryParam("start_at"), c.QueryParam("end_at"))
	return controllers.ResponseSuccess(scoreCount, c)
}

func Rank(c echo.Context) error {
	ctx := context.Background()
	userTotal, err := user.Count(ctx)
	socialUserID := c.QueryParam("fb_user_id")
	if socialUserID == "" {
		return controllers.ResponseFail(fmt.Errorf("lack fb_user_id"), c)
	}
	floatUserTotal := float32(userTotal)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	rankUsers, err := user.FindByRankBetterHeroScore(ctx, 10)
	subRankResp := make([]*subRankResponse, 0)
	for _, rankUser := range rankUsers {
		percentage := (floatUserTotal - (float32(countRankBetterME(rankUsers, rankUser.BetterHeroScore)))) / floatUserTotal * 100
		subRankResponse := &subRankResponse{
			LatestHeroScore: rankUser.LatestHeroScore,
			BetterHeroScore: rankUser.BetterHeroScore,
			SocialName:      rankUser.SocialName,
			SocialEmail:     rankUser.SocialEmail,
			SocialAvatarURL: rankUser.SocialAvatarURL,
			Percentage:      percentage,
		}
		subRankResp = append(subRankResp, subRankResponse)
	}
	if err != nil {
		return controllers.ResponseFail(err, c)
	}

	meUser, err := user.FindBySocialUserID(ctx, socialUserID)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	countBetterMe, err := user.CountBetterME(ctx, meUser.BetterHeroScore)
	if err != nil {
		return controllers.ResponseFail(err, c)
	}
	percentage := (floatUserTotal - float32(countBetterMe)) / floatUserTotal * 100
	me := &subRankResponse{
		LatestHeroScore: meUser.LatestHeroScore,
		BetterHeroScore: meUser.BetterHeroScore,
		SocialName:      meUser.SocialName,
		SocialEmail:     meUser.SocialEmail,
		SocialAvatarURL: meUser.SocialAvatarURL,
		Percentage:      percentage,
	}
	rankResponse := &RankResponse{
		Rank: subRankResp,
		ME:   me,
	}

	return controllers.ResponseSuccess(rankResponse, c)
}

func countRankBetterME(rankUsers []*ent.User, meScore int) int {
	betterME := 0
	for _, rankUser := range rankUsers {
		if rankUser.BetterHeroScore > meScore {
			betterME++
		}
	}
	return betterME
}
