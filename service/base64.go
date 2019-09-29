package service

import (
	"IxDServer/db"
	"IxDServer/param/base64"
	uuid "github.com/satori/go.uuid"
)

func Base64Add(p *base64.Add) (string, error) {
	var err error
	var id string
	if p.Id != "" {
		id = p.Id
		//更新
		err = db.UpdateBase64(id, p.Value)
	} else {
		//插入
		//添加
		id = uuid.NewV4().String()
		err = db.InsertBase64(id, p.Value)
	}

	if err != nil {
		return "", err
	}
	return id, nil
}

func Base64Get(p *base64.Get) (interface{}, error) {
	records, err := db.SelectBase64(p.Id)
	if err != nil {
		return nil, err
	}
	return records, nil
}
