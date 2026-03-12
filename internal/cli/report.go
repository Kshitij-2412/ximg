package cli

import "github.com/Kshitij-2412/ximg/internal/report"

// NewPlaceholderDiffReport returns a minimal report until the real analyzers exist.
// We keep report construction separate from printing so the future CLI can support
// text, JSON, and machine-readable outputs without duplicating logic.
func NewPlaceholderDiffReport(sourceImage, targetImage string) report.DiffReport {
	return report.DiffReport{
		SourceImage: sourceImage,
		TargetImage: targetImage,
		Status:      report.StatusPending,
		Summary:     "analysis engine not implemented yet",
	}
}
