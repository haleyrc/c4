package main

import (
	"context"
	"os"

	"github.com/haleyrc/c4"
)

func main() {
	ctx := context.Background()

	api, _ := c4.NewContainer(ctx, "api", c4.ContainerArgs{
		Name:         "API Application",
		Description:  "Provides Internet Banking functionality via a JSON/HTTPS API.",
		Technologies: []string{"Java", "Spring MVC"},
	})

	apache, _ := c4.NewDeploymentNode(ctx, "apache", c4.DeploymentNodeArgs{
		Name:        "Apache Tomcat",
		Type:        "Apache Tomcat 8.x",
		Description: "An open source Java EE web server.",
		Properties: []c4.Property{
			{Name: "Java Version", Value: "8"},
			{Name: "Xmx", Value: "512M"},
			{Name: "Xms", Value: "1024M"},
		},
		Elements: []c4.Element{api},
	})

	dn, _ := c4.NewDeploymentNode(ctx, "dn", c4.DeploymentNodeArgs{
		Name:        "bigbank-api***\tx8",
		Type:        "Ubuntu 16.04 LTS",
		Description: "A web server residing in the web server farm, accessed via F5 BIG-IP LTMs.",
		Properties:  []c4.Property{{Name: "Location", Value: "London and Reading"}},
		Elements:    []c4.Element{apache},
	})

	db, _ := c4.NewDatabase(ctx, "db", c4.DatabaseArgs{
		Name:         "Database",
		Description:  "Stores user registration information, hashed authentication credentials, access logs, etc.",
		Technologies: []string{"Relational Database Schema"},
	})

	oracle, _ := c4.NewDeploymentNode(ctx, "oracle", c4.DeploymentNodeArgs{
		Name:        "Oracle - Primary",
		Type:        "Oracle 12c",
		Description: "The primary, live database server.",
		Elements:    []c4.Element{db},
	})

	bigbankdb01, _ := c4.NewDeploymentNode(ctx, "bigbankdb01", c4.DeploymentNodeArgs{
		Name:        "bigbank-db01",
		Type:        "Ubuntu 16.04 LTS",
		Description: "The primary database server.",
		Properties:  []c4.Property{{Name: "Location", Value: "London"}},
		Elements:    []c4.Element{oracle},
	})

	db2, _ := c4.NewDatabase(ctx, "db2", c4.DatabaseArgs{
		Name:         "Database",
		Description:  "Stores user registration information, hashed authentication credentials, access logs, etc.",
		Technologies: []string{"Relational Database Schema"},
	})

	oracle2, _ := c4.NewDeploymentNode(ctx, "oracle2", c4.DeploymentNodeArgs{
		Name:        "Oracle - Secondary",
		Type:        "Oracle 12c",
		Description: "A secondary, standby database server, used for failover purposes only.",
		Elements:    []c4.Element{db2},
	})

	bigbankdb02, _ := c4.NewDeploymentNode(ctx, "bigbankdb02", c4.DeploymentNodeArgs{
		Name:        "bigbank-db02",
		Type:        "Ubuntu 16.04 LTS",
		Description: "The secondary database server.",
		Properties:  []c4.Property{{Name: "Location", Value: "Reading"}},
		Elements:    []c4.Element{oracle2},
	})

	web, _ := c4.NewContainer(ctx, "web", c4.ContainerArgs{
		Name:         "Web Application",
		Description:  "Delivers the static content and the Internet Banking single page application.",
		Technologies: []string{"Java", "Spring MVC"},
	})

	apache2, _ := c4.NewDeploymentNode(ctx, "apache2", c4.DeploymentNodeArgs{
		Name:        "Apache Tomcat",
		Type:        "Apache Tomcat 8.x",
		Description: "An open source Java EE web server.",
		Properties: []c4.Property{
			{Name: "Java Version", Value: "8"},
			{Name: "Xmx", Value: "512M"},
			{Name: "Xms", Value: "1024M"},
		},
		Elements: []c4.Element{web},
	})

	bb2, _ := c4.NewDeploymentNode(ctx, "bb2", c4.DeploymentNodeArgs{
		Name:        "bigbank-web***\tx4",
		Type:        "Ubuntu 16.04 LTS",
		Description: "A web server residing in the web server farm, accessed via F5 BIG-IP LTMs.",
		Properties:  []c4.Property{{Name: "Location", Value: "London and Reading"}},
		Elements:    []c4.Element{apache2},
	})

	plc, _ := c4.NewDeploymentNode(ctx, "plc", c4.DeploymentNodeArgs{
		Name:        "Live",
		Type:        "Big Bank plc",
		Description: "Big bank plc data center",
		Elements:    []c4.Element{dn, bigbankdb01, bigbankdb02, bb2},
	})

	mobile, _ := c4.NewContainer(ctx, "mobile", c4.ContainerArgs{
		Name:         "Mobile App",
		Description:  "Provides a limited subset of the Internet Banking functionality to customers via their mobile device.",
		Technologies: []string{"Xamarin"},
	})

	mob, _ := c4.NewDeploymentNode(ctx, "mob", c4.DeploymentNodeArgs{
		Name:     "Customer's mobile device",
		Type:     "Apple IOS or Android",
		Elements: []c4.Element{mobile},
	})

	spa, _ := c4.NewContainer(ctx, "spa", c4.ContainerArgs{
		Name:         "Single Page Application",
		Description:  "Provides all of the Internet Banking functionality to customers via their web browser.",
		Technologies: []string{"JavaScript", "Angular"},
	})

	browser, _ := c4.NewDeploymentNode(ctx, "browser", c4.DeploymentNodeArgs{
		Name:     "Web Browser",
		Type:     "Google Chrome, Mozilla Firefox, Apple Safari or Microsoft Edge",
		Elements: []c4.Element{spa},
	})

	comp, _ := c4.NewDeploymentNode(ctx, "comp", c4.DeploymentNodeArgs{
		Name:     "Customer's computer",
		Type:     "Microsoft Windows or Apple macOS",
		Elements: []c4.Element{browser},
	})

	d, _ := c4.NewDiagram(ctx, "Deployment Diagram")

	d.AddElement(ctx, plc)
	d.AddElement(ctx, mob)
	d.AddElement(ctx, comp)

	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          mobile,
			Dst:          api,
			Description:  "Makes API calls to",
			Technologies: []string{"json/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          spa,
			Dst:          api,
			Description:  "Makes API calls to",
			Technologies: []string{"json/HTTPS"},
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         web,
			Dst:         spa,
			Description: "Delivers to the customer's web browser",
		},
		c4.WithDirection(c4.DirectionUp),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          api,
			Dst:          db,
			Description:  "Reads from and writes to",
			Technologies: []string{"JDBC"},
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:          api,
			Dst:          db2,
			Description:  "Reads from and writes to",
			Technologies: []string{"JDBC"},
		},
		c4.WithDirection(c4.DirectionDown),
	)
	d.NewRelation(ctx,
		c4.RelationArgs{
			Src:         db,
			Dst:         db2,
			Description: "Replicates data to",
		},
		c4.WithDirection(c4.DirectionRight),
	)

	if err := d.PlantUML(ctx, os.Stdout); err != nil {
		panic(err)
	}
}
