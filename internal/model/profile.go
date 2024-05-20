package model

type Profile struct {
	ID          uint   // 唯一值
	Name        string // 名字
	Height      uint   // 身高
	Gender      Gender // 性別
	WantedDates uint   // 想要約會的次數
}
