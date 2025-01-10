package service

import (
	"encoding/json"
	"fmt"
	"time"
	"log"
	"net/http"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

func VoteMembersHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method != http.MethodGet {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	} else {
		data, err := getMembers()
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Code = 0
			res.Data = data
		}
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func VoteDumpHandler(w http.ResponseWriter, r *http.Request) {
	data := getAllVoteResult()
	ret := map[string]interface{}{
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		"vote": data,
	}

	msg, err := json.Marshal(ret)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodGet {
		data, err := getVoteResult()
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = data
		}
	} else if r.Method == http.MethodPost {
		err := upsertVote(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = ""
		}
	} else if r.Method == http.MethodDelete {
		// update version
		version = time.Now().Format("2006-01-02 15:04:05")

		err := clearVote()
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = ""
		}
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func upsertVote(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var body []string
	if err := decoder.Decode(&body); err != nil {
		return err
	}
	defer r.Body.Close()

	if err := dao.VoteImp.UpsertVote(body); err != nil {
		return err
	}
	return nil
}

func clearVote() error {
	return dao.VoteImp.ClearVote()
}

func getVoteResult() ([]model.MemberModel, error) {
	defaults := map[string]string{
		"Jason Team": "Kevin Kou",
		"Laura Team": "Daniel Sun",
		"Michael Team": "Patrick Zhou",
		"Oliver Team": "Fu Tang Wang",
		"Teddy Team": "Tony Tian",
		"Katie Team": "Lesley Liu",
	}

	data, _ := dao.VoteImp.GetVote()

	text, _ := json.Marshal(data)
	log.Printf("Vote Result %v\n", string(text))

	ret := []model.MemberModel{}
	for g, v := range data.Data {
		maxkey := ""
		maxval := 0

		for k, x := range v.Members {
			if int(x.Vote) > maxval {
				maxkey = k
				maxval = int(x.Vote)
			}
		}

		if maxkey == "" {
			ret = append(ret, model.MemberModel{Name: defaults[g], Vote: 1})
		} else {
			ret = append(ret, *data.Data[g].Members[maxkey])
		}
	}
	return ret, nil
}

func getAllVoteResult() map[string]*model.VoteModel {
	data, _ := dao.VoteImp.GetVote()

	text, _ := json.Marshal(data)
	log.Printf("Vote Result %v\n", string(text))

	return data.Data
}

func getMembers() ([]string, error) {
	data, _ := dao.VoteImp.GetVote()

	ret := []string{}
	for _, v := range data.Data {
		for k, _ := range v.Members {
			ret = append(ret, k)
		}
	}
	return ret, nil
}
