# c4

> **NOTICE:** This repository is presented as-is. No guarantees are made with regards to quality, suitability, or backwards-compatibility. Use at your own risk.

[![Build Status](https://github.com/haleyrc/c4/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/haleyrc/c4/actions?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/haleyrc/c4?status.svg)](https://pkg.go.dev/github.com/haleyrc/c4?tab=doc)

A library for describing software systems in Go following the [C4 model](https://c4model.com/).

Before getting started, it's important to note that this library is not intended as a complete implementation of the C4 extension for PlantUML. Functionality is only added as I have a use-case for it. See the [References](#references) section for links to alternatives.

## Install

```
$ go get -u github.com/haleyrc/c4
```

## Usage

> This section outlines the general workflow using the `c4` package. For detailed instructions on implementation, see the package documentation.

When generating diagrams with PlantUML, you generally write a specification in
the PlantUML "language" in a plain text file and then use either the [online server](https://www.plantuml.com/plantuml/) or the CLI to convert that specification to a PNG/SVG/etc. This package focuses exclusively on the first portion of this process by allowing you to write your specification in Go rather than the PlantUML DSL. This has a number of benefits including test, error handling, and re-use. If you are familiar with infrastructure-as-code tooling, this package draws a lot of inspiration from Pulumi where native PlantUML is more akin to writing Terraform configs in raw HCL.

Given the minimal program from the [basic example](./examples/basic/main.go), paraphrased here:

```go
ctx := context.Background()

internetBankingSystem, _ := c4.NewSystem(ctx, "internetBankingSystem", c4.SystemArgs{
    Name:        "Internet Banking System",
    Description: "Allows customers to view information about their bank accounts and make payments.",
})

d, _ := c4.NewDiagram(ctx, "Basic")
d.AddElement(ctx, internetBankingSystem)
d.PlantUML(ctx, os.Stdout)
```

You can generate a PNG by running the following:

```bash
$ go run ./examples/basic | java -jar plantuml.jar -p > basic.png
```

You are, of course, free to output your PlantUML specification to a file and pass that to the PlantUML CLI as an argument. The `c4` package doesn't make any real assumptions about how you are getting from the Go world to the PlantUML world.

## TODO

- [ ] Deployment diagram support
- [ ] Dynamic diagram support

### Future

- [ ] Single static definition e.g. the ability to write a single monolithic description of your systems, their containers, and their components, and then use those in building diagrams by only specififying the nodes you want.

## References

- [go-structurizer - A library for auto-generating C4 diagrams from Go applications](https://github.com/krzysztofreczek/go-structurizr)
