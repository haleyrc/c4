package c4

import (
	"context"
	"fmt"
	"io"
	"strings"
)

func plantUML(ctx context.Context, w io.Writer, el interface{}) error {
	switch v := el.(type) {
	case *Component:
		technologies := strings.Join(v.technologies, ", ")
		fmt.Fprintf(w, `Component(%s, "%s", "%s", "%s")`, v.ID(), v.name, technologies, v.description)
		fmt.Fprintln(w)
	case *Container:
		technologies := strings.Join(v.technologies, ", ")
		fmt.Fprintf(w, `Container(%s, "%s", "%s", "%s")`, v.ID(), v.name, technologies, v.description)
		fmt.Fprintln(w)
	case *containerBoundary:
		fmt.Fprintf(w, `Container_Boundary(%s, "%s") {`, v.ID(), v.name)
		fmt.Fprintln(w)
		for _, el := range v.elements {
			fmt.Fprintf(w, "\t")
			if err := plantUML(ctx, w, el); err != nil {
				return err
			}
		}
		fmt.Fprintln(w, "}")
	case *Database:
		technologies := strings.Join(v.technologies, ", ")
		fmt.Fprintf(w, `ContainerDb(%s, "%s", "%s", "%s")`, v.ID(), v.name, technologies, v.description)
		fmt.Fprintln(w)
	case *Person:
		fmt.Fprintf(w, `Person(%s, "%s", "%s")`, v.ID(), v.name, v.description)
		fmt.Fprintln(w)
	case *relation:
		prefix := "Rel"
		if v.direction != "" {
			prefix = fmt.Sprintf("Rel_%s", v.direction)
		}
		fmt.Fprintf(w, `%s(%s, %s, "%s", "%s")`, prefix, v.src.ID(), v.dst.ID(), v.description, strings.Join(v.technologies, ","))
		fmt.Fprintln(w)
	case *systemBoundary:
		fmt.Fprintf(w, `System_Boundary(%s, "%s") {`, v.ID(), v.name)
		fmt.Fprintln(w)
		for _, el := range v.elements {
			fmt.Fprintf(w, "\t")
			if err := plantUML(ctx, w, el); err != nil {
				return err
			}
		}
		fmt.Fprintln(w, "}")
	case *System:
		fmt.Fprintf(w, `System(%s, "%s", "%s")`, v.ID(), v.name, v.description)
		fmt.Fprintln(w)
	default:
		return fmt.Errorf("cannot create plantuml: invalid item type: %T", el)
	}

	return nil
}
