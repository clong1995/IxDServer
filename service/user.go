package service

import (
	. "IxDServer/common"
	"IxDServer/db"
	"IxDServer/param/user"
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//获取菜单列表
func UserList(uid string) (interface{}, error) {
	//用户信息
	info, err := db.SelectUserById(uid)
	if err != nil {
		return nil, err
	}
	//公司
	company := info["company"].(string)
	//等级
	rank := int(info["rank"].(uint32))
	rank += 1

	rows, err := db.SelectUserList(company, rank)
	if err != nil {
		return nil, err
	}
	return rows, err
}

//添加用户
func UserAdd(p *user.Add, uid string) error {
	//检查邮箱是否存在
	has, err := db.HasEmail(p.Email)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf(MULTIPLE_STR)
	}
	info, err := db.SelectUserById(uid)
	if err != nil {
		return err
	}
	//用户id
	id := uuid.NewV4().String()
	//添加者的公司，即被添加用户的所属上级公司
	company := info["company"].(string)
	//等级
	rank := int(info["rank"].(uint32))
	rank += 1
	//密码
	password := "123456"
	hash, err := bcrypt.GenerateFromPassword([]byte(password+id), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = db.InsertUser(id, p.Email, string(hash), company, rank)
	if err != nil {
		return err
	}
	return nil
}

//添加公司
func UserDelete(p *user.Delete) error {
	err := db.UserDelete(p.Uid)
	if err != nil {
		return err
	}
	return nil
}
