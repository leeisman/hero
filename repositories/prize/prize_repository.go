package prize

import (
	"context"
	"hero/database/ent"
	tablePrize "hero/database/ent/prize"
	"hero/pkg/db/mysql"
)

func FindBySocialUserIDAndDate(ctx context.Context, SocialUserID string, date string) (*ent.Prize, error) {
	return mysql.Client().Prize.Query().
		Where(
			tablePrize.And(
				tablePrize.SocialUserIDEQ(SocialUserID),
				tablePrize.DateEQ(date),
			),
		).
		First(ctx)
}

func Create(ctx context.Context, SocialUserID string, date string) (*ent.Prize, error) {
	return mysql.Client().Prize.Create().SetSocialUserID(SocialUserID).SetDate(date).Save(ctx)
}
