package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

var _cfg = beego.AppConfig

// db init
func Db() *sql.DB {
	// database config
	dbUser := _cfg.String("mysql_user")
	dbPass := _cfg.String("mysql_pass")
	dbHost := _cfg.String("mysql_host")
	dbPort := _cfg.String("mysql_port")
	dbName := _cfg.String("mysql_db")

	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FShanghai"

	db, err := sql.Open("mysql", dbLink)
	if err != nil {
		panic(err.Error())
		log.Printf("mysql connect error %v", err)
		return nil
	}
	//defer db.Close()

	return db
}

//插入
func Create(sqlstr string, args ...interface{}) (int64, error) {
	db := Db()
	defer db.Close()

	stmtIns, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	return result.LastInsertId()
}

//取一行数据，注意这类取出来的结果都是 string
func GetOne(sqlstr string, args ...interface{}) (*map[string]string, error) {
	db := Db()
	defer db.Close()

	stmtOut, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	res := make(map[string]string, len(scanArgs))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			res[columns[i]] = value
		}
		break //get the first row only
	}

	return &res, nil
}

//取多行，注意这类取出来的结果都是string
func GetAll(sqlstr string, args ...interface{}) (*[]map[string]string, error) {
	db := Db()
	defer db.Close()

	stmtOut, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	res := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		res = append(res, vmap)
	}
	return &res, nil
}

// 修改
func Update(sqlstr string, args ...interface{}) (int64, error) {
	db := Db()
	defer db.Close()

	stmtIns, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	return result.RowsAffected()
}

// 修改
func Delete(sqlstr string, args ...interface{}) (int64, error) {
	db := Db()
	defer db.Close()

	stmtIns, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	return result.RowsAffected()
}

// Exec: Update or Delete
func Exec(sqlstr string, args ...interface{}) (int64, error) {
	db := Db()
	defer db.Close()

	stmtIns, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	return result.RowsAffected()
}
