package db

import (
	. "IxDServer/common"
	"IxDServer/param/menu"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
)

//根据用户查询菜单
func SelectMenuListByUser(user string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
SELECT
	m.id AS id,
	m.name AS name,
	m.type AS type 
FROM
	menu m
	LEFT JOIN user__menu um ON um.menu = m.id 
WHERE
	um.user = ?`, user)

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

//查詢所有菜單，用於全線部分
func SelectMenuAllList() ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
SELECT
	m.id,
	m.name,
	m.type 
FROM
	menu m`)

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

func InsertMenu(id, user, menu string, sort int) error {
	_, err := Db.Exec("INSERT INTO user__menu (id,user,menu,sort) values (?,?,?,?)", id, user, menu, sort)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

//批量添加用户
func InsertMenuMulti(user []string, menus []menu.AddMultiMenu) error {
	//开启事务
	tx, _ := Db.Begin()
	defer tx.Rollback()

	//删除当前用户的
	args := make([]interface{}, 0)
	in := ""
	for _, u := range user {
		args = append(args, u)
		in += "?,"
	}
	in = strings.TrimRight(in, ",")
	sql := "DELETE FROM user__menu WHERE user in (" + in + ")"
	_, err := Db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}

	//添加新的
	args = make([]interface{}, 0)
	sql = "INSERT INTO user__menu (id,user,menu,sort) values"
	for _, u := range user {
		for _, m := range menus {
			args = append(args, uuid.NewV4().String(), u, m.Menu, m.Sort)
			sql += "(?,?,?,?),"
		}
	}
	//去掉最后的逗号
	sql = strings.TrimRight(sql, ",")
	_, err = Db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}

	//提交事务
	err = tx.Commit()
	return nil
}
