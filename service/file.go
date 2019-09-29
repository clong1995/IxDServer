package service

import (
	. "IxDServer/common"
	"IxDServer/db"
	"IxDServer/param/file"
	"fmt"
	"github.com/satori/go.uuid"
	"path"
)

func FileAddFolder(p *file.AddFolder, user string) error {
	//检查是否重名
	b, err := db.HasFileFolderName(p.Name, p.Pid, user)
	if err != nil {
		return nil
	}
	if b {
		return fmt.Errorf(MULTIPLE_STR)
	}
	//建立新的文件夹
	id := uuid.NewV4().String()
	err = db.InsertFileFolder(id, p.Name, p.Pid, user)
	if err != nil {
		return err
	}
	return nil
}

func FileAddFile(p *file.AddFile, user string) error {
	//查询文件是否存在
	//文件类型
	typee := "file"
	ext := path.Ext(p.Name) //获取文件后缀
	switch {
	case ext == ".png", ext == ".jpg", ext == ".jpeg", ext == ".gif", ext == ".bmp":
		typee = "image"
	case ext == ".mp4", ext == "wmv", ext == "asf", ext == "rm", ext == "rmvb", ext == "mov", ext == "avi",
		ext == "mpg", ext == "mpeg", ext == "mpeg1", ext == "mpej2", ext == "mpej3", ext == "mpej4",
		ext == "vob", ext == "dat", ext == "divx":
		typee = "video"
	case ext == ".js":
		typee = "js"
	case ext == ".css":
		typee = "css"
	case ext == ".zip", ext == ".rar", ext == ".7z", ext == ".tar", ext == ".zg":
		typee = "zip"
	case ext == ".html":
		typee = "html"
	case ext == ".psd":
		typee = "psd"
	case ext == ".xls", ext == ".xlsx":
		typee = "excel"
	case ext == ".json":
		typee = "json"
	case ext == ".doc", ext == ".docx":
		typee = "word"
	case ext == ".ppt", ext == ".pptx":
		typee = "ppt"
	case ext == ".exe":
		typee = "exe"
	case ext == ".txt":
		typee = "txt"
	default:
		typee = "file"
	}
	id := uuid.NewV4().String()
	err := db.InsertFile(id, p.Etag, p.Name, p.Pid, typee, user, p.MimeType, p.Size, p.State)
	if err != nil {
		return err
	}
	return nil
}

func FileDelete(p *file.Delete) error {
	//建立新的文件夹
	err := db.UpdateFileState(p.Id, 1)
	if err != nil {
		return err
	}
	return nil
}

func FileUploadFinish(p *file.UploadFinish) error {
	//建立新的文件夹
	err := db.UpdateFileState(p.Id, 0)
	if err != nil {
		return err
	}
	return nil
}

func FileListTopFolder() (interface{}, error) {
	rows, err := db.SelectFileTopFolderList()
	if err != nil {
		return nil, err
	}
	return rows, err
}

func FileList(p *file.List, user string) (interface{}, error) {
	rows, err := db.SelectFileList(p.Pid, user)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func FileDeleteList(user string) (interface{}, error) {
	rows, err := db.SelectFileDeleteList(user)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func FileTaskList(user string) (interface{}, error) {
	rows, err := db.SelectFileTaskList(user)
	if err != nil {
		return nil, err
	}
	return rows, err
}
