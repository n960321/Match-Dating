package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Profile struct {
	ID          uint   `json:"ID"`          // 唯一值
	Name        string `json:"Name"`        // 名字
	Height      uint   `json:"Height"`      // 身高
	Gender      Gender `json:"Gender"`      // 性別
	WantedDates uint   `json:"WantedDates"` // 想要約會的次數
}

func (p *Profile) GetKey() string {
	return fmt.Sprintf("%d_%d", p.ID, p.Height)
}

func GetIDAndHeightFromKey(s string) (uint, uint) {
	arr := strings.Split(s, "_")
	id, _ := strconv.Atoi(arr[0])
	height, _ := strconv.Atoi(arr[1])
	return uint(id), uint(height)
}
