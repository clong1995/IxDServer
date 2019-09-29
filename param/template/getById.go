package template

import (
	. "IxDServer/common"
	"IxDServer/network"
	"net/http"
)

//注册参数
type GetById struct {
	Id string `json:"id" required:"true"`
}

func (p *GetById) Format(w http.ResponseWriter, r *http.Request) error {

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
