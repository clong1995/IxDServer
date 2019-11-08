package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
)

//全部的部门
func SelectAllDepartment() ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
		select 
			id,
			name
		from department`)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()
	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

//公司下所有的部门
func SelectDepartmentByCompany(company string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
		select 
			id,
			name
		from department
		where company = ?`, company)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()
	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

//是否存在email
func HasDepartment(name, company string) (bool, error) {
	rows, err := Db.Query(`
		select id 
		from department 
		where name=? AND company = ? limit 1`, name, company)
	if err != nil {
		log.Println(err)
		return true, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()
	records, err := ParseRows(rows)
	if err != nil {
		return true, err
	}
	if len(records) > 0 {
		return true, nil
	}
	return false, nil
}

//插入用户
func InsertDepartment(id, name, company string) error {
	_, err := Db.Exec(
		"INSERT department (id,name,company) values (?,?,?)",
		id, name, company)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}
