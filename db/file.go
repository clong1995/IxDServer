package db

import (
	. "IxDServer/common"
	"fmt"
	"log"
	"strings"
)

func InsertFileFolder(id, name, pid, user string) error {
	_, err := Db.Exec("INSERT INTO file (id, name,type, pid, user) values (?,?,?,?,?)", id, name, "folder", pid, user)
	if err != nil {
		log.Println(err)
		return fmt.Errorf(SQL_STR)
	}
	return nil
}

func InsertFile(id, etag, name, pid, typee, user, mime, local string, size float64, state int) error {
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

func SelectFileList(pid, user string) ([]map[string]interface{}, error) {
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

func FileInfo(id string) (map[string]interface{}, error) {
	rows, err := Db.Query(`SELECT id,name,type,size,pid,state,user FROM file WHERE id=? LIMIT 1`, id)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(SQL_STR)
	}
	defer rows.Close()

	records, err := ParseRows(rows)
	if err != nil {
		return nil, err
	}
	if len(records) != 0 {
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
