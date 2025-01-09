package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodGet {
		data, err := getCurrentVote()
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
	body := make(model.VoteData)
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

func getCurrentVote() (model.VoteData, error) {
	data, err := dao.VoteImp.GetVote()
	if err != nil {
		return nil, err
	}

	return data, nil
}
