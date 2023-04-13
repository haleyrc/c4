package c4

import (
	"context"
)

// A Boundary represents a special type of element that can also accept child
// elements in order to group them visually in the resultant diagram. You can
// add a Boundary to a Diagram using Diagram.AddElement in the same way you can
// add other Element types.
type Boundary interface {
	Element
	AddElement(ctx context.Context, el Element)
}

// An Element is a type that can be added to a diagram to have it displayed
// visually.
type Element interface {
	ID() string
}
