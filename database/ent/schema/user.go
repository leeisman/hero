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
		field.String("social_user_id").Nillable().Default("").Optional(),
		field.String("social_avatar_url").Nillable().Default("").Optional(),
		field.String("social_email").Nillable().Default("").Optional(),
		field.String("social_name").Nillable().Default("").Optional(),
		field.String("social_type").Nillable().Default("").Optional(),
		field.String("social_payload").Nillable().Default("").Optional(),
		field.Time("created_at").Nillable().Optional(),
		field.Time("updated_at").Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
