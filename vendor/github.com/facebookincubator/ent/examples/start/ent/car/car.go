// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package car

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID           = "id"    // FieldModel holds the string denoting the model vertex property in the database.
	FieldModel        = "model" // FieldRegisteredAt holds the string denoting the registered_at vertex property in the database.
	FieldRegisteredAt = "registered_at"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the car in the database.
	Table = "cars"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "cars"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_cars"
)

// Columns holds all SQL columns for car fields.
var Columns = []string{
	FieldID,
	FieldModel,
	FieldRegisteredAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Car type.
var ForeignKeys = []string{
	"user_cars",
}
