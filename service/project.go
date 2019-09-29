package service

import (
	"IxDServer/db"
	"IxDServer/param/project"
	"github.com/satori/go.uuid"
)

//添加公司
func ProjectAdd(uid string, p *project.Add) error {
	id := uuid.NewV4().String()
	err := db.InsertProject(id, uid, p.Name)
	if err != nil {
		return err
	}
	return nil
}

func ProjectList(uid string) (interface{}, error) {
	records, err := db.SelectProject(uid)
	if err != nil {
		return nil, err
	}
	return records, nil
}

//删除设备
func ProjectDelete(p *project.Delete) error {
	err := db.DeleteProject(p.Project)
	if err != nil {
		return err
	}
	return nil
}
