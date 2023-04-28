// A gallery of the available elements.
package main

import (
	"context"
	"flag"
	"os"

	"github.com/haleyrc/c4"
)

func main() {
	ctx := context.Background()

	cfg, err := parseCommandLine(ctx)
	if err != nil {
		panic(err)
	}

	title := "Gallery"
	opts := []c4.DiagramOption{
		c4.WithLegend(),
	}
	if cfg.Sketch {
		title = "Sketch"
		opts = append(opts, c4.AsSketch())
	}

	d, _ := c4.NewDiagram(ctx, title, opts...)

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

	internalQueue, _ := c4.NewQueue(ctx, "internalQueue", c4.QueueArgs{
		Name:         "Internal Queue",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     false,
	})
	d.AddElement(ctx, internalQueue)

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

	externalQueue, _ := c4.NewQueue(ctx, "externalQueue", c4.QueueArgs{
		Name:         "External Queue",
		Description:  "Optional Description",
		Technologies: []string{"Technology"},
		External:     false,
	})
	d.AddElement(ctx, externalQueue)

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

	enterpriseBoundary, _ := c4.NewEnterpriseBoundary(ctx, "enterpriseBoundary", c4.EnterpriseBoundaryArgs{
		Name: "Enterprise Boundary",
	})
	boundedSystem, _ := c4.NewSystem(ctx, "boundedSystem", c4.SystemArgs{
		Name:        "Bounded System",
		Description: "Optional Description",
	})
	enterpriseBoundary.AddElement(ctx, boundedSystem)
	d.AddElement(ctx, enterpriseBoundary)

	parentNode, _ := c4.NewDeploymentNode(ctx, "parentNode", c4.DeploymentNodeArgs{
		Name:        "Parent Node",
		Description: "A deployment node containing another node.",
		Properties: []c4.Property{
			{Name: "Location", Value: "New York"},
		},
	})
	childNode, _ := c4.NewDeploymentNode(ctx, "childNode", c4.DeploymentNodeArgs{
		Name:        "Child Node",
		Description: "A deployment node inside another node.",
		Properties: []c4.Property{
			{Name: "Memory", Value: "100Mb"},
			{Name: "Storage", Value: "50Gb"},
		},
	})
	childContainer, _ := c4.NewContainer(ctx, "childContainer", c4.ContainerArgs{
		Name:         "Child Container",
		Description:  "A container inside a deployment node.",
		Technologies: []string{"Technology"},
	})
	childNode.AddElement(ctx, childContainer)
	parentNode.AddElement(ctx, childNode)
	d.AddElement(ctx, parentNode)

	if err := d.PlantUML(ctx, os.Stdout); err != nil {
		panic(err)
	}
}

type Config struct {
	Sketch bool
}

func parseCommandLine(ctx context.Context) (Config, error) {
	var cfg Config

	flag.BoolVar(&cfg.Sketch, "sketch", false, "Render the gallery as a sketch")
	flag.Parse()

	return cfg, nil
}
