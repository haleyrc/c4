package c4

import (
	"context"
)

// QueueArgs describes the parameters available for configuring a queue
// container.
type QueueArgs struct {
	// The human-readable name of the queue.
	Name string

	// A general description of the purpose of the queue.
	Description string

	// An optional list of technologies describing the queue e.g. Kafka, SQS.
	Technologies []string

	// Enables alternate styling reserved for external elements.
	External bool
}

// MustNewQueue is the same as NewQueue, but panics on any error.
func MustNewQueue(ctx context.Context, id string, args QueueArgs) *Queue {
	db, err := NewQueue(ctx, id, args)
	if err != nil {
		panic(err)
	}
	return db
}

// NewQueue constructs a queue container that can be used in a Diagram.
func NewQueue(ctx context.Context, id string, args QueueArgs) (*Queue, error) {
	c := &Queue{
		id:          id,
		name:        args.Name,
		description: args.Description,
		external:    args.External,
	}
	return c, nil
}

// Queue represents a C4 container (https://c4model.com/#ContainerDiagram)
// specifically for describing queues. The C4 documentation refers to
// queue alongside other containers, but conceptually they stand apart and we
// likewise represent them differently in the resultant diagram.
type Queue struct {
	id           string
	name         string
	description  string
	technologies []string
	external     bool
}

// ID satisfies the Element interface.
func (db *Queue) ID() string {
	return db.id
}
