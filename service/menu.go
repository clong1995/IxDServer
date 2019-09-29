package service

import (
	"IxDServer/db"
	"IxDServer/param/menu"
	"fmt"
	"github.com/satori/go.uuid"
)

//获取菜单列表
func MenuGetList(user string) (interface{}, error) {
	return nil, fmt.Errorf("站位")
	//全部菜單
	/*rows, err := db.SelectMenuAllList()
	if err != nil {
		return nil, err
	}
	//用戶的菜單
	if user != "" {
		rowsUser, err := db.SelectMenuList(user)
		if err != nil {
			return nil, err
		}
		//所有的
		for _, v := range rows {
			//用戶有的
			for _, vu := range rowsUser {
				if vu["id"] == v["id"] {
					v["is"] = true
					break
				}
			}
		}
	}
	return rows, err*/
}

//获取菜单列表
func MenuGetListByUserStatus(user string) (interface{}, error) {
	return nil, fmt.Errorf("站位")
	/*
		//全部菜單
		rows, err := db.SelectMenuAllList()
		if err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		result["status"] = 0

		if user != "" {
			//用户的状态
			row, err := db.SelectUserById(user)
			if err != nil {
				return nil, err
			}

			//用戶的菜單
			rowsUser, err := db.SelectMenuList(user)
			if err != nil {
				return nil, err
			}

			//所有的
			for _, v := range rows {
				//用戶有的
				v["is"] = false
				for _, vu := range rowsUser {
					if vu["id"] == v["id"] {
						v["is"] = true
						break
					}
				}
			}

			result["status"] = row["status"]
		}

		result["menu"] = rows

		return result, err*/
}

//添加公司
func MenuAdd(p *menu.Add) error {
	id := uuid.NewV4().String()
	err := db.InsertMenu(id, p.User, p.Menu, p.Sort)
	if err != nil {
		return err
	}
	return nil
}

//添加公司
func MenuAddMulti(p *menu.AddMulti) error {
	users := p.Users
	menus := p.Menus
	err := db.InsertMenuMulti(users, menus)
	if err != nil {
		return err
	}
	return nil
}
