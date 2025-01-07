package dao

import (
	"sync"
)

type VoteData map[string]map[string]int32

var (
	vote_data VoteData
	mu        sync.Mutex
)

func loadData() {
	vote_data = VoteData{
		"Jason Ye": {
			"Doris Ding": 0,
			"Sam Fang":   0,
		},
	}
}

func init() {
	loadData()
}

// ClearVote 清除Vote
func (imp *VoteInterfaceImp) ClearVote() error {
	mu.Lock()
	defer mu.Unlock()

	loadData()
	return nil
}

// UpsertVote 更新/写入counter
func (imp *VoteInterfaceImp) UpsertVote(data *VoteData) error {
	mu.Lock()
	defer mu.Unlock()

	for k, items := range *data {
		for v, _ := range items {
			vote_data[k][v] += 1
		}
	}

	return nil
}

// GetVote 查询Vote
func (imp *VoteInterfaceImp) GetVote() (*VoteData, error) {
	mu.Lock()
	defer mu.Unlock()

	return &vote_data, nil
}
