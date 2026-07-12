package check

import "context"

type Status int

const (
	StatusOK Status = iota
	StatusWarning
	StatusError
)

type Result struct {
	Name    string
	Status  Status
	Summary string
	Details []string
	Err     error
}

type Checker interface {
	ID() string
	Run(ctx context.Context) Result
}
