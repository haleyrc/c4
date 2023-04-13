package c4

import (
	"context"
)

// DatabaseArgs describes the parameters available for configuring a database
// container.
type DatabaseArgs struct {
	// The human-readable name of the database.
	Name string

	// A general description of the purpose of the database.
	Description string

	// An optional list of technologies describing the database e.g. PostgreSQL.
	Technologies []string
}

// MustNewDatabase is the same as NewDatabase, but panics on any error.
func MustNewDatabase(ctx context.Context, id string, args DatabaseArgs) *Database {
	db, err := NewDatabase(ctx, id, args)
	if err != nil {
		panic(err)
	}
	return db
}

// NewDatabase constructs a database container that can be used in a Diagram.
func NewDatabase(ctx context.Context, id string, args DatabaseArgs) (*Database, error) {
	c := &Database{
		id:          id,
		name:        args.Name,
		description: args.Description,
	}
	return c, nil
}

// Database represents a C4 container (https://c4model.com/#ContainerDiagram)
// specifically for describing databases. The C4 documentation refers to
// database alongside other containers, but conceptually they stand apart and we
// likewise represent them differently in the resultant diagram.
type Database struct {
	id           string
	name         string
	description  string
	technologies []string
}

// ID satisfies the Element interface.
func (db *Database) ID() string { return db.id }
