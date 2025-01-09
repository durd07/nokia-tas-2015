package model

import "time"

// CounterModel 计数器模型
type CounterModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Count     int32     `gorm:"column:count" json:"count"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

type MemberModel struct {
	Name string `gorm:"column:name" json:"name"`
	Vote int32  `gorm:"column:vote" json:"vote"`
}

type VoteModel struct {
	Name    string                 `gorm:"column:name" json:"name"`
	Members map[string]*MemberModel `gorm:"column:members" json:"members"`
}

type VoteData map[string]*VoteModel
