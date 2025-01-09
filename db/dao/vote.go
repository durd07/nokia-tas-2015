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

	// construct VoteData
	vote_data = make(model.VoteData)

	for k, v := range data {
		if _, exists := vote_data[k]; !exists {
			vote_data[k] = &model.VoteModel{
				Name: k,
				Members: make(map[string]*model.MemberModel),
			}
		}

		for _, x := range v {
			vote_data[k].Members[x] = &model.MemberModel{
				Name: x,
				Vote: 0,
			}
		}
	}
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
func (imp *VoteInterfaceImp) UpsertVote(data model.VoteData) error {
	mu.Lock()
	defer mu.Unlock()

	for k, manager := range data {
		for v, _ := range manager.Members {
			vote_data[k].Members[v].Vote += 1
		}
	}

	return nil
}

// GetVote 查询Vote
func (imp *VoteInterfaceImp) GetVote() (model.VoteData, error) {
	mu.Lock()
	defer mu.Unlock()

	return vote_data, nil
}
