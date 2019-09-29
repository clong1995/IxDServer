package service

import (
	"IxDServer/db"
	"IxDServer/param/template"
	uuid "github.com/satori/go.uuid"
)

func TemplateGetListByUser(uid string) (interface{}, error) {
	records, err := db.SelectTemplateListByUser(uid)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func TemplateAdd(uid string, p *template.Add) error {
	//添加
	id := uuid.NewV4().String()
	err := db.InsertTemplate(id, uid, p.Name, p.Data)
	if err != nil {
		return err
	}
	return nil
}

func TemplateGetById(p *template.GetById) (interface{}, error) {
	records, err := db.SelectTemplateById(p.Id)
	if err != nil {
		return nil, err
	}
	return records, nil
}
