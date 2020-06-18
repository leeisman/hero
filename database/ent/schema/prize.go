package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Prize struct {
	ent.Schema
}

// Fields of the User.
func (Prize) Fields() []ent.Field {
	return []ent.Field{
		field.String("social_user_id").Default(""),
		field.String("date").Optional(),
	}
}

// Edges of the User.
func (Prize) Edges() []ent.Edge {
	return nil
}
