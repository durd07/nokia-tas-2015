package service

import (
	"encoding/json"
	"fmt"
	"time"
	"net/http"
)

var (
	version string
)

func init() {
	version = time.Now().Format("2006-01-02 15:04:05")
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		res := map[string]string{"version": version}

		msg, err := json.Marshal(res)
		if err != nil {
			fmt.Fprint(w, "内部错误")
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(msg)
	} else if r.Method == http.MethodPost {
		version = time.Now().Format("2006-01-02 15:04:05")
		w.Write([]byte("OK"))
	}
}
