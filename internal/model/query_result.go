package model

type QueryResult struct {
	Infos []*DatingInfo
}

type DatingInfo struct {
	Male   *Profile
	Female *Profile
}
