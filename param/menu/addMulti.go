package menu

import (
	. "IxDServer/common"
	"IxDServer/network"
	"net/http"
)

type AddMultiMenu struct {
	Menu string `json:"menu" required:"true"`
	Sort int    `json:"sort" required:"true"`
}

//注册参数
type AddMulti struct {
	Users []string       `json:"users" required:"true"` //数组
	Menus []AddMultiMenu `json:"menus" required:"true"` //数组[{menu,sort},{menu,sort}]
}

func (p *AddMulti) Format(w http.ResponseWriter, r *http.Request) error {
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
