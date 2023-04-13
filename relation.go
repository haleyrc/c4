package c4

import (
	"context"
)

// Direction represents the "arrow key" direction of the relation in the
// resultant diagram. PlantUML will use the direction when laying out elements
// within the physical constraints of the diagram itself.
type Direction string

const (
	DirectionUp    Direction = "Up"
	DirectionDown  Direction = "Down"
	DirectionLeft  Direction = "Left"
	DirectionRight Direction = "Right"
)

// RelationArgs describes the parameters available for configuring a relation
// between two elements in a diagram.
type RelationArgs struct {
	// The subject of the relation e.g. "Src <<uses>> Dst".
	Src Element

	// The object of the relation e.g. "Src <<uses>> Dst".
	Dst Element

	// The verb of the relation. This should be in the indicative tense from the
	// perspective of the source element e.g. "Src uses Dst" as opposed to "Src is
	// used by Dst" or "Src used/use/etc. Dst".
	Description string

	// An optional list of technologies describing the nature of the interaction
	// between the elements of the relation e.g. "JSON/HTTPS" or "SQL/TCP".
	Technologies []string
}

// RelationOptions are used to modify display characteristics of a relation.
type RelationOption func(*relation)

// WithDirection allows you to set an explicit direction for the relation.
func WithDirection(d Direction) RelationOption {
	return func(r *relation) {
		r.direction = d
	}
}

func newRelation(ctx context.Context, args RelationArgs, opts ...RelationOption) (*relation, error) {
	rel := &relation{
		src:          args.Src,
		dst:          args.Dst,
		description:  args.Description,
		technologies: args.Technologies,
	}

	for _, opt := range opts {
		opt(rel)
	}

	return rel, nil
}

type relation struct {
	src          Element
	dst          Element
	description  string
	technologies []string
	direction    Direction
}
