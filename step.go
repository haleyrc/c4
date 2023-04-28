package c4

import (
	"context"
)

// StepArgs describes the parameters available for configuring a step
// between two elements in a diagram.
type StepArgs struct {
	// The subject of the step e.g. "Src <<uses>> Dst".
	Src Element

	// The object of the step e.g. "Src <<uses>> Dst".
	Dst Element

	// The verb of the step. This should be in the indicative tense from the
	// perspective of the source element e.g. "Src uses Dst" as opposed to "Src
	// is used by Dst" or "Src used/use/etc. Dst".
	Description string

	// An optional list of technologies describing the nature of the interaction
	// between the elements of the step e.g. "JSON/HTTPS" or "SQL/TCP".
	Technologies []string

	Direction Direction
}

func newStep(ctx context.Context, args StepArgs) (*step, error) {
	s := &step{
		src:          args.Src,
		dst:          args.Dst,
		description:  args.Description,
		technologies: args.Technologies,
		direction:    args.Direction,
	}

	return s, nil
}

type step struct {
	src          Element
	dst          Element
	description  string
	technologies []string
	direction    Direction
}
