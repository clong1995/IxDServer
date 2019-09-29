package service

import (
	. "IxDServer/common"
	"IxDServer/db"
	"IxDServer/param/scene"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func SceneGetListByProject(p *scene.GetListByProject) (interface{}, error) {
	records, err := db.SelectSceneListByProject(p.Project)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SceneAdd(p *scene.Add) (string, error) {
	if p.Scene != "" {
		//更新
		err := db.UpdateScene(p.Scene, p.Name, p.Data)
		if err != nil {
			return "", err
		}
		return p.Scene, nil
	} else {
		//添加
		if p.Project == "" {
			return "", fmt.Errorf(PARAM_STR)
		}
		id := uuid.NewV4().String()
		err := db.InsertScene(id, p.Project, p.Name, p.Data)
		if err != nil {
			return "", err
		}
		return id, nil
	}
}

func SceneGetById(p *scene.GetById) (interface{}, error) {
	records, err := db.SelectSceneById(p.Id)
	if err != nil {
		return nil, err
	}
	return records, nil
}
