package model

import "github.com/sonatype-nexus-community/go-sona-types/ossindex/types"

// ReportInfo this is report information
type ReportInfo struct {
	Audited       []types.Coordinate    `json:"audited"`
	Excluded      []types.Vulnerability `json:"excluded"`
	Exclusions    []string              `json:"exclusions"`
	Invalid       []types.Coordinate    `json:"invalid"`
	NumAudited    int                   `json:"num_audited"`
	NumExclusions int                   `json:"num_exclusions"`
	NumVulnerable int                   `json:"num_vulnerable"`
	Version       string                `json:"version"`
	Vulnerable    []types.Coordinate    `json:"vulnerable"`
}
