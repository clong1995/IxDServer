package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/user"
	"IxDServer/service"
	"net/http"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		//读取head
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//业务
		rows, err := service.UserList(uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		//读取head
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//业务
		rows, err := service.UserInfo(uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}

func UserAdd(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		//读取head
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(user.Add)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.UserAdd(p, uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

func UserAddCompanyUser(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		//读取head
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(user.AddCompanyUser)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.UserAddCompanyUser(p, uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

func UserAddDepartmentUser(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		//读取head
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(user.AddDepartmentUser)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.UserAddDepartmentUser(p, uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

func UserAddSuperAdmin(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		//参数
		//业务
		err := service.UserAddSuperAdmin()
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
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
		p := new(user.Delete)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.UserDelete(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}
