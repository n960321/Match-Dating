package model

// 性別
type Gender int

const (
	GenderUnknow Gender = iota // 未知
	GenderMale                 // 男生
	GenderFemale               // 女生
)
