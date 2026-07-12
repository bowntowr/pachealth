package output

import (
	"fmt"
	"strings"

	"github.com/bowntowr/pachealth/internal/check"
)

type TextFormatter struct {
	Verbose bool
}

func (f *TextFormatter) Format(results []check.Result, summary check.Status) (string, error) {
	var b strings.Builder

	fmt.Fprintln(&b, "pachealth")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "System Status: %s\n\n", statusLabel(summary))

	for _, r := range results {
		indicator := symbol(r.Status)

		if r.Err != nil {
			fmt.Fprintf(&b, "%s %s (error: %v)\n", indicator, r.Name, r.Err)
		} else {
			fmt.Fprintf(&b, "%s %s\n", indicator, r.Summary)
		}

		if f.Verbose && len(r.Details) > 0 {
			for _, d := range r.Details {
				fmt.Fprintf(&b, "  %s\n", d)
			}
		}
	}

	return b.String(), nil
}

func statusLabel(s check.Status) string {
	switch s {
	case check.StatusOK:
		return "OK"
	case check.StatusWarning:
		return "ATTENTION REQUIRED"
	case check.StatusError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func symbol(s check.Status) string {
	switch s {
	case check.StatusOK:
		return "\u2714"
	case check.StatusWarning:
		return "\u26A0"
	case check.StatusError:
		return "\u2716"
	default:
		return "?"
	}
}
