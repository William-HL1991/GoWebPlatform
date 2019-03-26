package handlers

import (
	"../model"
	"../utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//var appSession *sessions.Manager
//
//// 初始化session manager
//func init() {
//	appSession, _ = sessions.NewSessionManager("memory", "sessionid", 3600)
//	go appSession.GC()
//}


// cookies

func setCookie(w http.ResponseWriter, name string , vlaue string) {
	now := time.Now()
	d, _ := time.ParseDuration("1H")
	newNow := now.Add(d)
	fmt.Println(newNow)
	cookie := http.Cookie{Name: name, Value: vlaue, Expires:newNow}
	http.SetCookie(w, &cookie)
	fmt.Println("cookie 过期时间", cookie.Expires)
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		fmt.Println("cookie:", cookie.Name)
		fmt.Println("cookie", cookie.Value)
		fmt.Println("cookie", cookie.Expires)
	}

}

func Index(w http.ResponseWriter, r *http.Request) {
	readCookie(w, r)
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp := utils.NewBaseJsonBean()
		resp.Code = utils.REQUESTMETHODERROR
		resp.Message = "不支持get请求"
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp := utils.NewBaseJsonBean()
		resp.Code = utils.REQUESTMETHODERROR
		resp.Message = "只支持POST请求"
		respJson, err := json.Marshal(*resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(respJson)
	} else if r.Method == "POST" {
		username := r.Form.Get("username")
		passwd := r.Form.Get("password")
		// todo 简单的数据格式判断
		has, err := model.GetUserByName(username)
		fmt.Println(has, err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if has {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			resp := utils.NewBaseJsonBean()
			resp.Code = utils.USERNAMEUSED
			resp.Message = "该用户名已经被注册"
			respJson, err := json.Marshal(*resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(respJson)
		} else {
			err := model.NewUser(username, passwd)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// cookies
			setCookie(w, username, "logined")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			resp := utils.NewBaseJsonBean()
			resp.Code = 0
			resp.Message = "创建用户成功"
			respJson, err := json.Marshal(*resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(respJson)
		}
	}
	
}