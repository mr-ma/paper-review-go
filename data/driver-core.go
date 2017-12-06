package data

import (
	"database/sql"
	"fmt"
	//overriding MySqlDriver
	_ "../mysql"
)

//Driver is the minimum functions for a DB engine
type DriverCore interface {
	OpenDB() (*sql.DB, error)
	Query(query string) (*sql.DB, *sql.Stmt, error)
	Insert(tableName string, columns string, values ...interface{}) (affected int64, id int64, err error)
}

//MySQLDriver mysql startup settings
type MySQLDriver struct {
	username string
	pass     string
	database string
}


//OpenDB opens a db connection
func (d MySQLDriver) OpenDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", d.username+":"+d.pass+"@/"+d.database)
	if err != nil {
		checkErr(err)// Just for example purpose. You should use proper error handling instead of panic
	}
	return db, err
}

//Query queries the db according to the query string
func (d MySQLDriver) Query(query string) (*sql.DB, *sql.Stmt, error) {
	db, err := d.OpenDB()
	if err != nil {
		checkErr(err)
	}
	//defer db.Close()
	fmt.Println(query)
	stmtOut, err := db.Prepare(query)
	checkErr(err)
	return db, stmtOut, err
}


//Insert general insert function
func (d MySQLDriver) Insert(tableName string, columns string, values ...interface{}) (affected int64, id int64, err error) {
	db, err := d.OpenDB()
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	fmt.Printf("tablename:%+v columns:%+v values:%+v\n", tableName, columns, values)

	stmt, err := db.Prepare("INSERT " + tableName + " SET " + columns)
	checkErr(err)
	res, err := stmt.Exec(values...)
	checkErr(err)
	fmt.Printf("res:%+v \n", res)
	id, err = res.LastInsertId()
	affect, err := res.RowsAffected()
	checkErr(err)
	return affect, id, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
