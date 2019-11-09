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
	folderType := "folder"
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
	err = db.InsertFileFolder(id, p.Name, p.Pid, user, folderType)
	if err != nil {
		return err
	}
	return nil
}

func FileAddFile(p *file.AddFile, user string) (string, error) {
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
	case ext == ".mp3", ext == ".flac", ext == ".wav", ext == ".apc":
		typee = "audio"
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
	case ext == ".pdf":
		typee = "pdf"
	case ext == ".sql":
		typee = "database"
	case ext == ".dmg":
		typee = "dmg"
	case ext == ".pkg":
		typee = "pkg"
	case ext == ".apk":
		typee = "apk"
	case ext == ".jar":
		typee = "jar"
	default:
		typee = "file"
	}
	id := uuid.NewV4().String()
	err := db.InsertFile(id, p.Etag, p.Name, p.Pid, typee, user, p.MimeType, p.Local, p.Size, p.State)
	if err != nil {
		return "", err
	}
	return id, nil
}

func FileDelete(p *file.Delete, user string) error {
	//检查是否属于自己
	flag, err := db.SelectFileOwn(p.Id, user)
	if err != nil {
		return err
	}
	if !flag {
		return fmt.Errorf("无权删除别人的文件")
	}

	//TODO 等级是部门管理员，则可删除成员的

	err = db.UpdateFileState(p.Id, 1)
	if err != nil {
		return err
	}
	return nil
}
func FileRename(p *file.Rename, user string) error {
	//检查是否属于自己
	flag, err := db.SelectFileOwn(p.Id, user)
	if err != nil {
		return err
	}
	if !flag {
		return fmt.Errorf("无权修改别人的文件")
	}
	err = db.UpdateFileName(p.Id, p.Name)
	if err != nil {
		return err
	}
	return nil
}

func FileCut(p *file.Cut, user string) error {
	//检查是否属于自己
	flag, err := db.SelectFileOwn(p.File, user)
	if err != nil {
		return err
	}
	if !flag {
		return fmt.Errorf("无权剪切别人的文件")
	}
	err = db.UpdateFilePid(p.File, p.Dist)
	if err != nil {
		return err
	}
	return nil
}

func FileCopy(p *file.Copy, user string) error {
	//复制
	err := copyFile(p.File, p.Dist, user)
	if err != nil {
		return err
	}
	//自动重命名
	return nil
}

func FileRemove(p *file.Remove) error {
	/*//删除数据
	err := db.DeleteFile(p.Id)
	if err != nil {
		return err
	}
	return nil*/
	err := db.UpdateFileState(p.Id, 3)
	if err != nil {
		return err
	}
	return nil
}

func FileUploadFinish(p *file.UploadFinish) error {
	//建立新的文件夹
	err := db.UpdateFileFinish(p.Id)
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
	rows, err := db.SelectFileList(p.Pid)
	if err != nil {
		return nil, err
	}
	for _, v := range rows {
		if v["user"] == user {
			v["user"] = "own"
		} else {
			v["user"] = ""
		}
	}
	return rows, err
}

func FileMyList(p *file.List, user string) (interface{}, error) {
	rows, err := db.SelectFileMyList(p.Pid, user)
	if err != nil {
		return nil, err
	}
	for _, v := range rows {
		if v["user"] == user {
			v["user"] = "own"
		} else {
			v["user"] = ""
		}
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

func FileInfo(p *file.Info) (interface{}, error) {
	row, err := db.SelectFileInfoById(p.Id)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func FileCheckFinish(p *file.CheckFinish, user string) (interface{}, error) {
	if len(p.Ids) > 0 {
		rows, err := db.SelectFileCheckFinish(p.Ids, user)
		if err != nil {
			return nil, err
		}
		return rows, nil
	}
	return nil, nil
}

func FileUploading(p *file.Uploading, user string) (interface{}, error) {
	if len(p.Etags) > 0 {
		rows, err := db.SelectFileUploading(p.Etags, user)
		if err != nil {
			return nil, err
		}
		for _, v := range rows {
			if v["user"] == user {
				v["user"] = "own"
			} else {
				v["user"] = ""
			}
		}
		return rows, nil
	}
	return nil, nil
}

func FileListByEtags(p *file.ListByEtags, user string) (interface{}, error) {
	if len(p.Etags) > 0 {
		rows, err := db.SelectFileByEtags(p.Etags, user)
		if err != nil {
			return nil, err
		}
		return rows, nil
	}
	return [0]int{}, nil
}

func FileDepartmentList(user string) (interface{}, error) {
	//查询当前用户的等级
	userInfo, err := db.SelectUserById(user)
	if err != nil {
		return nil, err
	}
	rank := userInfo["rank"].(uint8)
	department := userInfo["department"].(string)
	company := userInfo["company"].(string)

	var rows []map[string]interface{}
	switch rank {
	case 0: //如果是Super Admin，能看到全部公司
		rows, err = db.SelectFileFolderCompany()
		break
	case 1: //如果是公司管理员 xxx Admin，看到本公司下所有的部门文件
		rows, err = db.SelectFileByPidList(company)
		break
	case 2: //如果是部门用户，看到自己部门下的员工的文件
		rows, err = db.SelectFileByPidList(department)
		break
	case 3: //看到自己
		rows, err = db.SelectFileByUserList(user)
		break
	}

	if err != nil {
		return nil, err
	}
	for _, v := range rows {
		if v["user"] == user {
			v["user"] = "own"
		} else {
			delete(v, "user")
		}
	}
	return rows, nil
}

func FileDepartmentPublic(user string) (interface{}, error) {
	userInfo, err := db.SelectUserById(user)
	if err != nil {
		return nil, err
	}
	department := userInfo["department"].(string)
	//查询这个部门下所有的顶级部门文件
	rows, err := db.SelectFileDepartmentBucket(department)
	if err != nil {
		return nil, err
	}
	for _, v := range rows {
		if v["user"] == user {
			v["user"] = "own"
		} else {
			delete(v, "user")
		}
	}
	return rows, nil
}

//TODO 这里高并发，阻塞，单线程。需要优化
func copyFile(fId, pid, user string) error {
	//首先复制自己
	row, err := db.SelectFileById(fId)
	id := uuid.NewV4().String()
	err = db.InsertFile(id, row["etag"].(string), row["name"].(string), pid, row["type"].(string), user, row["mime"].(string), "", row["size"], 0)
	if err != nil {
		return err
	}

	//向前移动
	pid = id

	//查询子目录文件
	rows, err := db.SelectFileList(fId)
	if err != nil {
		return err
	}
	//复制子目录
	if len(rows) > 0 {
		for _, v := range rows {
			err := copyFile(v["id"].(string), pid, user)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
