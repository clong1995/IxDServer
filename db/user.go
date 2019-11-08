package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
)

//插入用户
func InsertUser(id, email, password string, rank int, belong, company, department interface{}) error {
	_, err := Db.Exec(
		"INSERT user (id,email,password,`rank`,belong,company,department) values (?,?,?,?,?,?,?)",
		id, email, password, rank, belong, company, department)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

//增加超级管理员
func InsertSuperAdmin(id, email, password string) error {
	_, err := Db.Exec(
		"INSERT user (id,email,password) values (?,?,?)",
		id, email, password)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

//插入用户
func UserDelete(uid string) error {
	_, err := Db.Exec("DELETE FROM user WHERE id = ?", uid)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

//是否存在email
func HasEmail(email string) (bool, error) {
	rows, err := Db.Query("select email from user where email=? limit 1", email)
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

//根据email查询用户信息
func SelectUserByEmail(email string) (map[string]interface{}, error) {
	rows, err := Db.Query(`
		select 
			id,
			email,
			password,
			status,
			create_time
		from user 
		where email=? limit 1`, email)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()
	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	if len(records) != 1 {
		return nil, fmt.Errorf(EMPTY_STR)
	}
	return records[0], nil
}

//根据id查询用户信息
func SelectUserById(id string) (map[string]interface{}, error) {
	rows, err := Db.Query(`
		select 
			id,
			email,
			status,
			rank,
		    department,
			company,
			belong,
			create_time
		from user
		where id= ? limit 1`, id)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()
	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	if len(records) != 1 {
		log.Println(EMPTY_STR)
		return nil, fmt.Errorf(EMPTY_STR)
	}
	return records[0], nil
}

//获取用户诶列表
func SelectUserList(company, department interface{}, rank int) ([]map[string]interface{}, error) {
	args := make([]interface{}, 0)
	where := "where u.belong = ? "
	args = append(args, company)
	if department != nil {
		args = append(args, department)
		where += "AND u.department = ? "
	}
	args = append(args, rank)
	where += "AND rank = ? "

	rows, err := Db.Query(
		fmt.Sprintf(`SELECT 
       				u.id,
       				u.email,
       				u.status,
       				c.name as company,
		   			d.name as department,
       				u.create_time 
				FROM user u
				LEFT JOIN company c on u.company = c.id
				left join department d on u.department = d.name
				%s
				ORDER BY u.update_time DESC`, where),
		args...)
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

//获取用户诶列表
func SelectUserByDepartment(department string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(
		`SELECT 
       				id,
       				email,
       				status,
       				create_time 
				FROM user u
				where department = ?
				ORDER BY u.update_time DESC`,
		department)
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

//自己部门的人员
func SelectUserBySelfDepartment(user string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
		select 
			d.id,
			d.name
		from department d
		left join user u on u.department = d.id
		where u.id = ?`, user)

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

//设置用户状态
func UpdateUserPassword(id, password string) error {
	_, err := Db.Exec(
		"UPDATE user SET password = ? WHERE id=?", password, id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}
