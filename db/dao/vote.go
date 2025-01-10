package dao

import (
	"os"
	"sync"
	"gopkg.in/yaml.v3"
	"log"
	"wxcloudrun-golang/db/model"
)

var (
	vote_data model.VoteData
	mu        sync.Mutex
)

func loadData() {
	text, err := os.ReadFile("ntas.yaml")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var data map[string][]string
	err = yaml.Unmarshal(text, &data)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	vote_data = model.InitVoteData(data)
}

func init() {
	mu.Lock()
	defer mu.Unlock()

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
func (imp *VoteInterfaceImp) UpsertVote(data []string) error {
	mu.Lock()
	defer mu.Unlock()

	for _, person := range data {
		manager := vote_data.EmployeeManagerMapping[person]
		vote_data.Data[manager].Members[person].Vote += 1
	}

	return nil
}

// GetVote 查询Vote
func (imp *VoteInterfaceImp) GetVote() (model.VoteData, error) {
	mu.Lock()
	defer mu.Unlock()

	return vote_data, nil
}
