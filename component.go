package c4

import (
	"context"
)

// ComponentArgs describes the parameters available for configuring a component.
type ComponentArgs struct {
	// The human-readable name of the component.
	Name string

	// A general description of the purpose of the component.
	Description string

	// An optional list of technologies describing the component e.g. Spring Bean.
	Technologies []string
}

// MustNewComponent is the same as NewComponent, but panics on any error.
func MustNewComponent(ctx context.Context, id string, args ComponentArgs) *Component {
	c, err := NewComponent(ctx, id, args)
	if err != nil {
		panic(err)
	}
	return c
}

// NewComponent constructs a component that can be used in a Diagram.
func NewComponent(ctx context.Context, id string, args ComponentArgs) (*Component, error) {
	c := &Component{
		id:           id,
		name:         args.Name,
		description:  args.Description,
		technologies: args.Technologies,
	}
	return c, nil
}

// Component represents a C4 component (https://c4model.com/#ComponentDiagram),
// which can be described as the consituent pieces of a container. Some examples
// of potential components are: controllers, models, packages, etc.
type Component struct {
	id           string
	name         string
	description  string
	technologies []string
}

// ID satisfies the Element interface.
func (c *Component) ID() string { return c.id }
