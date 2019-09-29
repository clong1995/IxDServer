package file

import (
	. "IxDServer/common"
	"IxDServer/network"
	"net/http"
)

//注册参数
type AddFile struct {
	Etag     string  `json:"etag" required:"true"`
	Name     string  `json:"name" required:"true"`
	Size     float64 `json:"size" required:"true"`
	MimeType string  `json:"mimeType"`
	Pid      string  `json:"pid" required:"true"`
	State    int     `json:"state" required:"true"`
}

func (c *AddFile) Format(w http.ResponseWriter, r *http.Request) error {
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
