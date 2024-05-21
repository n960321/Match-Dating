package model

type QueryResult struct {
	Infos []*DatingInfo `json:"Infos"`
}

type DatingInfo struct {
	Male   *Profile `json:"Male"`
	Female *Profile `json:"Female"`
}
