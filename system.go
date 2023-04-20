package c4

import (
	"context"
)

// SystemArgs describes the parameters available for configuring a system.
type SystemArgs struct {
	// The human-readable name of the system.
	Name string

	// A general description of the purpose of the system.
	Description string

	// Enables alternate styling reserved for external elements.
	External bool
}

// MustNewSystem is the same as NewSystem, but panics on any error.
func MustNewSystem(ctx context.Context, id string, args SystemArgs) *System {
	s, err := NewSystem(ctx, id, args)
	if err != nil {
		panic(err)
	}
	return s
}

// NewSystem constructs a system that can be used in a Diagram.
func NewSystem(ctx context.Context, id string, args SystemArgs) (*System, error) {
	s := &System{
		id:          id,
		name:        args.Name,
		description: args.Description,
		external:    args.External,
	}
	return s, nil
}

// System represents a C4 system (https://c4model.com/#SystemContextDiagram),
// which is the highest level of abstraction in the C4 model. A system generally
// describes a thing that "delivers value to its users" and can include both the
// system you are modelling as well as its dependent systems.
type System struct {
	id          string
	name        string
	description string
	external    bool
}

// Boundary returns a system boundary which can be used to group sub-containers
// at the container diagram level.
func (s *System) Boundary() Boundary {
	return &systemBoundary{System: s}
}

// ID satisfies the Element interface.
func (s *System) ID() string { return s.id }

type systemBoundary struct {
	*System
	elements []Element
}

func (sb *systemBoundary) AddElement(ctx context.Context, el Element) {
	sb.elements = append(sb.elements, el)
}
