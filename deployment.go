package c4

import "context"

// Property represents a key/value pair decribing an aspect of a deployment
// node.
type Property struct {
	Name  string
	Value string
}

// DeploymentNodeArgs describes the parameters available for configuring a
// container.
type DeploymentNodeArgs struct {
	Name        string
	Type        string
	Description string
	Properties  []Property
	Elements    []Element
}

// MustNewDeploymentNode is the same as NewDeploymentNode, but panics on any
// error.
func MustNewDeploymentNode(ctx context.Context, id string, args DeploymentNodeArgs) *DeploymentNode {
	n, err := NewDeploymentNode(ctx, id, args)
	if err != nil {
		panic(err)
	}
	return n
}

// NewDeploymentNode constructs a deployment node that can be used in a Diagram.
func NewDeploymentNode(ctx context.Context, id string, args DeploymentNodeArgs) (*DeploymentNode, error) {
	n := &DeploymentNode{
		id:          id,
		name:        args.Name,
		nodeType:    args.Type,
		description: args.Description,
		properties:  args.Properties,
		elements:    args.Elements,
	}
	return n, nil
}

// DeploymentNode represents a C4 deployment node
// (https://c4model.com/#DeploymentDiagram), which the C4 documentation
// describes as "where an instance of a software system/container is running"
// e.g. a virtual machine, container, etc.
type DeploymentNode struct {
	id          string
	name        string
	nodeType    string
	description string
	properties  []Property
	elements    []Element
}

func (dn *DeploymentNode) ID() string {
	return dn.id
}

func (dn *DeploymentNode) AddElement(ctx context.Context, el Element) {
	dn.elements = append(dn.elements, el)
}
