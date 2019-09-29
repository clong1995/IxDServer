package handler

import (
	. "IxDServer/common"
	"IxDServer/network"
	"IxDServer/param/file"
	"IxDServer/service"
	"net/http"
)

func FileAddFolder(w http.ResponseWriter, r *http.Request) {
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
		p := new(file.AddFolder)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.FileAddFolder(p, uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

func FileAddFile(w http.ResponseWriter, r *http.Request) {
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
		p := new(file.AddFile)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.FileAddFile(p, uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

func FileDelete(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		token := r.Header.Get("Authorization")
		_, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(file.Delete)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.FileDelete(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

func FileUploadFinish(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//认证
		token := r.Header.Get("Authorization")
		_, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(file.UploadFinish)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		err = service.FileUploadFinish(p)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}

		network.Succ(w, "")
	}
}

func FileListTopFolder(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//参数
		//读取head
		token := r.Header.Get("Authorization")
		_, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//业务
		rows, err := service.FileListTopFolder()
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}

func FileList(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//读取head
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		p := new(file.List)
		err = p.Format(w, r)
		if err != nil {
			return
		}

		//业务
		rows, err := service.FileList(p, uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}

func FileDeleteList(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//读取head
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}

		//参数
		//业务
		rows, err := service.FileDeleteList(uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}

func FileTaskList(w http.ResponseWriter, r *http.Request) {
	//跨域
	network.Origin(w)
	if r.Method == http.MethodGet {
		network.FbdReq(w)
	} else if r.Method == http.MethodPost {
		//读取head
		token := r.Header.Get("Authorization")
		uid, err := service.AuthUnToken(token)
		if err != nil {
			network.ErrStrCode(w, err.Error(), AUTH)
			return
		}
		//参数
		//业务
		rows, err := service.FileTaskList(uid)
		if err != nil {
			network.ErrStr(w, err.Error())
			return
		}
		network.Succ(w, rows)
	}
}
