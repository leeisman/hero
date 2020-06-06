package user

import (
	"context"
	"hero/database/ent"
	tableUser "hero/database/ent/user"
	"hero/pkg/db"
)

func FindUserBySocialUserID(ctx context.Context, socialUserID string) (*ent.User, error) {
	return db.Client().User.Query().
		Where(
			tableUser.And(
				tableUser.SocialUserIDEqualFold(socialUserID),
			),
		).
		First(ctx)
}

func CreateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	return db.Client().User.Create().
		SetID(user.ID).
		SetSocialUserID(*user.SocialUserID).
		SetSocialAvatarURL(*user.SocialAvatarURL).
		SetSocialEmail(*user.SocialEmail).
		SetSocialName(*user.SocialName).
		Save(ctx)
}
