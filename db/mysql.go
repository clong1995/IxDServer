package db

import (
	. "IxDServer/common"
	. "IxDServer/config"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"reflect"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", CONF.DataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	Db.SetMaxOpenConns(2000)
	Db.SetMaxIdleConns(1000)
	err = Db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Query(sql string, dest interface{}, args ...interface{}) (interface{}, error) {
	log.Println(sql)
	log.Println(dest)
	log.Println(args)

	return nil, nil
}

func ParseRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(ROW_STR)
	}
	columnLen := len(columnTypes)

	scanArgs := make([]interface{}, columnLen)

	for i, v := range columnTypes {
		t := v.ScanType()
		v := reflect.New(t).Interface()
		scanArgs[i] = v
	}

	records := make([]map[string]interface{}, 0)

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf(ROW_STR)
		}
		record := make(map[string]interface{}, columnLen)

		for i, ct := range columnTypes {
			switch vv := scanArgs[i].(type) {
			//这特么竟然是字符串
			case *sql.RawBytes:
				record[ct.Name()] = string(*vv)
			//时间在这里直接格式化了
			case *mysql.NullTime:
				record[ct.Name()] = vv.Time.Format("2006-01-02 15:04:05")
			//特殊的空判断，这由于接受者是interface,不需要Valid
			case *sql.NullBool:
				value, _ := vv.Value()
				record[ct.Name()] = value
			case *sql.NullFloat64:
				value, _ := vv.Value()
				record[ct.Name()] = value
			case *sql.NullInt64:
				value, _ := vv.Value()
				record[ct.Name()] = value
			case *sql.NullString:
				value, _ := vv.Value()
				record[ct.Name()] = value
			//貌似是正常的，发现问题再处理
			//int
			case *int8:
				record[ct.Name()] = *vv
			case *uint8:
				record[ct.Name()] = *vv
			case *int32:
				record[ct.Name()] = *vv
			case *uint32:
				record[ct.Name()] = *vv
			case *int64:
				record[ct.Name()] = *vv
			case *uint64:
				record[ct.Name()] = *vv
			//float
			case *float32:
				record[ct.Name()] = *vv
			case *float64:
				record[ct.Name()] = *vv
			//未知
			default:
				log.Println("未知类型出现")
				record[ct.Name()] = vv
			}
		}
		records = append(records, record)
	}
	return records, nil
}

//过滤传入sql的参数
//TODO 这里需要判断类型
func FilterNull(args []interface{}) {
	for k, v := range args {
		if v == "" {
			args[k] = nil
		}
	}
}
