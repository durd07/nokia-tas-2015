package dao

import (
	"os"
	"sync"
	"gopkg.in/yaml.v3"
	"log"
)

type VoteData map[string]map[string]int32

var (
	vote_data VoteData
	vote_data_list map[string][]string

	mu        sync.Mutex
	vote_data_list_lock sync.RWMutex
)

func loadData() {
	data, err := os.ReadFile("./ntas.yaml")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	vote_data_list_lock.Lock()
	err = yaml.Unmarshal(data, &vote_data_list)
	vote_data_list_lock.Unlock()
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	mu.Lock()
	defer mu.Unlock()

	vote_data_list_lock.RLock()
	defer vote_data_list_lock.RUnlock()

	// construct VoteData
	vote_data = make(VoteData)

	for k, v := range vote_data_list {
		if _, exists := vote_data[k]; !exists {
			vote_data[k] = make(map[string]int32)
		}

		for _, x := range v {
			vote_data[k][x] = 0
		}
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

func (imp *VoteInterfaceImp) GetVoteList() (*map[string][]string, error) {
	vote_data_list_lock.RLock()
	defer vote_data_list_lock.RUnlock()

	return &vote_data_list, nil
}
