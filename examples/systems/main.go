// A demonstration of a component diagram based on
// https://c4model.com/#SystemContextDiagram.
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

	d, _ := c4.NewDiagram(ctx, "Systems Context")

	d.AddElement(ctx, personalBankingCustomer)
	d.AddElement(ctx, internetBankingSystem)
	d.AddElement(ctx, emailSystem)
	d.AddElement(ctx, mainframeBankingSystem)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         personalBankingCustomer,
			Dst:         internetBankingSystem,
			Description: "Views account balances and makes payments using",
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         internetBankingSystem,
			Dst:         emailSystem,
			Description: "Sends e-mail using",
		},
		c4.WithDirection(c4.DirectionRight),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         internetBankingSystem,
			Dst:         mainframeBankingSystem,
			Description: "Gets account information from and makes payments using",
		},
		c4.WithDirection(c4.DirectionDown),
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
