package handlers

import (
	"encoding/json"
	"net/http"
	"../utils"
	)

func Index(w http.ResponseWriter, r *http.Request)  {
	resp := utils.NewBaseJsonBean()
	resp.Code = 0
	resp.Message = "链接正常"
	respJson, err := json.Marshal(*resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func Trigger(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		resp := utils.NewBaseJsonBean()
		resp.Code = utils.CLIENT_ERROR
		resp.Message = "不支持get请求"
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}

}
