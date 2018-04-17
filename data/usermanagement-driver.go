package data

import (
	"strconv"
	"strings"
	//overriding MySqlDriver
	//"github.com/go-sql-driver/mysql"
	"../model"
)

type UsermanagementDriver interface {
	DriverCore
	Login(string, string) (model.LoginResult, error)
	QueryDB(string) (model.Result, error)
	SaveUser(string, string) (model.Result, error)
	CreateUser(string, string, int) (model.Result, error)
	UpdateUser(string, int) (model.Result, error)
	DeleteUser(string) (model.Result, error)
	GetUser(string) (model.User, error)
	GetUsers() ([]model.User, error)
	GetTaxonomyPermissions(string) ([]model.Taxonomy, error)
	UpdateTaxonomyPermissions(string, string) (model.Result, error)
}

//InitMySQLDriver initialize a new my sql driver instance
func InitUsermanagementDriver(user string, password string, server string) UsermanagementDriver {
	return MySQLDriver{username: user, pass: password, database: "classification", server: server}
}

func (d MySQLDriver) Login(email string, password string) (result model.LoginResult, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT email, taxonomies, admin FROM user WHERE email = ? AND password = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email, password)
	checkErr(err)
	if rows.Next() {
		a := model.User{}
		rows.Scan(&a.Email, &a.Taxonomies, &a.Admin)
		result.Success = true
		result.User = a
	} else {
		result.Success = false
	}
	defer rows.Close()
	return result, err
}

func (d MySQLDriver) QueryDB(query string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query(query)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		result.Success = false
	} else {
		result.Success = true
	}
	return result, err
}

func (d MySQLDriver) SaveUser(email string, password string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT COUNT(email) as userCount FROM user WHERE email = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email)
	checkErr(err)
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	defer rows.Close()
	if count > 0 {
		result.Success = false
		return result, err
	}
	dbRef.Exec("INSERT IGNORE INTO user (email, name, password) VALUES (?, ?, ?);", email, email, password)
	result.Success = true
	return result, err
}

func (d MySQLDriver) CreateUser(email string, password string, admin int) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT COUNT(email) as userCount FROM user WHERE email = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email)
	checkErr(err)
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	defer rows.Close()
	if count > 0 {
		result.Success = false
		return result, err
	}
	adminStr := strconv.Itoa(admin)
	dbRef.Exec("INSERT IGNORE INTO user (email, name, password, admin) VALUES (?, ?, ?, ?);", email, email, password, adminStr)
	result.Success = true
	return result, err
}

func (d MySQLDriver) UpdateUser(email string, admin int) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	adminStr := strconv.Itoa(admin)
	dbRef.Exec("UPDATE user SET admin = ? WHERE email = ?;", adminStr, email)
	result.Success = true
	return result, err
}

func (d MySQLDriver) DeleteUser(email string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	dbRef.Exec("DELETE FROM user WHERE email = ?;", email)
	result.Success = true
	return result, err
}

func (d MySQLDriver) GetUser(email string) (user model.User, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT id, name FROM user WHERE email = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email)
	checkErr(err)
	if rows.Next() {
		a := model.User{}
		rows.Scan(&a.ID, &a.Name)
		user = a
	}
	defer rows.Close()
	user.Email = email
	return user, err
}

func (d MySQLDriver) GetUsers() (users []model.User, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT email, admin FROM user;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		a := model.User{}
		rows.Scan(&a.Email, &a.Admin)
		users = append(users, a)
	}
	defer rows.Close()
	return users, err
}

func (d MySQLDriver) GetTaxonomyPermissions(email string) (taxonomies []model.Taxonomy, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT taxonomies FROM user WHERE email = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email)
	checkErr(err)
	var taxonomyPermissions string
	taxonomyPermissions = ""
	for rows.Next() {
		rows.Scan(&taxonomyPermissions)
	}
	if taxonomyPermissions != "" {
		array := strings.Split(taxonomyPermissions, ",")
		for _, elem := range array {
			id, err := strconv.Atoi(elem)
			if err == nil {
				a := model.Taxonomy{ID: id}
				taxonomies = append(taxonomies, a)
			}
		}
	}
	defer rows.Close()
	return taxonomies, err
}

func (d MySQLDriver) UpdateTaxonomyPermissions(email string, permissions string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	dbRef.Exec("UPDATE user SET taxonomies = ? WHERE email = ?;", permissions, email)
	result.Success = true
	return result, err
}