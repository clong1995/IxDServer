package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/auth"
	"IxDServer/service"
	"net/http"
)

//注册
/*
func AuthSignup(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)

	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		p := new(auth.SigninAndSignUp)
		err := p.Format(w, r)
		if err != nil {
			return
		}
		//业务
		token, err := service.AuthSignup(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, token)
	}
}*/

//登录
func AuthSignin(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)

	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		p := new(auth.Signin)
		err := p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		rows, err := service.AuthSignin(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}

//注册和登录
/*func AuthSigninAndSignup(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)

	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		p := new(auth.SigninAndSignUp)
		err := p.Format(w, r)
		if err != nil {
			return
		}
		//业务
		token, err := service.AuthSigninAndSignup(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, token)
	}
}*/

//TODO 这个方法很危险，是测试用的。解析token获取用户id
func AuthUnToken(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//读取head
		token := r.Header.Get("Authorization")
		str, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}
		if str == "" {
			network.ErrStrCode(w, AUTH_STR, AUTH)
			return
		}
		network.Succ(w, str)
	}
}

//根据token获取id，给客户端用的
func AuthGetUid(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//读取head
		//参数
		p := new(auth.GetUid)
		err := p.Format(w, r)
		if err != nil {
			return
		}
		str, err := service.AuthUnToken(p.Token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}
		if str == "" {
			network.ErrStrCode(w, AUTH_STR, AUTH)
			return
		}
		network.Succ(w, str)
	}
}

//续期
func AuthProlongToken(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//读取head
		token := r.Header.Get("Authorization")
		str, err := service.AuthProlongToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}
		network.Succ(w, str)
	}
}
