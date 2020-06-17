// Code generated by entc, DO NOT EDIT.

package ent

import (
	"hero/database/ent/schema"
	"hero/database/ent/user"
	"hero/database/ent/useractiverecord"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescLatestHeroScore is the schema descriptor for latest_hero_score field.
	userDescLatestHeroScore := userFields[1].Descriptor()
	// user.DefaultLatestHeroScore holds the default value on creation for the latest_hero_score field.
	user.DefaultLatestHeroScore = userDescLatestHeroScore.Default.(int)
	// userDescBetterHeroScore is the schema descriptor for better_hero_score field.
	userDescBetterHeroScore := userFields[2].Descriptor()
	// user.DefaultBetterHeroScore holds the default value on creation for the better_hero_score field.
	user.DefaultBetterHeroScore = userDescBetterHeroScore.Default.(int)
	// userDescSocialUserID is the schema descriptor for social_user_id field.
	userDescSocialUserID := userFields[4].Descriptor()
	// user.DefaultSocialUserID holds the default value on creation for the social_user_id field.
	user.DefaultSocialUserID = userDescSocialUserID.Default.(string)
	// userDescSocialAvatarURL is the schema descriptor for social_avatar_url field.
	userDescSocialAvatarURL := userFields[5].Descriptor()
	// user.DefaultSocialAvatarURL holds the default value on creation for the social_avatar_url field.
	user.DefaultSocialAvatarURL = userDescSocialAvatarURL.Default.(string)
	// userDescSocialEmail is the schema descriptor for social_email field.
	userDescSocialEmail := userFields[6].Descriptor()
	// user.DefaultSocialEmail holds the default value on creation for the social_email field.
	user.DefaultSocialEmail = userDescSocialEmail.Default.(string)
	// userDescSocialName is the schema descriptor for social_name field.
	userDescSocialName := userFields[7].Descriptor()
	// user.DefaultSocialName holds the default value on creation for the social_name field.
	user.DefaultSocialName = userDescSocialName.Default.(string)
	// userDescSocialType is the schema descriptor for social_type field.
	userDescSocialType := userFields[8].Descriptor()
	// user.DefaultSocialType holds the default value on creation for the social_type field.
	user.DefaultSocialType = userDescSocialType.Default.(string)
	// userDescSocialPayload is the schema descriptor for social_payload field.
	userDescSocialPayload := userFields[9].Descriptor()
	// user.DefaultSocialPayload holds the default value on creation for the social_payload field.
	user.DefaultSocialPayload = userDescSocialPayload.Default.(string)
	// userDescHeroPlayed is the schema descriptor for hero_played field.
	userDescHeroPlayed := userFields[10].Descriptor()
	// user.DefaultHeroPlayed holds the default value on creation for the hero_played field.
	user.DefaultHeroPlayed = userDescHeroPlayed.Default.(uint)
	// userDescHeroRepeat is the schema descriptor for hero_repeat field.
	userDescHeroRepeat := userFields[11].Descriptor()
	// user.DefaultHeroRepeat holds the default value on creation for the hero_repeat field.
	user.DefaultHeroRepeat = userDescHeroRepeat.Default.(uint)
	useractiverecordFields := schema.UserActiveRecord{}.Fields()
	_ = useractiverecordFields
	// useractiverecordDescScore is the schema descriptor for score field.
	useractiverecordDescScore := useractiverecordFields[3].Descriptor()
	// useractiverecord.DefaultScore holds the default value on creation for the score field.
	useractiverecord.DefaultScore = useractiverecordDescScore.Default.(int)
	// useractiverecordDescIsFinished is the schema descriptor for is_finished field.
	useractiverecordDescIsFinished := useractiverecordFields[4].Descriptor()
	// useractiverecord.DefaultIsFinished holds the default value on creation for the is_finished field.
	useractiverecord.DefaultIsFinished = useractiverecordDescIsFinished.Default.(uint)
}
