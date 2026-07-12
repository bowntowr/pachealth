package runner

import (
	"context"
	"sync"
	"time"

	"github.com/bowntowr/pachealth/internal/check"
)

type Runner struct {
	checks  []check.Checker
	Timeout time.Duration
}

func New(checks []check.Checker) *Runner {
	return &Runner{
		checks:  checks,
		Timeout: 30 * time.Second,
	}
}

func (r *Runner) Run(ctx context.Context) ([]check.Result, check.Status) {
	ctx, cancel := context.WithTimeout(ctx, r.Timeout)
	defer cancel()

	results := make([]check.Result, len(r.checks))
	var wg sync.WaitGroup

	for i, c := range r.checks {
		i, c := i, c
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = c.Run(ctx)
		}()
	}

	wg.Wait()

	status := check.StatusOK
	for _, res := range results {
		if res.Status > status {
			status = res.Status
		}
	}

	return results, status
}
