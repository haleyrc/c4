package c4

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

// Layout represents the overall layout flow of the resultant diagram.
type Layout string

const (
	LayoutTopDown  Layout = "LAYOUT_TOP_DOWN"
	LayoutPortrait Layout = LayoutTopDown

	LayoutLandscape Layout = "LAYOUT_LANDSCAPE"

	// Deprecated: The PlantUML algorithm for LAYOUT_LEFT_RIGHT is insane and
	// also rotates the directions nonsensically (a reflection across the line
	// y = -x; a reflection across the y-axis and a 90 degree counter-clockwise
	// rotation). This causes the relations in the resultant diagram to be
	// different from the stated directions in the code.
	//
	// In the next major version of this library, this option will switch to
	// LAYOUT_LANDSCAPE under the hood, making it more in line with user
	// expectations.
	LayoutLeftRight Layout = "LAYOUT_LEFT_RIGHT"

	DefaultLayout Layout = LayoutTopDown
)

// NewDiagram constructs a diagram that can be converted to a C4-enabled
// PlantUML representation.
func NewDiagram(ctx context.Context, title string, opts ...DiagramOption) (*Diagram, error) {
	d := &Diagram{
		title:  title,
		layout: DefaultLayout,
		theme:  DefaultTheme(),
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
	title            string
	layout           Layout
	theme            Theme
	elements         []Element
	relations        []*relation
	sketch           bool
	legend           bool
	hideElementTypes bool
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

	if err := d.writePreamble(ctx, &buff, d.title, d.layout); err != nil {
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

	if err := d.writeEpilogue(ctx, &buff); err != nil {
		return err
	}

	if _, err := io.Copy(w, &buff); err != nil {
		return err
	}

	return nil
}

func (d *Diagram) writePreamble(ctx context.Context, buff *bytes.Buffer, title string, layout Layout) error {
	fmt.Fprintln(buff, "@startuml", title)
	fmt.Fprintln(buff, "!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml")
	fmt.Fprintln(buff, "!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml")
	fmt.Fprintln(buff, "!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Deployment.puml")
	fmt.Fprintln(buff)
	fmt.Fprintln(buff, "WithoutPropertyHeader()")
	fmt.Fprintln(buff)
	fmt.Fprintf(buff, "%s()\n", layout)
	if d.sketch {
		fmt.Fprintln(buff, `LAYOUT_AS_SKETCH()`)
	}
	fmt.Fprintln(buff)
	fmt.Fprintf(buff, `UpdateElementStyle(system, $bgColor="%s", $fontColor="%s")`, d.theme.System.BackgroundColor, d.theme.System.FontColor)
	fmt.Fprintln(buff)
	fmt.Fprintf(buff, `UpdateElementStyle(container, $bgColor="%s", $fontColor="%s")`, d.theme.Container.BackgroundColor, d.theme.Container.FontColor)
	fmt.Fprintln(buff)
	fmt.Fprintf(buff, `UpdateElementStyle(component, $bgColor="%s", $fontColor="%s")`, d.theme.Component.BackgroundColor, d.theme.Component.FontColor)
	fmt.Fprintln(buff)
	fmt.Fprintf(buff, `UpdateElementStyle(person, $bgColor="%s", $fontColor="%s")`, d.theme.Person.BackgroundColor, d.theme.Person.FontColor)
	fmt.Fprintln(buff)
	return nil
}

func (d *Diagram) writeEpilogue(ctx context.Context, buff *bytes.Buffer) error {
	if d.hideElementTypes {
		fmt.Fprintln(buff, `HIDE_STEREOTYPE()`)
		fmt.Fprintln(buff)
	}
	if d.legend {
		// The hideStereotype paramater is hard-coded to false here in order to
		// make the default behavior more consistent with expectations. In
		// concert with the individual option to hide stereotypes, you can still
		// easily achieve the same result in an opt-in way.
		fmt.Fprintln(buff, `SHOW_LEGEND($hideStereotype=false)`)
		fmt.Fprintln(buff)
	}
	fmt.Fprintln(buff, "@enduml")
	return nil
}

// DiagramOptions are used to modify the display characteristics of a diagram.
type DiagramOption func(*Diagram)

// AsSketch causes the diagram to be rendered in a sketch-like style. The
// diagram also gets a disclaimer indicating that it still needs to be validated
// and is for discussion only.
func AsSketch() DiagramOption {
	return func(d *Diagram) {
		d.sketch = true
	}
}

// HideElementTypes hides the type line that appears at the top of rendered
// diagram elements. This can be especially useful in concert with the
// WithLegend option or if the type of elements is unimportant to your diagram.
func HideElementTypes() DiagramOption {
	return func(d *Diagram) {
		d.hideElementTypes = true
	}
}

// WithLayout allows you to set an explicit layout direction for the diagram.
func WithLayout(l Layout) DiagramOption {
	return func(d *Diagram) {
		d.layout = l
	}
}

// WithLegend enables a legend mapping colors to element types.
func WithLegend() DiagramOption {
	return func(d *Diagram) {
		d.legend = true
	}
}

// WithTheme allows you to set a custom theme for the diagram. You can either
// create a theme from scratch or use c4.DefaultTheme() and modify the values
// you care about.
func WithTheme(t Theme) DiagramOption {
	return func(d *Diagram) {
		d.theme = t
	}
}
