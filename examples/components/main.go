// A demonstration of a component diagram based on
// https://c4model.com/#ComponentDiagram.
package main

import (
	"context"
	"os"

	"github.com/haleyrc/c4"
)

func main() {
	ctx := context.Background()

	mainframeBankingSystem, _ := c4.NewSystem(ctx, "mainframeBankingSystem", c4.SystemArgs{
		Name:        "Mainframe Banking System",
		Description: "Stores all of the core banking information about customers, accounts, transactions, etc.",
	})
	emailSystem, _ := c4.NewSystem(ctx, "emailSystem", c4.SystemArgs{
		Name:        "Email System",
		Description: "The internal Microsoft Exchange e-mail system.",
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

	signInController, _ := c4.NewComponent(ctx, "signInController", c4.ComponentArgs{
		Name:         "Sign In Controller",
		Description:  "Allows users to sign in to the Internet Banking System.",
		Technologies: []string{"Spring MVC Rest Controller"},
	})
	securityComponent, _ := c4.NewComponent(ctx, "securityComponent", c4.ComponentArgs{
		Name:         "Security Component",
		Description:  "Provides functionality related to signing in, changing passwords, etc.",
		Technologies: []string{"Spring Bean"},
	})
	resetPasswordController, _ := c4.NewComponent(ctx, "resetPasswordController", c4.ComponentArgs{
		Name:         "Reset Password Controller",
		Description:  "Allows users to reset their passwords with a single use URL.",
		Technologies: []string{"Spring MVC Rest Controller"},
	})
	emailComponent, _ := c4.NewComponent(ctx, "emailComponent", c4.ComponentArgs{
		Name:         "E-mail Component",
		Description:  "Sends e-mails to users.",
		Technologies: []string{"Spring Bean"},
	})
	accountsSummaryController, _ := c4.NewComponent(ctx, "accountsSummaryController", c4.ComponentArgs{
		Name:         "Accounts Summary Controller",
		Description:  "Provides customers with a summary of their bank accounts.",
		Technologies: []string{"Spring MVC Rest Controller"},
	})
	mainframeBankingSystemFacade, _ := c4.NewComponent(ctx, "mainframeBankingSystemFacade", c4.ComponentArgs{
		Name:         "Mainframe Banking System Facade",
		Description:  "A facade onto the mainframe banking system.",
		Technologies: []string{"Spring Bean"},
	})

	apiApplicationBoundary := apiApplication.Boundary()
	apiApplicationBoundary.AddElement(ctx, signInController)
	apiApplicationBoundary.AddElement(ctx, resetPasswordController)
	apiApplicationBoundary.AddElement(ctx, accountsSummaryController)
	apiApplicationBoundary.AddElement(ctx, securityComponent)
	apiApplicationBoundary.AddElement(ctx, emailComponent)
	apiApplicationBoundary.AddElement(ctx, mainframeBankingSystemFacade)

	d, _ := c4.NewDiagram(ctx, "Components")

	d.AddElement(ctx, singlePageApplication)
	d.AddElement(ctx, mobileApp)
	d.AddElement(ctx, apiApplicationBoundary)
	d.AddElement(ctx, database)
	d.AddElement(ctx, emailSystem)
	d.AddElement(ctx, mainframeBankingSystem)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          singlePageApplication,
			Dst:          signInController,
			Description:  "Makes API calls to",
			Technologies: []string{"JSON/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          singlePageApplication,
			Dst:          resetPasswordController,
			Description:  "Makes API calls to",
			Technologies: []string{"JSON/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          singlePageApplication,
			Dst:          accountsSummaryController,
			Description:  "Makes API calls to",
			Technologies: []string{"JSON/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          mobileApp,
			Dst:          signInController,
			Description:  "Makes API calls to",
			Technologies: []string{"JSON/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          mobileApp,
			Dst:          resetPasswordController,
			Description:  "Makes API calls to",
			Technologies: []string{"JSON/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          mobileApp,
			Dst:          accountsSummaryController,
			Description:  "Makes API calls to",
			Technologies: []string{"JSON/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         signInController,
			Dst:         securityComponent,
			Description: "Uses",
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         resetPasswordController,
			Dst:         emailComponent,
			Description: "Uses",
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         resetPasswordController,
			Dst:         securityComponent,
			Description: "Uses",
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         accountsSummaryController,
			Dst:         mainframeBankingSystemFacade,
			Description: "Uses",
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          securityComponent,
			Dst:          database,
			Description:  "Reads from and writes to",
			Technologies: []string{"SQL/TCP"},
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         emailComponent,
			Dst:         emailSystem,
			Description: "Sends e-mail using",
		},
		c4.WithDirection(c4.DirectionDown),
	)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          mainframeBankingSystemFacade,
			Dst:          mainframeBankingSystem,
			Description:  "Makes API calls to",
			Technologies: []string{"XML/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)

	if err := d.PlantUML(ctx, os.Stdout); err != nil {
		panic(err)
	}
}
