// A demonstration of theming a diagram.
package main

import (
	"context"
	"os"

	"github.com/haleyrc/c4"
)

func main() {
	ctx := context.Background()

	internetBankingSystem, _ := c4.NewSystem(ctx, "internetBankingSystem", c4.SystemArgs{
		Name:        "Internet Banking System",
		Description: "Allows customers to view information about their bank accounts and make payments.",
	})
	singlePageApplication, _ := c4.NewContainer(ctx, "singlePageApplication", c4.ContainerArgs{
		Name:         "Single-Page Application",
		Description:  "Provides all of the Internet banking functionality to customers via their web browser.",
		Technologies: []string{"Javascript", "Angular"},
	})
	signInController, _ := c4.NewComponent(ctx, "signInController", c4.ComponentArgs{
		Name:         "Sign In Controller",
		Description:  "Allows users to sign in to the Internet Banking System.",
		Technologies: []string{"Spring MVC Rest Controller"},
	})
	personalBankingCustomer, _ := c4.NewPerson(ctx, "personalBankingCustomer", c4.PersonArgs{
		Name:        "Personal Banking Customer",
		Description: "A customer of the bank with personal bank accounts.",
	})

	d, _ := c4.NewDiagram(ctx, "Theming", c4.WithTheme(c4.Theme{
		System: c4.Palette{
			BackgroundColor: "red",
			FontColor:       "white",
		},
		Container: c4.Palette{
			BackgroundColor: "blue",
			FontColor:       "orange",
		},
		Component: c4.Palette{
			BackgroundColor: "yellow",
			FontColor:       "black",
		},
		Person: c4.Palette{
			BackgroundColor: "green",
			FontColor:       "grey",
		},
	}))
	d.AddElement(ctx, internetBankingSystem)
	d.AddElement(ctx, singlePageApplication)
	d.AddElement(ctx, signInController)
	d.AddElement(ctx, personalBankingCustomer)
	d.PlantUML(ctx, os.Stdout)
}
