// A demonstration of a component diagram based on
// https://c4model.com/#ContainerDiagram.
package main

import (
	"context"
	"os"

	"github.com/haleyrc/c4"
)

func main() {
	ctx := context.Background()

	personalBankingCustomer, _ := c4.NewPerson(ctx, "personalBankingCustomer", c4.PersonArgs{
		Name:        "Personal Banking Customer",
		Description: "A customer of the bank with personal bank accounts.",
	})

	internetBankingSystem, _ := c4.NewSystem(ctx, "internetBankingSystem", c4.SystemArgs{
		Name:        "Internet Banking System",
		Description: "Allows customers to view information about their bank accounts and make payments.",
	})
	emailSystem, _ := c4.NewSystem(ctx, "emailSystem", c4.SystemArgs{
		Name:        "Email System",
		Description: "The internal Microsoft Exchange e-mail system.",
	})
	mainframeBankingSystem, _ := c4.NewSystem(ctx, "mainframeBankingSystem", c4.SystemArgs{
		Name:        "Mainframe Banking System",
		Description: "Stores all of the core banking information about customers, accounts, transactions, etc.",
	})

	webApplication, _ := c4.NewContainer(ctx, "webApplication", c4.ContainerArgs{
		Name:         "Web Application",
		Description:  "Delivers the static content and the Internet banking single page application.",
		Technologies: []string{"Java", "Spring MVC"},
	})
	singlePageApplication, _ := c4.NewContainer(ctx, "singlePageApplication", c4.ContainerArgs{
		Name:         "Single-Page Application",
		Description:  "Provides all of the Internet banking functionality to customers via their web browser.",
		Technologies: []string{"Javascript", "Angular"},
	})
	mobileApp, _ := c4.NewContainer(ctx, "mobileApp", c4.ContainerArgs{
		Name:         "Mobile App",
		Description:  "Provides a limited subset of the Internet banking functionality to customers via their mobile device.",
		Technologies: []string{"Xamarin"},
	})
	apiApplication, _ := c4.NewContainer(ctx, "apiApplication", c4.ContainerArgs{
		Name:         "API Application",
		Description:  "Provides Internet banking functionality via a JSON/HTTPS API.",
		Technologies: []string{"Java", "Spring MVC"},
	})
	database, _ := c4.NewDatabase(ctx, "database", c4.DatabaseArgs{
		Name:         "Database",
		Description:  "Stores user registration information, hashed authentication credentials, access logs, etc.",
		Technologies: []string{"Oracle Database Schema"},
	})

	ibsSystemBoundary := internetBankingSystem.Boundary()
	ibsSystemBoundary.AddElement(ctx, webApplication)
	ibsSystemBoundary.AddElement(ctx, singlePageApplication)
	ibsSystemBoundary.AddElement(ctx, mobileApp)
	ibsSystemBoundary.AddElement(ctx, database)
	ibsSystemBoundary.AddElement(ctx, apiApplication)

	d, _ := c4.NewDiagram(ctx, "Containers")

	d.AddElement(ctx, personalBankingCustomer)
	d.AddElement(ctx, ibsSystemBoundary)
	d.AddElement(ctx, emailSystem)
	d.AddElement(ctx, mainframeBankingSystem)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          personalBankingCustomer,
			Dst:          webApplication,
			Description:  "Visits bigbank.com/ib using",
			Technologies: []string{"HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         personalBankingCustomer,
			Dst:         singlePageApplication,
			Description: "Views account balances and makes payments using",
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         personalBankingCustomer,
			Dst:         mobileApp,
			Description: "Views account balances and makes payments using",
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         webApplication,
			Dst:         singlePageApplication,
			Description: "Delivers to the customer's web browser",
		},
		c4.WithDirection(c4.DirectionRight),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          singlePageApplication,
			Dst:          apiApplication,
			Description:  "Makes API calls to",
			Technologies: []string{"JSON/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          mobileApp,
			Dst:          apiApplication,
			Description:  "Makes API calls to",
			Technologies: []string{"JSON/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          apiApplication,
			Dst:          database,
			Description:  "Reads from and writes to",
			Technologies: []string{"SQL/TCP"},
		},
		c4.WithDirection(c4.DirectionLeft),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         apiApplication,
			Dst:         emailSystem,
			Description: "Sends e-mail using",
		},
		c4.WithDirection(c4.DirectionUp),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          apiApplication,
			Dst:          mainframeBankingSystem,
			Description:  "Makes API calls to",
			Technologies: []string{"XML/HTTPS"},
		},
		c4.WithDirection(c4.DirectionRight),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         emailSystem,
			Dst:         personalBankingCustomer,
			Description: "Sends e-mails to",
		},
		c4.WithDirection(c4.DirectionUp),
	)

	if err := d.PlantUML(ctx, os.Stdout); err != nil {
		panic(err)
	}

}
