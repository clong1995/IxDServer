package service

import (
	"IxDServer/db"
)

func DepartmentList(uid string) (interface{}, error) {
	//用户信息
	info, err := db.SelectUserById(uid)
	if err != nil {
		return nil, err
	}
	//公司
	company := info["company"].(string)

	rows, err := db.SelectDepartmentByCompany(company)
	if err != nil {
		return nil, err
	}
	return rows, err
}
