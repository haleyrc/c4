package c4

import "context"

// EnterpriseBoundaryArgs describes the parameters available for configuring an
// EnterpriseBoundary.
type EnterpriseBoundaryArgs struct {
	// The human-readable name of the enterprise.
	Name string
}

// MustNewEnterpriseBoundary is the same as NewEnterpriseBoundary, but panics on
// any error.
func MustNewEnterpriseBoundary(ctx context.Context, id string, args EnterpriseBoundaryArgs) *EnterpriseBoundary {
	b, err := NewEnterpriseBoundary(ctx, id, args)
	if err != nil {
		panic(err)
	}
	return b
}

// NewEnterpriseBoundary constructs an enterprise boundary which can be used to
// group elements belonging to a common parent enterprise.
func NewEnterpriseBoundary(ctx context.Context, id string, args EnterpriseBoundaryArgs) (*EnterpriseBoundary, error) {
	b := &EnterpriseBoundary{
		id:   id,
		name: args.Name,
	}
	return b, nil
}

// EnterpriseBoundary represents a logical grouping of elements that belong to
// a common enterprise e.g. systems owned and operated by the same company.
type EnterpriseBoundary struct {
	id       string
	name     string
	elements []Element
}

// AddElement adds child elements to the parent EnterpriseBoundary.
func (eb *EnterpriseBoundary) AddElement(ctx context.Context, el Element) {
	eb.elements = append(eb.elements, el)
}

// ID satisfies the Element interface.
func (eb *EnterpriseBoundary) ID() string { return eb.id }
