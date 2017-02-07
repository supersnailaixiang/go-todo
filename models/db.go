package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	link_id *sql.DB
	tx      *sql.Tx
	result  sql.Result
}

// mysql 全局函数
func db() *Mysql {

	dbms := beego.AppConfig.String("dbms")
	dbuser := beego.AppConfig.String("dbuser")
	dbhost := beego.AppConfig.String("dbhost")
	dbport, err := beego.AppConfig.Int("dbport")
	dbname := beego.AppConfig.String("dbname")
	dbpwd := beego.AppConfig.String("dbpwd")
	if nil != err {
		dbport = 3306
	}
	db, err := sql.Open(dbms, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbuser, dbpwd, dbhost, dbport, dbname))

	mysql := new(Mysql)
	mysql.link_id = db
	return mysql
}
func (db *Mysql) begin() {
	tx, err := db.link_id.Begin()
	checkErr(err)
	db.tx = tx

}
func (db *Mysql) rollBack() {
	db.tx.Rollback()
}
func (db *Mysql) commit() {
	db.tx.Commit()
}
func (db *Mysql) prepareExecute(sql string, args ...interface{}) bool {
	stmt, err := db.link_id.Prepare(sql)
	checkErr(err)
	result, err := stmt.Exec(args[:]...)
	checkErr(err)
	db.result = result
	return true
}
func (db *Mysql) lastInsertId() int64 {
	id, _ := db.result.LastInsertId()
	return id
}
func (db *Mysql) prepareQueryRow(sql string, args ...interface{}) *sql.Row {
	stmt, err := db.link_id.Prepare(sql)
	checkErr(err)

	row := stmt.QueryRow(args[:]...)

	return row
}
func (db *Mysql) prepareQuery(sql string, args ...interface{}) *sql.Rows {
	stmt, err := db.link_id.Prepare(sql)
	checkErr(err)
	rows, err := stmt.Query(args[:]...)
	checkErr(err)
	return rows
}
func (db *Mysql) close() {
	err := db.link_id.Close()
	checkErr(err)
}

// 需要在外部判断，不需要在这里做回滚
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
