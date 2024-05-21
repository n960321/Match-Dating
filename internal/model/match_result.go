package model

type MatchResult struct {
	Profile    *Profile   `json:"Profile"`
	Candidates []*Profile `json:"Candidates"`
}
