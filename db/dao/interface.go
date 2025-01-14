package dao

import (
	"wxcloudrun-golang/db/model"
)

// CounterInterface 计数器数据模型接口
type CounterInterface interface {
	GetCounter(id int32) (*model.CounterModel, error)
	UpsertCounter(counter *model.CounterModel) error
	ClearCounter(id int32) error
}

// CounterInterfaceImp 计数器数据模型实现
type CounterInterfaceImp struct{}

// Imp 实现实例
var Imp CounterInterface = &CounterInterfaceImp{}

// VoteInterface 计数器数据模型接口
type VoteInterface interface {
	GetVote() (model.VoteData, error)
	UpsertVote(data []string) error
	ClearVote() error
}

// CounterInterfaceImp 计数器数据模型实现
type VoteInterfaceImp struct{}

// Imp 实现实例
var VoteImp VoteInterface = &VoteInterfaceImp{}
