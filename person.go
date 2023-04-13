package c4

import (
	"context"
)

// PersonArgs describes the parameters available for configuring a person.
type PersonArgs struct {
	// The human-readable name of the person.
	Name string

	// A general description of the purpose of the person.
	Description string
}

// MustNewPerson is the same as NewPerson, but panics on any error.
func MustNewPerson(ctx context.Context, id string, args PersonArgs) *Person {
	p, err := NewPerson(ctx, id, args)
	if err != nil {
		panic(err)
	}
	return p
}

// NewPerson constructs a person that can be used in a Diagram.
func NewPerson(ctx context.Context, id string, args PersonArgs) (*Person, error) {
	p := &Person{
		id:          id,
		name:        args.Name,
		description: args.Description,
	}
	return p, nil
}

// Component represents a C4 person (https://c4model.com/#Abstractions) which is
// a user/persona/etc. that interacts with your system.
type Person struct {
	id          string
	name        string
	description string
}

// ID satisfies the Element interface.
func (p *Person) ID() string { return p.id }
