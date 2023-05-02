package main

import (
	"context"
	"os"

	"github.com/haleyrc/c4"
)

func main() {
	ctx := context.Background()

	d, _ := c4.NewDiagram(ctx, "Dynamic Diagram", c4.WithLegend(), c4.HideElementTypes())

	c4db, _ := c4.NewDatabase(ctx, "c4", c4.DatabaseArgs{
		Name:         "Database",
		Description:  "Stores user registration information, hashed authentication credentials, access logs, etc.",
		Technologies: []string{"Relational Database Schema"},
	})

	c1, _ := c4.NewContainer(ctx, "c1", c4.ContainerArgs{
		Name:         "Single-Page Application",
		Description:  "Provides all of the Internet banking functionality to customers via their web browser.",
		Technologies: []string{"Javascript and Angular"},
	})

	b, _ := c4.NewContainer(ctx, "b", c4.ContainerArgs{
		Name: "API Application",
	})

	bBoundary := b.Boundary()
	c3, _ := c4.NewComponent(ctx, "c3", c4.ComponentArgs{
		Name:         "Security Component",
		Description:  "Provides functionality Related to signing in, changing passwords, etc.",
		Technologies: []string{"Spring Bean"},
	})
	c2, _ := c4.NewComponent(ctx, "c2", c4.ComponentArgs{
		Name:         "Sign In Controller",
		Description:  "Allows users to sign in to the Internet Banking System.",
		Technologies: []string{"Spring MVC Rest Controller"},
	})
	bBoundary.AddElement(ctx, c3)
	bBoundary.AddElement(ctx, c2)

	d.AddElement(ctx, c4db)
	d.AddElement(ctx, c1)
	d.AddElement(ctx, bBoundary)

	d.AddSteps(ctx,
		c4.StepArgs{
			Src:          c1,
			Dst:          c2,
			Description:  "Submits credentials to",
			Technologies: []string{"JSON/HTTPS"},
			Direction:    c4.Direction(c4.DirectionRight),
		},
		c4.StepArgs{
			Src:         c2,
			Dst:         c3,
			Description: "Calls isAuthenticated() on",
		},
		c4.StepArgs{
			Src:          c3,
			Dst:          c4db,
			Description:  "select * from users where username = ?",
			Technologies: []string{"JDBC"},
			Direction:    c4.Direction(c4.DirectionRight),
		},
	)

	if err := d.PlantUML(ctx, os.Stdout); err != nil {
		panic(err)
	}
}
