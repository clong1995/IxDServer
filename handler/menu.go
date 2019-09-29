package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/menu"
	"IxDServer/service"
	"net/http"
)

//TODO 获取所有菜单的列表，後面可能根權限有關
func MenuGetList(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		//读取head
		token := r.Header.Get("Authorization")
		user, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//业务
		rows, err := service.MenuGetList(user)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}

//获取所有菜单的列表
func MenuGetListByUser(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		p := new(menu.GetListByUser)
		err := p.Format(w, r)
		if err != nil {
			return
		}
		//读取head
		token := r.Header.Get("Authorization")
		_, err = service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//业务
		rows, err := service.MenuGetList(p.User)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}

//获取所有菜单带用户状态的列表
func MenuGetListByUserStatus(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		p := new(menu.GetListByUser)
		err := p.Format(w, r)
		if err != nil {
			return
		}
		//读取head
		token := r.Header.Get("Authorization")
		_, err = service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//业务
		rows, err := service.MenuGetListByUserStatus(p.User)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}

func MenuAdd(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		//读取head
		token := r.Header.Get("Authorization")
		_, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(menu.Add)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.MenuAdd(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

//批量为多个用户设置菜单
func MenuAddMulti(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		//读取head
		token := r.Header.Get("Authorization")
		_, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(menu.AddMulti)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//校验
		//用户
		if len(p.Users) < 1 {
			network.ErrStrCode(w, PARAM_STR, PARAM)
			return
		}
		//菜单
		if len(p.Menus) < 1 {
			network.ErrStrCode(w, PARAM_STR, PARAM)
			return
		}

		//业务
		err = service.MenuAddMulti(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}
