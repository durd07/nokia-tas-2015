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

type VoteData struct {
	EmployeeManagerMapping map[string]string
	Data map[string]*VoteModel
}

func InitVoteData(data map[string][]string) VoteData {
	voteData := VoteData{}

	for k, v := range data {
		if _, exists := voteData.Data[k]; !exists {
			voteData.Data[k] = &VoteModel{
				Name: k,
				Members: make(map[string]*MemberModel),
			}
		}

		for _, x := range v {
			voteData.EmployeeManagerMapping[x] = k
			voteData.Data[k].Members[x] = &MemberModel{
				Name: x,
				Vote: 0,
			}
		}
	}
	return voteData
}
