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
		prefix := "Component"
		if v.external {
			prefix += "_Ext"
		}
		technologies := strings.Join(v.technologies, ", ")
		fmt.Fprintf(w, `%s(%s, "%s", "%s", "%s")`, prefix, v.ID(), v.name, technologies, v.description)
		fmt.Fprintln(w)
	case *Container:
		prefix := "Container"
		if v.external {
			prefix += "_Ext"
		}
		technologies := strings.Join(v.technologies, ", ")
		fmt.Fprintf(w, `%s(%s, "%s", "%s", "%s")`, prefix, v.ID(), v.name, technologies, v.description)
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
	case *EnterpriseBoundary:
		fmt.Fprintf(w, `Enterprise_Boundary(%s, "%s") {`, v.ID(), v.name)
		fmt.Fprintln(w)
		for _, el := range v.elements {
			fmt.Fprintf(w, "\t")
			if err := plantUML(ctx, w, el); err != nil {
				return err
			}
		}
		fmt.Fprintln(w, "}")
	case *Database:
		prefix := "ContainerDb"
		if v.external {
			prefix += "_Ext"
		}
		technologies := strings.Join(v.technologies, ", ")
		fmt.Fprintf(w, `%s(%s, "%s", "%s", "%s")`, prefix, v.ID(), v.name, technologies, v.description)
		fmt.Fprintln(w)
	case *DeploymentNode:
		for _, property := range v.properties {
			fmt.Fprintf(w, `AddProperty("%s", "%s")`, property.Name, property.Value)
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, `Deployment_Node(%s, "%s", "%s", "%s") {`, v.id, v.name, v.nodeType, v.description)
		fmt.Fprintln(w)
		for _, el := range v.elements {
			fmt.Fprintf(w, "\t")
			if err := plantUML(ctx, w, el); err != nil {
				return err
			}
		}
		fmt.Fprintln(w, `}`)
	case *Person:
		prefix := "Person"
		if v.external {
			prefix += "_Ext"
		}
		fmt.Fprintf(w, `%s(%s, "%s", "%s")`, prefix, v.ID(), v.name, v.description)
		fmt.Fprintln(w)
	case *Queue:
		prefix := "ContainerQueue"
		if v.external {
			prefix += "_Ext"
		}
		technologies := strings.Join(v.technologies, ", ")
		fmt.Fprintf(w, `%s(%s, "%s", "%s", "%s")`, prefix, v.ID(), v.name, technologies, v.description)
		fmt.Fprintln(w)
	case *relation:
		prefix := "Rel"
		if v.direction != "" {
			prefix = fmt.Sprintf("Rel_%s", v.direction)
		}
		fmt.Fprintf(w, `%s(%s, %s, "%s", "%s")`, prefix, v.src.ID(), v.dst.ID(), v.description, strings.Join(v.technologies, ","))
		fmt.Fprintln(w)
	case *step:
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
		prefix := "System"
		if v.external {
			prefix += "_Ext"
		}
		fmt.Fprintf(w, `%s(%s, "%s", "%s")`, prefix, v.ID(), v.name, v.description)
		fmt.Fprintln(w)
	default:
		return fmt.Errorf("cannot create plantuml: invalid item type: %T", el)
	}

	return nil
}
