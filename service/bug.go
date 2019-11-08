package service

import (
	"IxDServer/db"
	"IxDServer/param/bug"
	uuid "github.com/satori/go.uuid"
)

//获取菜单列表
func BugList(uid string) (interface{}, error) {
	//用户信息
	info, err := db.SelectUserById(uid)
	if err != nil {
		return nil, err
	}
	//公司
	company := info["company"].(string)
	//等级
	rank := int(info["rank"].(uint8))
	rank += 1

	rows, err := db.SelectBugList(company, rank)
	if err != nil {
		return nil, err
	}
	return rows, err
}

//获取菜单列表
func BugInfo(p *bug.Info) (interface{}, error) {
	//用户信息
	row, err := db.SelectBugById(p.Id)
	if err != nil {
		return nil, err
	}
	return row, err
}

//添加bug
func BugAdd(p *bug.Add, uid string) error {
	id := uuid.NewV4().String()
	err := db.InsertBug(uid, id, p.Title, p.Position, p.Reappear, p.Expect, p.Type, p.Severity, p.Priority)
	if err != nil {
		return err
	}
	return nil
}
