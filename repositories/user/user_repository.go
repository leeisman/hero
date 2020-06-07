package user

import (
	"context"
	"hero/database/ent"
	tableUser "hero/database/ent/user"
	"hero/pkg/db/mysql"
	"time"
)

func FindBySocialUserID(ctx context.Context, socialUserID string) (*ent.User, error) {
	return mysql.Client().User.Query().
		Where(
			tableUser.And(
				tableUser.SocialUserIDEqualFold(socialUserID),
			),
		).
		First(ctx)
}

func CountRepeatUser(ctx context.Context, repeatStatus uint, startAt, endAt time.Time) (int, error) {
	return mysql.Client().User.Query().
		Where(
			tableUser.And(
				tableUser.HeroRepeatEQ(repeatStatus),
				tableUser.HeroPlayed(1),
				tableUser.CreatedAtGT(startAt),
				tableUser.CreatedAtLTE(endAt),
			),
		).Count(ctx)
}

func CountUnFinishedUser(ctx context.Context, startAt, endAt time.Time) (int, error) {
	return mysql.Client().User.Query().
		Where(
			tableUser.And(
				tableUser.HeroPlayed(0),
				tableUser.CreatedAtGTE(startAt),
				tableUser.CreatedAtLTE(endAt),
			),
		).Count(ctx)
}

func Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	return mysql.Client().User.Create().
		SetID(user.ID).
		SetSocialType(user.SocialType).
		SetSocialUserID(user.SocialUserID).
		SetSocialAvatarURL(user.SocialAvatarURL).
		SetSocialEmail(user.SocialEmail).
		SetSocialName(user.SocialName).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		Save(ctx)
}
