package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const version = "0.1.0"

// App owns CLI execution and output streams.
// Keeping this as a struct makes the code easier to test later.
type App struct {
	stdout io.Writer
	stderr io.Writer
}

// New constructs the CLI app with standard output streams.
func New() *App {
	return NewWithWriters(os.Stdout, os.Stderr)
}

// NewWithWriters lets tests inject buffers instead of using the real terminal.
func NewWithWriters(stdout, stderr io.Writer) *App {
	return &App{
		stdout: stdout,
		stderr: stderr,
	}
}

// Run routes CLI arguments to commands.
func (a *App) Run(args []string) error {
	if len(args) == 0 {
		return a.printUsage()
	}

	switch args[0] {
	case "help", "-h", "--help":
		return a.printUsage()
	case "version":
		_, err := fmt.Fprintf(a.stdout, "ximg %s\n", version)
		return err
	case "diff":
		return a.runDiff(args[1:])
	default:
		return fmt.Errorf("unknown command %q\n\n%s", args[0], usageText)
	}
}

func (a *App) printUsage() error {
	_, err := fmt.Fprint(a.stdout, usageText)
	return err
}

func (a *App) runDiff(args []string) error {
	if len(args) != 2 {
		return errors.New("usage: ximg diff <image-a> <image-b>")
	}

	report := NewPlaceholderDiffReport(args[0], args[1])

	_, err := fmt.Fprintf(
		a.stdout,
		"comparing %s -> %s\nstatus: %s\nsummary: %s\n",
		report.SourceImage,
		report.TargetImage,
		report.Status,
		report.Summary,
	)
	return err
}

const usageText = `ximg compares container images deeply.

Usage:
  ximg version
  ximg diff <image-a> <image-b>
`
