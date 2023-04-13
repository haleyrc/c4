package c4

import (
	"context"
)

// ContainerArgs describes the parameters available for configuring a container.
type ContainerArgs struct {
	// The human-readable name of the container.
	Name string

	// A general description of the purpose of the container.
	Description string

	// An optional list of technologies describing the container e.g. Javascript.
	Technologies []string
}

// MustNewContainer is the same as NewContainer, but panics on any error.
func MustNewContainer(ctx context.Context, id string, args ContainerArgs) *Container {
	c, err := NewContainer(ctx, id, args)
	if err != nil {
		panic(err)
	}
	return c
}

// NewContainer constructs a container that can be used in a Diagram.
func NewContainer(ctx context.Context, id string, args ContainerArgs) (*Container, error) {
	c := &Container{
		id:          id,
		name:        args.Name,
		description: args.Description,
	}
	return c, nil
}

// Container represents a C4 container (https://c4model.com/#ContainerDiagram),
// which the C4 documentation describes as a "separately runnable/deployable
// unit" e.g. a single-page application, mobile app, etc.
//
// Note that while the C4 documentation mentions databases as containers, this
// package provides a separate data type (Database) for database containers.
type Container struct {
	id           string
	name         string
	description  string
	technologies []string
}

// Boundary returns a container boundary which can be used to group
// sub-components at the component diagram level.
func (c *Container) Boundary() Boundary {
	return &containerBoundary{Container: c}
}

// ID satisfies the Element interface.
func (c *Container) ID() string { return c.id }

type containerBoundary struct {
	*Container
	elements []Element
}

func (cb *containerBoundary) AddElement(ctx context.Context, el Element) {
	cb.elements = append(cb.elements, el)
}
