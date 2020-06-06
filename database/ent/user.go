// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hero/database/ent/user"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// SocialUserID holds the value of the "social_user_id" field.
	SocialUserID *string `json:"social_user_id,omitempty"`
	// SocialAvatarURL holds the value of the "social_avatar_url" field.
	SocialAvatarURL *string `json:"social_avatar_url,omitempty"`
	// SocialEmail holds the value of the "social_email" field.
	SocialEmail *string `json:"social_email,omitempty"`
	// SocialName holds the value of the "social_name" field.
	SocialName *string `json:"social_name,omitempty"`
	// SocialType holds the value of the "social_type" field.
	SocialType *string `json:"social_type,omitempty"`
	// SocialPayload holds the value of the "social_payload" field.
	SocialPayload *string `json:"social_payload,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // id
		&sql.NullString{}, // social_user_id
		&sql.NullString{}, // social_avatar_url
		&sql.NullString{}, // social_email
		&sql.NullString{}, // social_name
		&sql.NullString{}, // social_type
		&sql.NullString{}, // social_payload
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		u.ID = value.String
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field social_user_id", values[0])
	} else if value.Valid {
		u.SocialUserID = new(string)
		*u.SocialUserID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field social_avatar_url", values[1])
	} else if value.Valid {
		u.SocialAvatarURL = new(string)
		*u.SocialAvatarURL = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field social_email", values[2])
	} else if value.Valid {
		u.SocialEmail = new(string)
		*u.SocialEmail = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field social_name", values[3])
	} else if value.Valid {
		u.SocialName = new(string)
		*u.SocialName = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field social_type", values[4])
	} else if value.Valid {
		u.SocialType = new(string)
		*u.SocialType = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field social_payload", values[5])
	} else if value.Valid {
		u.SocialPayload = new(string)
		*u.SocialPayload = value.String
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[6])
	} else if value.Valid {
		u.CreatedAt = new(time.Time)
		*u.CreatedAt = value.Time
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[7])
	} else if value.Valid {
		u.UpdatedAt = new(time.Time)
		*u.UpdatedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	if v := u.SocialUserID; v != nil {
		builder.WriteString(", social_user_id=")
		builder.WriteString(*v)
	}
	if v := u.SocialAvatarURL; v != nil {
		builder.WriteString(", social_avatar_url=")
		builder.WriteString(*v)
	}
	if v := u.SocialEmail; v != nil {
		builder.WriteString(", social_email=")
		builder.WriteString(*v)
	}
	if v := u.SocialName; v != nil {
		builder.WriteString(", social_name=")
		builder.WriteString(*v)
	}
	if v := u.SocialType; v != nil {
		builder.WriteString(", social_type=")
		builder.WriteString(*v)
	}
	if v := u.SocialPayload; v != nil {
		builder.WriteString(", social_payload=")
		builder.WriteString(*v)
	}
	if v := u.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := u.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
