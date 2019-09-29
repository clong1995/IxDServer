package menu

import (
	. "IxDServer/common"
	"IxDServer/network"
	"net/http"
)

//注册参数
type Add struct {
	User string `json:"user" required:"true"`
	Menu string `json:"menu" required:"true"`
	Sort int    `json:"sort" required:"true"`
}

func (c *Add) Format(w http.ResponseWriter, r *http.Request) error {
	//json转结构体
	err := network.GetReqJson(r, c)
	if err != nil {
		network.ErrStrCode(w, err.Error(), PARAM)
		return err
	}

	//非空检验
	err = network.CheckEmptyReqParam(c)
	if err != nil {
		network.ErrStrCode(w, err.Error(), PARAM)
		return err
	}

	//其他业务校验
	return nil
}
