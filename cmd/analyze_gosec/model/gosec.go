package model

// ReportInfo this is report information
type ReportInfo struct {
	Errors       map[string][]Error `json:"Golang errors"`
	Issues       []*Issue           `json:"Issues"`
	Stats        *Metrics           `json:"Stats"`
	GosecVersion string             `json:"GosecVersion"`
}

// Issue is returned by a gosec rule if it discovers an issue with the scanned code.
type Issue struct {
	Severity     string            `json:"severity"`     // issue severity (how problematic it is)
	Confidence   string            `json:"confidence"`   // issue confidence (how sure we are we found it)
	Cwe          Cwe               `json:"cwe"`          // Cwe associated with RuleID
	RuleID       string            `json:"rule_id"`      // Human readable explanation
	What         string            `json:"details"`      // Human readable explanation
	File         string            `json:"file"`         // File name we found it in
	Code         string            `json:"code"`         // Impacted code line
	Line         string            `json:"line"`         // Line number in file
	Col          string            `json:"column"`       // Column number in line
	NoSec        bool              `json:"nosec"`        // true if the issue is nosec
	Suppressions []SuppressionInfo `json:"suppressions"` // Suppression info of the issue
}

type Metrics struct {
	NumFiles int `json:"files"`
	NumLines int `json:"lines"`
	NumNosec int `json:"nosec"`
	NumFound int `json:"found"`
}

// Error is used when there are golang errors while parsing the AST
type Error struct {
	Line   int    `json:"line"`
	Column int    `json:"column"`
	Err    string `json:"error"`
}

// Cwe defines a url link
type Cwe struct {
	ID  string
	Url string
}

type SuppressionInfo struct {
	Kind          string `json:"kind"`
	Justification string `json:"justification"`
}
