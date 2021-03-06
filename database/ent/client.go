// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"hero/database/ent/migrate"

	"hero/database/ent/prize"
	"hero/database/ent/user"
	"hero/database/ent/useractiverecord"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Prize is the client for interacting with the Prize builders.
	Prize *PrizeClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserActiveRecord is the client for interacting with the UserActiveRecord builders.
	UserActiveRecord *UserActiveRecordClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Prize = NewPrizeClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserActiveRecord = NewUserActiveRecordClient(c.config)
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:           cfg,
		Prize:            NewPrizeClient(cfg),
		User:             NewUserClient(cfg),
		UserActiveRecord: NewUserActiveRecordClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:           cfg,
		Prize:            NewPrizeClient(cfg),
		User:             NewUserClient(cfg),
		UserActiveRecord: NewUserActiveRecordClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Prize.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Prize.Use(hooks...)
	c.User.Use(hooks...)
	c.UserActiveRecord.Use(hooks...)
}

// PrizeClient is a client for the Prize schema.
type PrizeClient struct {
	config
}

// NewPrizeClient returns a client for the Prize from the given config.
func NewPrizeClient(c config) *PrizeClient {
	return &PrizeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `prize.Hooks(f(g(h())))`.
func (c *PrizeClient) Use(hooks ...Hook) {
	c.hooks.Prize = append(c.hooks.Prize, hooks...)
}

// Create returns a create builder for Prize.
func (c *PrizeClient) Create() *PrizeCreate {
	mutation := newPrizeMutation(c.config, OpCreate)
	return &PrizeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Prize.
func (c *PrizeClient) Update() *PrizeUpdate {
	mutation := newPrizeMutation(c.config, OpUpdate)
	return &PrizeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PrizeClient) UpdateOne(pr *Prize) *PrizeUpdateOne {
	mutation := newPrizeMutation(c.config, OpUpdateOne, withPrize(pr))
	return &PrizeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PrizeClient) UpdateOneID(id int) *PrizeUpdateOne {
	mutation := newPrizeMutation(c.config, OpUpdateOne, withPrizeID(id))
	return &PrizeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Prize.
func (c *PrizeClient) Delete() *PrizeDelete {
	mutation := newPrizeMutation(c.config, OpDelete)
	return &PrizeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PrizeClient) DeleteOne(pr *Prize) *PrizeDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PrizeClient) DeleteOneID(id int) *PrizeDeleteOne {
	builder := c.Delete().Where(prize.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PrizeDeleteOne{builder}
}

// Create returns a query builder for Prize.
func (c *PrizeClient) Query() *PrizeQuery {
	return &PrizeQuery{config: c.config}
}

// Get returns a Prize entity by its id.
func (c *PrizeClient) Get(ctx context.Context, id int) (*Prize, error) {
	return c.Query().Where(prize.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PrizeClient) GetX(ctx context.Context, id int) *Prize {
	pr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pr
}

// Hooks returns the client hooks.
func (c *PrizeClient) Hooks() []Hook {
	return c.hooks.Prize
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserActiveRecordClient is a client for the UserActiveRecord schema.
type UserActiveRecordClient struct {
	config
}

// NewUserActiveRecordClient returns a client for the UserActiveRecord from the given config.
func NewUserActiveRecordClient(c config) *UserActiveRecordClient {
	return &UserActiveRecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `useractiverecord.Hooks(f(g(h())))`.
func (c *UserActiveRecordClient) Use(hooks ...Hook) {
	c.hooks.UserActiveRecord = append(c.hooks.UserActiveRecord, hooks...)
}

// Create returns a create builder for UserActiveRecord.
func (c *UserActiveRecordClient) Create() *UserActiveRecordCreate {
	mutation := newUserActiveRecordMutation(c.config, OpCreate)
	return &UserActiveRecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for UserActiveRecord.
func (c *UserActiveRecordClient) Update() *UserActiveRecordUpdate {
	mutation := newUserActiveRecordMutation(c.config, OpUpdate)
	return &UserActiveRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserActiveRecordClient) UpdateOne(uar *UserActiveRecord) *UserActiveRecordUpdateOne {
	mutation := newUserActiveRecordMutation(c.config, OpUpdateOne, withUserActiveRecord(uar))
	return &UserActiveRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserActiveRecordClient) UpdateOneID(id string) *UserActiveRecordUpdateOne {
	mutation := newUserActiveRecordMutation(c.config, OpUpdateOne, withUserActiveRecordID(id))
	return &UserActiveRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserActiveRecord.
func (c *UserActiveRecordClient) Delete() *UserActiveRecordDelete {
	mutation := newUserActiveRecordMutation(c.config, OpDelete)
	return &UserActiveRecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserActiveRecordClient) DeleteOne(uar *UserActiveRecord) *UserActiveRecordDeleteOne {
	return c.DeleteOneID(uar.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserActiveRecordClient) DeleteOneID(id string) *UserActiveRecordDeleteOne {
	builder := c.Delete().Where(useractiverecord.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserActiveRecordDeleteOne{builder}
}

// Create returns a query builder for UserActiveRecord.
func (c *UserActiveRecordClient) Query() *UserActiveRecordQuery {
	return &UserActiveRecordQuery{config: c.config}
}

// Get returns a UserActiveRecord entity by its id.
func (c *UserActiveRecordClient) Get(ctx context.Context, id string) (*UserActiveRecord, error) {
	return c.Query().Where(useractiverecord.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserActiveRecordClient) GetX(ctx context.Context, id string) *UserActiveRecord {
	uar, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return uar
}

// Hooks returns the client hooks.
func (c *UserActiveRecordClient) Hooks() []Hook {
	return c.hooks.UserActiveRecord
}
