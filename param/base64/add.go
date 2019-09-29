package base64

import (
	. "IxDServer/common"
	"IxDServer/network"
	"net/http"
)

//注册参数
type Add struct {
	Id    string `json:"id"`                    //id
	Value string `json:"value" required:"true"` //数据
}

func (p *Add) Format(w http.ResponseWriter, r *http.Request) error {
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
