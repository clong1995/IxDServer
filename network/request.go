package network

import (
	. "IxDServer/common"
	"IxDServer/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

//json字符串请求转结构体
func GetReqJson(r *http.Request, i interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(PARAM_STR)
	} else {
		err := json.Unmarshal(data, i)
		if err != nil {
			log.Println(err)
			return fmt.Errorf(PARAM_STR)
		}
	}
	return nil
}

//TODO 非空检测，后期支持多层级校验
func CheckEmptyReqParam(itf interface{}) error {
	t := reflect.TypeOf(itf).Elem()
	v := reflect.ValueOf(itf).Elem()
	var f reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		f = t.Field(i)
		//必填项校验
		if f.Tag.Get("required") == "true" {
			//string类型
			if f.Type.Name() == "string" && v.Field(i).String() == "" {
				errStr := fmt.Sprintf("%s 的参数列表中，%s 不得为空", t.Name(), f.Name)
				log.Println(errStr)
				return fmt.Errorf(errStr)
			}
			//TODO 其他类型
		}
		//TODO 其他校验
	}
	return nil
}

//认证，权限
func Auth(w http.ResponseWriter, r *http.Request) (string, error) {
	//获取token
	auth := r.Header.Get("Authorization")
	if len(auth) < 100 {
		ErrStrCode(w, AUTH_STR, AUTH)
		return "", fmt.Errorf(AUTH_STR)
	}
	//解码token
	uId, err := util.UnMarshalToken(auth)
	if err != nil {
		ErrStrCode(w, EMPTY_STR, EMPTY)
		return "", err
	}
	return uId, nil
}

//跨域
func Origin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Accept,Authorization")
}
