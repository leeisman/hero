package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("pk"),
		field.Int("latest_hero_score").Default(0),
		field.Int("better_hero_score").Default(0),
		field.Time("better_hero_score_at").Optional(),
		field.String("social_user_id").Default(""),
		field.String("social_avatar_url").Default(""),
		field.String("social_email").Default(""),
		field.String("social_name").Default(""),
		field.String("social_type").Default(""),
		field.String("social_payload").Default(""),
		field.Uint("hero_played").Default(0).Comment("完成過hero"),
		field.Uint("hero_repeat").Default(0).Comment("重過玩hero"),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
