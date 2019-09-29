package auth

import (
	. "IxDServer/common"
	"IxDServer/network"
	"net/http"
)

//注册参数
type Signup struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

func (a *Signup) Format(w http.ResponseWriter, r *http.Request) error {
	//json转结构体
	err := network.GetReqJson(r, a)
	if err != nil {
		network.ErrStrCode(w, err.Error(), PARAM)
		return err
	}

	//非空检验
	err = network.CheckEmptyReqParam(a)
	if err != nil {
		network.ErrStrCode(w, err.Error(), PARAM)
		return err
	}

	//其他业务校验
	return nil
}
