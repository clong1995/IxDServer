package router

import (
	"IxDServer/handler"
	"net/http"
)

func init() {
	//登录
	http.HandleFunc("/auth/signin", handler.AuthSignin)
	//注册
	//http.HandleFunc("/auth/signup", handler.AuthSignup)
	//注册和登录
	//http.HandleFunc("/auth/signinAndSignup", handler.AuthSigninAndSignup)
	//TODO 解密token，这个接口和危险，用来测试的，上线后要删掉
	http.HandleFunc("/auth/unToken", handler.AuthUnToken)
	http.HandleFunc("/auth/getUid", handler.AuthGetUid)
	//续期
	http.HandleFunc("/auth/prolong", handler.AuthProlongToken)
	//重置密码
	http.HandleFunc("/auth/resetPassword", handler.AuthResetPassword)
}
