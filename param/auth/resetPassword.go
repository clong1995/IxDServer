package auth

import (
	. "IxDServer/common"
	"IxDServer/network"
	"net/http"
)

//参数
type ResetPassword struct {
	OldPassword string `json:"oldPassword" required:"true"`
	NewPassword string `json:"newPassword" required:"true"`
}

func (p *ResetPassword) Format(w http.ResponseWriter, r *http.Request) error {
	//json转结构体
	err := network.GetReqJson(r, p)
	if err != nil {
		network.ErrStrCode(w, err.Error(), PARAM)
		return err
	}

	//非空检验
	err = network.CheckEmptyReqParam(p)
	if err != nil {
		network.ErrStrCode(w, err.Error(), PARAM)
		return err
	}

	//其他业务校验

	return nil
}
