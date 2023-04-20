// A gallery of the available elements.
package main

import (
	"context"
	"os"

	"github.com/haleyrc/c4"
)

func main() {
	ctx := context.Background()

	d, _ := c4.NewDiagram(ctx, "Gallery")

	internalSystem, _ := c4.NewSystem(ctx, "internalSystem", c4.SystemArgs{
		Name:        "Internal System",
		Description: "Optional Description",
		External:    false,
	})
	d.AddElement(ctx, internalSystem)

	internalPerson, _ := c4.NewPerson(ctx, "internalPerson", c4.PersonArgs{
		Name:        "Internal Person",
		Description: "Optional Description",
		External:    false,
	})
	d.AddElement(ctx, internalPerson)

	internalContainer, _ := c4.NewContainer(ctx, "internalContainer", c4.ContainerArgs{
		Name:         "Internal Container",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     false,
	})
	d.AddElement(ctx, internalContainer)

	internalDatabase, _ := c4.NewDatabase(ctx, "internalDatabase", c4.DatabaseArgs{
		Name:         "Internal Database",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     false,
	})
	d.AddElement(ctx, internalDatabase)

	internalComponent, _ := c4.NewComponent(ctx, "internalComponent", c4.ComponentArgs{
		Name:         "Internal Component",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     false,
	})
	d.AddElement(ctx, internalComponent)

	externalSystem, _ := c4.NewSystem(ctx, "externalSystem", c4.SystemArgs{
		Name:        "External System",
		Description: "Optional Description",
		External:    true,
	})
	d.AddElement(ctx, externalSystem)

	externalPerson, _ := c4.NewPerson(ctx, "externalPerson", c4.PersonArgs{
		Name:        "External Person",
		Description: "Optional Description",
		External:    true,
	})
	d.AddElement(ctx, externalPerson)

	externalContainer, _ := c4.NewContainer(ctx, "externalContainer", c4.ContainerArgs{
		Name:         "External Container",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     true,
	})
	d.AddElement(ctx, externalContainer)

	externalDatabase, _ := c4.NewDatabase(ctx, "externalDatabase", c4.DatabaseArgs{
		Name:         "External Database",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     true,
	})
	d.AddElement(ctx, externalDatabase)

	externalComponent, _ := c4.NewComponent(ctx, "externalComponent", c4.ComponentArgs{
		Name:         "External Component",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     true,
	})
	d.AddElement(ctx, externalComponent)

	boundingSystem, _ := c4.NewSystem(ctx, "boundingSystem", c4.SystemArgs{
		Name:        "Bounding System",
		Description: "Optional Description",
		External:    false,
	})
	boundedContainer, _ := c4.NewContainer(ctx, "boundedContainer", c4.ContainerArgs{
		Name:         "Bounded Container",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     false,
	})
	systemBoundary := boundingSystem.Boundary()
	systemBoundary.AddElement(ctx, boundedContainer)
	d.AddElement(ctx, systemBoundary)

	boundingContainer, _ := c4.NewContainer(ctx, "boundingContainer", c4.ContainerArgs{
		Name:        "Bounding Container",
		Description: "Optional Description",
		External:    false,
	})
	boundedComponent, _ := c4.NewComponent(ctx, "boundedComponent", c4.ComponentArgs{
		Name:         "Bounded Component",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     false,
	})
	containerBoundary := boundingContainer.Boundary()
	containerBoundary.AddElement(ctx, boundedComponent)
	d.AddElement(ctx, containerBoundary)

	d.PlantUML(ctx, os.Stdout)
}
