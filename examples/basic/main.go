// A demonstration of a minimal c4 program.
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

	d, _ := c4.NewDiagram(ctx, "Basic")
	d.AddElement(ctx, internetBankingSystem)
	d.PlantUML(ctx, os.Stdout)
}
