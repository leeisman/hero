// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "hero_score", Type: field.TypeInt},
		{Name: "social_user_id", Type: field.TypeString, Default: ""},
		{Name: "social_avatar_url", Type: field.TypeString, Default: ""},
		{Name: "social_email", Type: field.TypeString, Default: ""},
		{Name: "social_name", Type: field.TypeString, Default: ""},
		{Name: "social_type", Type: field.TypeString, Default: ""},
		{Name: "social_payload", Type: field.TypeString, Default: ""},
		{Name: "hero_played", Type: field.TypeUint},
		{Name: "hero_repeat", Type: field.TypeUint},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UserActiveRecordsColumns holds the columns for the "user_active_records" table.
	UserActiveRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "active_type", Type: field.TypeString},
		{Name: "score", Type: field.TypeInt},
		{Name: "is_finished", Type: field.TypeUint},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "ended_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// UserActiveRecordsTable holds the schema information for the "user_active_records" table.
	UserActiveRecordsTable = &schema.Table{
		Name:        "user_active_records",
		Columns:     UserActiveRecordsColumns,
		PrimaryKey:  []*schema.Column{UserActiveRecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
		UserActiveRecordsTable,
	}
)

func init() {
}
