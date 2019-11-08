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
	var department interface{} = nil
	//等级
	rank := int(info["rank"].(uint8))
	rank += 1

	//部门查看普通用户加部门限制
	if rank == 3 {
		department = info["department"].(string)
	}

	rows, err := db.SelectUserList(company, department, rank)
	if err != nil {
		return nil, err
	}
	return rows, err
}

//获取用户信息
func UserInfo(uid string) (interface{}, error) {
	//用户信息
	row, err := db.SelectUserById(uid)
	if err != nil {
		return nil, err
	}
	return row, err
}

//添加公司用户
func UserAddCompanyUser(p *user.AddCompanyUser, uid string) error {
	//检查公司存在
	b, err := db.HasCompany(p.Company)
	if err != nil {
		return err
	}
	if b {
		//公司存在
		return fmt.Errorf(EMPTY_STR)
	}
	//新建公司
	companyId := uuid.NewV4().String()
	err = db.InsertCompany(companyId, p.Company)
	if err != nil {
		return err
	}

	//新建公司用户
	userId := uuid.NewV4().String()
	err = userAdd(p.Email, uid, userId, companyId, nil)
	if err != nil {
		return err
	}

	//新建公司文件
	err = db.InsertFileFolder(companyId, p.Company, "departmentBucket", userId, "folder-company")
	if err != nil {
		return err
	}

	return nil
}

//添加部门用户
func UserAddDepartmentUser(p *user.AddDepartmentUser, uid string) error {
	//用户信息
	info, err := db.SelectUserById(uid)
	if err != nil {
		return err
	}
	//公司
	company := info["company"].(string)

	//检查部门存在
	b, err := db.HasDepartment(p.Department, company)
	if err != nil {
		return err
	}
	if b {
		//部门存在
		return fmt.Errorf(EMPTY_STR)
	}

	//新建部门
	departmentId := uuid.NewV4().String()
	err = db.InsertDepartment(departmentId, p.Department, company)
	if err != nil {
		return err
	}

	//新建用户
	userId := uuid.NewV4().String()
	err = userAdd(p.Email, uid, userId, nil, departmentId)
	if err != nil {
		return err
	}

	//新建部门文件
	err = db.InsertFileFolder(departmentId, p.Department, company, userId, "folder-department")
	if err != nil {
		return err
	}

	//建立一个同名的用户文件夹
	err = db.InsertFileFolder(userId, p.Email, departmentId, userId, "folder-user")
	if err != nil {
		return err
	}

	return nil
}

//添加普通用户
func UserAdd(p *user.Add, uid string) error {
	//用户信息
	info, err := db.SelectUserById(uid)
	if err != nil {
		return err
	}
	department := info["department"].(string)

	//新建用户
	userId := uuid.NewV4().String()
	err = userAdd(p.Email, uid, userId, nil, nil)
	if err != nil {
		return err
	}

	//新建用户文件
	err = db.InsertFileFolder(userId, p.Email, department, userId, "folder-user")
	if err != nil {
		return err
	}

	return nil
}

//添加管理员
func UserAddSuperAdmin() error {
	email := "Super Admin"
	id := "0-0-0-0-0"
	password := "Zoolon872112"
	company := "0-0-0-0-0"
	rank := 0

	hash, err := bcrypt.GenerateFromPassword([]byte(password+id), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = db.InsertUser(id, email, string(hash), rank, company, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

//删除用户
func UserDelete(p *user.Delete) error {
	err := db.UserDelete(p.Uid)
	if err != nil {
		return err
	}
	return nil
}

//添加用户
func userAdd(email, uid, userId string, company, department interface{}) error {
	//检查邮箱是否存在
	has, err := db.HasEmail(email)
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
	//添加者的公司，即被添加用户的所属上级公司

	//等级
	rank := int(info["rank"].(uint8))
	rank += 1
	//密码
	password := "123456"
	hash, err := bcrypt.GenerateFromPassword([]byte(password+userId), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	var belongId interface{} = nil
	var companyId interface{} = nil
	var departmentId interface{} = nil

	if rank == 1 {
		belongId = info["company"].(string)
		companyId = company
	} else if rank == 2 { //公司部门用户
		belongId = info["company"].(string)
		companyId = info["company"].(string)
		departmentId = department
	} else if rank == 3 { //普通用户
		belongId = info["company"].(string)
		companyId = info["company"].(string)
		departmentId = info["department"].(string)
	}
	err = db.InsertUser(userId, email, string(hash), rank, belongId, companyId, departmentId)
	if err != nil {
		return err
	}
	return nil
}
