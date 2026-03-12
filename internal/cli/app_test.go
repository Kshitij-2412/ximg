package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunVersion(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	app := NewWithWriters(&stdout, &stderr)

	if err := app.Run([]string{"version"}); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if got := stdout.String(); !strings.Contains(got, "ximg 0.1.0") {
		t.Fatalf("unexpected stdout: %q", got)
	}
}

func TestRunDiffRequiresTwoImages(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	app := NewWithWriters(&stdout, &stderr)

	err := app.Run([]string{"diff", "only-one-image"})
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if got := err.Error(); got != "usage: ximg diff <image-a> <image-b>" {
		t.Fatalf("unexpected error: %q", got)
	}
}
