package c4

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

// Layout represents the overall layout flow of the resultant diagram.
type Layout string

const DefaultLayout Layout = LayoutTopDown

const (
	LayoutTopDown   Layout = "LAYOUT_TOP_DOWN"
	LayoutLeftRight Layout = "LAYOUT_LEFT_RIGHT"
)

// NewDiagram constructs a diagram that can be converted to a C4-enabled
// PlantUML representation.
func NewDiagram(ctx context.Context, title string, opts ...DiagramOption) (*Diagram, error) {
	d := &Diagram{
		title:  title,
		layout: DefaultLayout,
	}

	for _, opt := range opts {
		opt(d)
	}

	return d, nil
}

// Diagram represents the top-level container for systems, containers, etc. and
// their relations. Once the diagram is populated, it can be converted to a
// C4-enabled PlantUML representation that can be used to generate visual
// diagrams using the PlantUML CLI.
type Diagram struct {
	title     string
	layout    Layout
	elements  []Element
	relations []*relation
}

// AddElement adds an element to the resultant PlantUML specification.
func (d *Diagram) AddElement(ctx context.Context, el Element) {
	d.elements = append(d.elements, el)
}

// NewRelation adds a relation between two elements to the PlantUML
// specification. Note that if the constituent elements aren't added to the
// diagram, they will not appear and this relation will have no effect on the
// resultant visual diagram.
func (d *Diagram) NewRelation(ctx context.Context, args RelationArgs, opts ...RelationOption) error {
	rel, err := newRelation(ctx, args, opts...)
	if err != nil {
		return err
	}

	d.relations = append(d.relations, rel)

	return nil
}

// PlantUML renders the Diagram as a C4-enabled PlantUML specification to the
// provided writer.
func (d *Diagram) PlantUML(ctx context.Context, w io.Writer) error {
	var buff bytes.Buffer

	if err := writePreamble(ctx, &buff, d.title, d.layout); err != nil {
		return err
	}

	for _, el := range d.elements {
		if err := plantUML(ctx, &buff, el); err != nil {
			return err
		}
	}

	for _, rel := range d.relations {
		if err := plantUML(ctx, &buff, rel); err != nil {
			return err
		}
	}

	if err := writeEpilogue(ctx, &buff); err != nil {
		return err
	}

	if _, err := io.Copy(w, &buff); err != nil {
		return err
	}

	return nil
}

// DiagramOptions are used to modify the display characteristics of a diagram.
type DiagramOption func(*Diagram)

// WithLayout allows you to set an explicit layout direction for the diagram.
func WithLayout(l Layout) DiagramOption {
	return func(d *Diagram) {
		d.layout = l
	}
}

func writePreamble(ctx context.Context, buff *bytes.Buffer, title string, layout Layout) error {
	fmt.Fprintln(buff, "@startuml", title)
	fmt.Fprintln(buff, "!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml")
	fmt.Fprintln(buff, "!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml")
	fmt.Fprintln(buff)
	fmt.Fprintf(buff, "%s()\n", layout)
	fmt.Fprintln(buff)
	return nil
}

func writeEpilogue(ctx context.Context, buff *bytes.Buffer) error {
	fmt.Fprintln(buff, "@enduml")
	return nil
}
