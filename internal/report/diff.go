package report

// Status describes the maturity or state of an analysis result.
type Status string

const (
	StatusPending Status = "pending"
	StatusReady   Status = "ready"
	StatusFailed  Status = "failed"
)

// DiffReport is the top-level output model for image comparison.
// It is intentionally small for now and will grow as analyzers are added.
type DiffReport struct {
	SourceImage string `json:"sourceImage"`
	TargetImage string `json:"targetImage"`
	Status      Status `json:"status"`
	Summary     string `json:"summary"`
}
