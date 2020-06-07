package user_active_record

import (
	"context"
	"github.com/rs/xid"
	"hero/database/ent"
	"hero/pkg/db/mysql"
)

func Create(ctx context.Context, userActiveRecord *ent.UserActiveRecord) (*ent.UserActiveRecord, error) {
	userActiveRecordXid := xid.New()
	userActiveRecordID := "UAR_" + userActiveRecordXid.String()
	return mysql.Client().UserActiveRecord.Create().
		SetID(userActiveRecordID).
		SetNillableUserID(userActiveRecord.UserID).
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
