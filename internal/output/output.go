package output

import "github.com/bowntowr/pachealth/internal/check"

type Formatter interface {
	Format(results []check.Result, summary check.Status) (string, error)
}
