package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// UserActiveRecord holds the schema definition for the UserActiveRecord entity.
type UserActiveRecord struct {
	ent.Schema
}

// Fields of the UserActiveRecord.
func (UserActiveRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("user_id"),
		field.String("active_type"),
		field.Int("score").Default(0),
		field.Uint("is_finished").Comment("0:未完成, 1:完成").Default(0),
		field.Time("started_at").Nillable().Optional(),
		field.Time("ended_at").Nillable().Optional(),
		field.Time("created_at").Nillable().Optional(),
		field.Time("updated_at").Nillable().Optional(),
	}
}

// Edges of the UserActiveRecord.
func (UserActiveRecord) Edges() []ent.Edge {
	return nil
}
