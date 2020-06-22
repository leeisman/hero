package user_active_record

import (
	"context"
	"github.com/rs/xid"
	"hero/database/ent"
	tableUserActiveRecord "hero/database/ent/useractiverecord"
	"hero/pkg/db/mysql"
	"hero/pkg/logger"
	"time"
)

type SelectScoreCounts struct {
	ScoreCounts []struct {
		UserID string `json:"user_id"`
	}
}

func Create(ctx context.Context, userActiveRecord *ent.UserActiveRecord) (*ent.UserActiveRecord, error) {
	userActiveRecordXid := xid.New()
	userActiveRecordID := "UAR_" + userActiveRecordXid.String()
	return mysql.Client().UserActiveRecord.Create().
		SetID(userActiveRecordID).
		SetUserID(userActiveRecord.UserID).
		SetActiveType(userActiveRecord.ActiveType).
		SetNillableStartedAt(userActiveRecord.StartedAt).
		SetNillableEndedAt(userActiveRecord.EndedAt).
		SetNillableCreatedAt(userActiveRecord.CreatedAt).
		SetNillableUpdatedAt(userActiveRecord.UpdatedAt).
		Save(ctx)
}

func FindByID(ctx context.Context, ID string) (*ent.UserActiveRecord, error) {
	return mysql.Client().UserActiveRecord.Get(ctx, ID)
}

func CountRecord(ctx context.Context, activeType string, startAt, endAt time.Time) (int, error) {
	return mysql.Client().UserActiveRecord.Query().
		Where(
			tableUserActiveRecord.And(
				tableUserActiveRecord.ActiveType(activeType),
				tableUserActiveRecord.CreatedAtGTE(startAt),
				tableUserActiveRecord.CreatedAtLTE(endAt),
			),
		).Count(ctx)
}

func LeaveCount(ctx context.Context, activeType string, startAt, endAt time.Time) (int, error) {
	return mysql.Client().UserActiveRecord.Query().
		Where(
			tableUserActiveRecord.And(
				tableUserActiveRecord.IsFinished(0),
				tableUserActiveRecord.ActiveType(activeType),
				tableUserActiveRecord.CreatedAtGTE(startAt),
				tableUserActiveRecord.CreatedAtLTE(endAt),
			),
		).Count(ctx)
}

func CountScore(ctx context.Context, score int, startAt, endAt string) int {
	rows, err := mysql.DB().QueryContext(ctx, "select count(distinct(user_id)) from user_active_records where score = ? and started_at >= ? and ended_at <= ? ", score, startAt, endAt)
	if err != nil {
		logger.Print("err", err.Error())
	}
	defer rows.Close()

	userIDCount := 0
	for rows.Next() {
		var count int
		if err := rows.Scan(&count); err != nil {
			logger.Print("err", err.Error())
		}
		userIDCount = count
	}

	return userIDCount
}
