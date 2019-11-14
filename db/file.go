package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
	"strings"
)

func InsertFileFolder(id, name, pid, user, folderType string) error {
	_, err := Db.Exec("INSERT INTO file (id, name,type, pid, user) values (?,?,?,?,?)", id, name, folderType, pid, user)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

func InsertFile(id, etag, name, pid, typee, user, mime, local string, size interface{}, state int) error {
	_, err := Db.Exec("INSERT INTO file (id,etag,name,pid,type,user,mime,local,size,state) values (?,?,?,?,?,?,?,?,?,?)",
		id, etag, name, pid, typee, user, mime, local, size, state)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

func UpdateFileState(id string, state int) error {
	result, err := Db.Exec("UPDATE file SET state = ? WHERE id = ?", state, id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	i, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	if i < 0 {
		return fmt.Errorf(EMPTY_STR)
	}
	return nil
}

func UpdateFileName(id, name string) error {
	result, err := Db.Exec("UPDATE file SET name = ? WHERE id = ?", name, id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	i, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	if i < 0 {
		return fmt.Errorf(EMPTY_STR)
	}
	return nil
}

func DeleteFile(id string) error {
	result, err := Db.Exec("DELETE FROM file WHERE id = ? AND state = 1", id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	i, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	if i < 0 {
		return fmt.Errorf(EMPTY_STR)
	}
	return nil
}

func UpdateFileFinish(id string) error {
	result, err := Db.Exec("UPDATE file SET state = 0,local = '' WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	i, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	if i < 0 {
		return fmt.Errorf(EMPTY_STR)
	}
	return nil
}

func UpdateFilePid(file string, pid string) error {
	result, err := Db.Exec("UPDATE file SET pid = ? WHERE id = ?", pid, file)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	i, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	if i < 0 {
		return fmt.Errorf(EMPTY_STR)
	}
	return nil
}

func SelectFileTopFolderList() ([]map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT id,name,type FROM file WHERE pid IS NULL AND state=0 ORDER BY sort ASC`)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileList(pid string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
		SELECT 
			id, etag, name, mime, type, state, user,size, update_time 
		FROM 
			file 
		WHERE 
			pid=? 
		AND 
			state!=1 
		ORDER BY 
			update_time 
		DESC
		`, pid)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileOwn(fId string, uId string) (bool, error) {
	rows, err := Db.Query(`
		SELECT 
			id
		FROM 
			file 
		WHERE 
			id=? 
		AND user = ?
		`, fId, uId)
	if err != nil {
		log.Println(err)
		return false, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return false, err
	}
	if len(records) > 0 {
		return true, nil
	}
	return false, nil
}

func SelectFileMyList(pid, user string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
		SELECT 
			id, etag, name, mime, type, state, user, update_time 
		FROM 
			file 
		WHERE 
			pid=? 
		AND 
			user=?
		AND 
			state!=1 
		ORDER BY 
			update_time 
		DESC
		`, pid, user)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileDepartmentBucket(department string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
		SELECT f.id, f.etag, f.name, f.mime, f.type, f.state, f.user, f.update_time 
		FROM file f
		left join user u on u.id = f.user
		left join department d on d.id = u.department
		WHERE f.pid= 'departmentBucket'  
		AND f.state!=1 
		AND d.id=?
		ORDER BY 
			f.update_time 
		DESC
		`, department)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileFolderCompany() ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
		SELECT 
			id, etag, name, mime, type, state, user, update_time 
		FROM 
			file 
		WHERE 
			pid= "departmentBucket" 
		AND 
			type= "folder-company"
		AND 
			state!=1 
		ORDER BY 
			update_time 
		DESC
		`)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileByPidList(pid string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
		SELECT id, etag, name, mime, type, state, user, update_time 
		FROM file 
		WHERE pid=? 
		AND state!=1 
		ORDER BY update_time DESC
		`, pid)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileByUserList(user string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`
		SELECT id, etag, name, mime, type, state, user, update_time 
		FROM file 
		WHERE id=? 
		AND state!=1 
		ORDER BY update_time DESC
		`, user)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func HasFileFolderName(name, pid, user string) (bool, error) {
	rows, err := Db.Query(`SELECT id FROM file WHERE pid=? AND name=? AND user=? LIMIT 1`, pid, name, user)
	if err != nil {
		log.Println(err)
		return true, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return true, err
	}
	if len(records) != 0 {
		return true, nil
	}
	return false, nil
}

func SelectFileById(id string) (map[string]interface{}, error) {
	rows, err := Db.Query(`
		SELECT 
			id,etag,name,type,size,pid,state,mime,user
		FROM 
			file
		WHERE 
			id=? 
		AND 
			state = 0
		LIMIT 1
	`, id)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	if len(records) != 1 {
		return nil, nil
	}
	return records[0], nil
}

func SelectFileInfoById(id string) (map[string]interface{}, error) {
	rows, err := Db.Query(`
		SELECT 
			f.id,f.etag,f.name,f.type,f.size,f.pid,f.state,u.email 
		FROM 
			file f 
		LEFT JOIN 
			user u 
		ON 
			f.user = u.id
		WHERE 
			f.id=? 
		AND 
			f.state = 0
		LIMIT 1
	`, id)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	if len(records) != 1 {
		return nil, nil
	}
	return records[0], nil
}

func SelectFileDeleteList(user string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT id,etag,name,mime,type,update_time FROM file WHERE user=? AND state=1 ORDER BY update_time DESC`, user)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileTaskList(user string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT id,etag,name,mime,type,update_time FROM file WHERE user=? AND state=2 ORDER BY update_time DESC`, user)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileCheckFinish(ids []string, user string) ([]map[string]interface{}, error) {
	idStr := strings.Join(ids, "','")
	sqlRaw := fmt.Sprintf(`SELECT id FROM file WHERE user=? AND state=0 AND id IN ('%s')`, idStr)
	rows, err := Db.Query(sqlRaw, user)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileUploading(etags []string, user string) ([]map[string]interface{}, error) {
	idStr := strings.Join(etags, "','")
	sqlRaw := fmt.Sprintf(`
		SELECT 
			id, etag, name, mime, type, state, user, update_time, local 
		FROM 
			file 
		WHERE 
			user=? 
		AND state=2 
		AND local!='' 
		AND etag IN ('%s') 
		ORDER BY 
			update_time 
		DESC
	`, idStr)
	rows, err := Db.Query(sqlRaw, user)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func SelectFileByEtags(etags []string, user string) ([]map[string]interface{}, error) {
	idStr := strings.Join(etags, "','")
	sqlRaw := fmt.Sprintf(`
		SELECT 
			id, etag, name, mime, type, user, update_time, local 
		FROM 
			file 
		WHERE 
			user=? 
		AND etag IN ('%s') 
		ORDER BY 
			update_time 
		DESC
	`, idStr)
	rows, err := Db.Query(sqlRaw, user)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	return records, nil
}
