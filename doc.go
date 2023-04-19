// Package c4 provides an API for describing systems following the C4 model
// (https://c4model.com/).  That description can then be used to produce
// C4-enabled PlantUML for generating diagrams.
//
// The API of c4 has been designed to encourage re-use, making it much easier to
// build multiple diagrams at the same level as well as diagrams at varying
// levels of "resolution" using common system/container/component definitions.
// Each diagram retains its own set of relations between the elements, so you
// can adequately describe the nature of how elements interact based on what
// slice of the overall system you are currently diagramming.
//
// At a basic level, a c4 definition consists of a couple high-level steps:
//
//  1. Declare the elements of your architecture.
//  2. Create a new diagram.
//  3. Add the relevant elements to the diagram.
//  4. Describe the relations between the elements.
//
// # Declaring Elements
//
// You declare c4 elements using the provided constructors e.g. NewPerson,
// NewSystem, etc.
//
//	personalBankingCustomer, _ := c4.NewPerson(ctx, "personalBankingCustomer", c4.PersonArgs{
//		Name:        "Personal Banking Customer",
//		Description: "A customer of the bank with personal bank accounts.",
//	})
//
//	internetBankingSystem, _ := c4.NewSystem(ctx, "internetBankingSystem", c4.SystemArgs{
//		Name:        "Internet Banking System",
//		Description: "Allows customers to view information about their bank accounts and make payments.",
//	})
//
// Every c4 constructor takes a context, an identifier, and a set of arguments
// for creating the element. The identifier must be unique among all of the
// elements in your diagram or PlantUML will only show one of the elements and
// its relations. If you intend to use the same element declaration in multiple
// diagrams, you will need to take care that identifiers don't clash in any of
// them. A best practice to help avoid any issues is to give elements meaningful
// identifiers that are unique across all of the elements in your c4 program.
// This will prevent any difficult to diagnose issues as your diagrams grow and
// change.
//
// # Creating Diagrams
//
// Creating diagrams is also fairly straightforward. The only required argument
// is the name of the diagram which, by default, PlantUML will use as the file
// name.
//
//	d, _ := c4.NewDiagram(ctx, "Demo")
//
// Optionally, you can pass in any number of configuration options that are used
// to control things like layout, etc.
//
//	d, _ := c4.NewDiagram(ctx, "Demo", c4.WithLayout(c4.LayoutLeftRight))
//
// # Adding Elements
//
// Declaring elements in a c4 program doesn't do anything by default. In order
// to use those elements in a diagram, you have to add them manually.
//
//	d.AddElement(ctx, personalBankingCustomer)
//	d.AddElement(ctx, internetBankingSystem)
//
// The biggest benefit of this manual approach is that you can add different
// sets of elements to different diagrams without having to redeclare them.
//
// # Describing Relations
//
// Once you add elements to a digram, you can export the PlantUML and see those
// elements in the produced PNG/SVG/etc., but the diagram won't have a lot of
// meaning. To be useful, you need to add relations to your diagram. These
// describe how the elements in your architecture interact and in C4 style are
// designed as a one-way directed link between a source and a destination.
// Relations also have a description and in C4 form should usually be voiced in
// the indicative tense from the perspective of the source element as in the
// following example:
//
//	d.NewRelation(ctx, c4.RelationArgs{
//		Src:         personalBankingCustomer,
//		Dst:         internetBankingSystem,
//		Description: "Views account balances and makes payments using",
//	})
//
// This relation can be expressed in prose as "the personalBankingCustomer
// [v]iews account balances and makes payments using the internetBankingSystem".
//
// Relations also take two optional parameters: technologies and direction. The
// technologies parameter lets you describe which technologies are involved in
// the interaction between elements in a relation e.g. a single-page application
// related to an API application may indicate a technology of "JSON/HTTPS". The
// direction parameter is used to give the developer a bit of control over how
// the produced diagram is laid out. You can specify that the relation should
// flow in one of the four "arrow key" directions (up, down, left, or right) and
// PlantUML will use those to position elements relative to each other within
// the physical constraints of the diagram.
//
// # Boundaries
//
// At the container level and below, you will often need to include an element
// as a boundary that other elements live within, but you don't want to see that
// element in its raw form. To accomodate this, c4 provides Boundary methods on
// any type that can be represented as such. You can then add elements to the
// boundary in the same way that you would the diagram. When you get to the
// diagram itself, you can then add the boundary rather than the parent element
// and your child elements will all be neatly grouped.
//
// In the following example, the API application is acting as a container
// boundary for the sign in and reset password controllers. We want to see the
// controller components grouped within the boundary, but we don't want to see
// the API application container itself:
//
//	apiApplication, _ := c4.NewContainer(ctx, "apiApplication", c4.ContainerArgs{
//		Name:         "API Application",
//		Description:  "Provides Internet banking functionality via a JSON/HTTPS API.",
//	})
//	signInController, _ := c4.NewComponent(ctx, "signInController", c4.ComponentArgs{
//		Name:         "Sign In Controller",
//		Description:  "Allows users to sign in to the Internet Banking System.",
//		Technologies: []string{"Spring MVC Rest Controller"},
//	})
//	resetPasswordController, _ := c4.NewComponent(ctx, "resetPasswordController", c4.ComponentArgs{
//		Name:         "Reset Password Controller",
//		Description:  "Allows users to reset their passwords with a single use URL.",
//		Technologies: []string{"Spring MVC Rest Controller"},
//	})
//
//	// ...
//
//	apiApplicationBoundary := apiApplication.Boundary()
//	apiApplicationBoundary.AddElement(ctx, signInController)
//	apiApplicationBoundary.AddElement(ctx, resetPasswordController)
//
//	// ...
//
//	d.AddElement(ctx, apiApplicationBoundary)
//
// # Theming
//
// By default, diagrams are styled using a default theme designed to be neutral
// and with enough contrast to be accessible. If you would like to style your
// diagram differently, for instance to match branding guidelines, you can do
// so by providing a theme object using the WithTheme option.
//
// If you do provide a theme, it is important that you fully specify it by
// providing a value for all available options. Since custom themes are not
// merged with the defaults provided by this package, any values not specified
// will pick up the defaults of the C4 extension to PlantUML.
//
//	d, _ := c4.NewDiagram(ctx, "Example", c4.WithTheme(c4.Theme{
//		System: c4.Palette{
//			BackgroundColor: "red",
//			FontColor:       "white",
//		},
//		Container: c4.Palette{
//			BackgroundColor: "blue",
//			FontColor:       "orange",
//		},
//		Component: c4.Palette{
//			BackgroundColor: "yellow",
//			FontColor:       "black",
//		},
//		Person: c4.Palette{
//			BackgroundColor: "green",
//			FontColor:       "grey",
//		},
//	}))
//
// As a convenenience, if you are only interested in tweaking a few values of
// the default theme, you can do this simply by obtaining a copy of the default
// with the DefaultTheme function and setting those values explicitly:
//
//	theme := c4.DefaultTheme()
//	theme.System.BackgroundColor = "green"
//	theme.Person.FontColor = "red"
//	d, _ := c4.NewDiagram(ctx, "Example", c4.WithTheme(theme))
package c4
