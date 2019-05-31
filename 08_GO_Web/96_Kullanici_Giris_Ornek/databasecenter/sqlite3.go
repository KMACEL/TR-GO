package databasecenter

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

/*
██████╗  █████╗ ████████╗ █████╗         ██████╗  █████╗ ███████╗███████╗         ██████╗███████╗███╗   ██╗████████╗███████╗██████╗
██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗        ██╔══██╗██╔══██╗██╔════╝██╔════╝        ██╔════╝██╔════╝████╗  ██║╚══██╔══╝██╔════╝██╔══██╗
██║  ██║███████║   ██║   ███████║        ██████╔╝███████║███████╗█████╗          ██║     █████╗  ██╔██╗ ██║   ██║   █████╗  ██████╔╝
██║  ██║██╔══██║   ██║   ██╔══██║        ██╔══██╗██╔══██║╚════██║██╔══╝          ██║     ██╔══╝  ██║╚██╗██║   ██║   ██╔══╝  ██╔══██╗
██████╔╝██║  ██║   ██║   ██║  ██║        ██████╔╝██║  ██║███████║███████╗        ╚██████╗███████╗██║ ╚████║   ██║   ███████╗██║  ██║
╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝        ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝         ╚═════╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚══════╝╚═╝  ╚═╝
*/
// Database Ceter is used for database util. "SQLite3" has been selected as the database.
// It is aimed at fast, light, easy to use. Driver "github.com/mattn/go-sqlite3" has been selected.
// You need to perform the necessary downloads before using the program.
//     sudo apt-get install sqlite3
//     go get github.com/mattn/go -sqlite3

// The DB was created to use the functions of the database center package.
type DB struct{}

// Driver Name Constant
const (
	driverName = "sqlite3"
)

/*
 ██████╗ ██████╗ ███████╗███╗   ██╗
██╔═══██╗██╔══██╗██╔════╝████╗  ██║
██║   ██║██████╔╝█████╗  ██╔██╗ ██║
██║   ██║██╔═══╝ ██╔══╝  ██║╚██╗██║
╚██████╔╝██║     ███████╗██║ ╚████║
╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═══╝
*/

// Open function is used to open and use a database. This function returns data as "* sql.DB".
// Through this data, transactions are carried out. The returned data is a ready-to-use database.
// If a named database does not exist on the specified path, it creates a new database.
//   For Use Example :
//      db := dataBase.Open("foo.db")
func (d DB) Open(dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ████████╗ █████╗ ██████╗ ██╗     ███████╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ╚══██╔══╝██╔══██╗██╔══██╗██║     ██╔════╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗             ██║   ███████║██████╔╝██║     █████╗
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝             ██║   ██╔══██║██╔══██╗██║     ██╔══╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗           ██║   ██║  ██║██████╔╝███████╗███████╗
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝           ╚═╝   ╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝
*/

// CreateTable
func (d DB) CreateTable(db *sql.DB, tableName string, items string) {
	sqlStmt := "CREATE TABLE " + tableName + "(" + items + ");"
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

/*
██╗███╗   ██╗███████╗███████╗██████╗ ████████╗        ██╗███╗   ██╗████████╗ ██████╗
██║████╗  ██║██╔════╝██╔════╝██╔══██╗╚══██╔══╝        ██║████╗  ██║╚══██╔══╝██╔═══██╗
██║██╔██╗ ██║███████╗█████╗  ██████╔╝   ██║           ██║██╔██╗ ██║   ██║   ██║   ██║
██║██║╚██╗██║╚════██║██╔══╝  ██╔══██╗   ██║           ██║██║╚██╗██║   ██║   ██║   ██║
██║██║ ╚████║███████║███████╗██║  ██║   ██║           ██║██║ ╚████║   ██║   ╚██████╔╝
╚═╝╚═╝  ╚═══╝╚══════╝╚══════╝╚═╝  ╚═╝   ╚═╝           ╚═╝╚═╝  ╚═══╝   ╚═╝    ╚═════╝
*/

func (d DB) InsertInto(db *sql.DB, tableName string, column string, args ...interface{}) {

	sql := "INSERT INTO " + tableName + "(" + column + ") values( " + strings.SplitAfterN(strings.Repeat(",?", len(args)), ",", 2)[1] + ")"
	_, err := d.GenericQuery(db, sql, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func (d DB) InsertQuery(db *sql.DB, tableName string, column string, values string) {
	_, err := db.Exec("INSERT INTO " + tableName + "(" + column + ") values" + values)
	if err != nil {
		log.Fatal(err)
	}
}

/*
███████╗███████╗██╗     ███████╗ ██████╗████████╗
██╔════╝██╔════╝██║     ██╔════╝██╔════╝╚══██╔══╝
███████╗█████╗  ██║     █████╗  ██║        ██║
╚════██║██╔══╝  ██║     ██╔══╝  ██║        ██║
███████║███████╗███████╗███████╗╚██████╗   ██║
╚══════╝╚══════╝╚══════╝╚══════╝ ╚═════╝   ╚═╝
*/

// Select is
func (d DB) Select(db *sql.DB, tableName string) []map[string]interface{} {
	sql := "select * from " + tableName
	users, err := d.GenericQuery(db, sql)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

func (d DB) Find(db *sql.DB, tableName string, findColumn string, where string) map[string]interface{} {
	sql := "select * from " + tableName + " where " + findColumn + "= ?"
	users, err := d.GenericQuery(db, sql, where)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if len(users) > 0 {
		return users[0]
	} else {
		return nil
	}
}

func (d DB) SelectAll(db *sql.DB, tableName string) {
	sql := "select * from " + tableName
	users, err := d.GenericQuery(db, sql)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	for i, user := range users {
		fmt.Print(i)
		for j, u := range user {
			v := reflect.ValueOf(u)
			switch v.Kind() {
			case reflect.Int, reflect.Int64:
				fmt.Print(" - ", j, " : ", u)
			case reflect.Int32:
				fmt.Print(" , ", j, " : ", string(u.([]byte)))
			default:
				fmt.Println(" - ", j, " : ", u)
			}
		}
		fmt.Println()
	}
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝
*/

// Update is
func (d DB) Update(db *sql.DB, tableName string, column string, changeValue string, whereColumn string, where string) {
	_, err := db.Exec("UPDATE " + tableName + " SET " + column + "= '" + changeValue + "' WHERE " + whereColumn + "=" + where)
	if err != nil {
		log.Fatal(err)
	}
}

/*
██████╗ ███████╗██╗     ███████╗████████╗███████╗        ████████╗ █████╗ ██████╗ ██╗     ███████╗
██╔══██╗██╔════╝██║     ██╔════╝╚══██╔══╝██╔════╝        ╚══██╔══╝██╔══██╗██╔══██╗██║     ██╔════╝
██║  ██║█████╗  ██║     █████╗     ██║   █████╗             ██║   ███████║██████╔╝██║     █████╗
██║  ██║██╔══╝  ██║     ██╔══╝     ██║   ██╔══╝             ██║   ██╔══██║██╔══██╗██║     ██╔══╝
██████╔╝███████╗███████╗███████╗   ██║   ███████╗           ██║   ██║  ██║██████╔╝███████╗███████╗
╚═════╝ ╚══════╝╚══════╝╚══════╝   ╚═╝   ╚══════╝           ╚═╝   ╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝
*/
//DeleteTable is
func (d DB) DeleteTable(db *sql.DB, tableName string) {
	_, err := db.Exec("delete from " + tableName)
	if err != nil {
		log.Fatal(err)
	}
}

/*
 ██████╗ ██╗   ██╗███████╗██████╗ ██╗   ██╗
██╔═══██╗██║   ██║██╔════╝██╔══██╗╚██╗ ██╔╝
██║   ██║██║   ██║█████╗  ██████╔╝ ╚████╔╝
██║▄▄ ██║██║   ██║██╔══╝  ██╔══██╗  ╚██╔╝
╚██████╔╝╚██████╔╝███████╗██║  ██║   ██║
 ╚══▀▀═╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝   ╚═╝
*/

// Query is
func (d DB) Query(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func (d DB) GenericQuery(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	res := make([]map[string]interface{}, 0)

	for rows.Next() {
		container := make([]interface{}, len(cols))
		dest := make([]interface{}, len(cols))
		for i, _ := range container {
			dest[i] = &container[i]
		}
		rows.Scan(dest...)
		r := make(map[string]interface{})
		for i, colname := range cols {
			val := dest[i].(*interface{})
			r[colname] = *val
		}
		res = append(res, r)
	}

	return res, nil
}

/*
 ██████╗██╗      ██████╗ ███████╗███████╗
██╔════╝██║     ██╔═══██╗██╔════╝██╔════╝
██║     ██║     ██║   ██║███████╗█████╗
██║     ██║     ██║   ██║╚════██║██╔══╝
╚██████╗███████╗╚██████╔╝███████║███████╗
 ╚═════╝╚══════╝ ╚═════╝ ╚══════╝╚══════╝
*/

// Close is
func (d DB) Close(db *sql.DB) {
	db.Close()
}

func example() {

	var dataBase DB

	os.Remove("foo.db")

	db := dataBase.Open("foo.db")

	dataBase.CreateTable(db, "bar",
		"id INTEGER NOT NULL PRIMARY KEY,"+
			"name TEXT")

	dataBase.InsertQuery(db, "bar", "id,name", "(4, 'mert'), (5, 'acel'), (6, 'ek')")

	dataBase.Update(db, "bar", "name", "caca", "id", "5")

	dataBase.Query(db, "insert into bar (id,name) values (88,'Nettin')")

	getSelect := dataBase.Select(db, "bar")[0]["name"].([]byte)
	fmt.Println(string(getSelect))

	getFind := dataBase.Find(db, "bar", "id", "5")["name"].([]byte)
	fmt.Println(string(getFind))

	dataBase.InsertInto(db, "bar", "id,name", 22, "hahaha")

	dataBase.SelectAll(db, "foo")

	dataBase.Close(db)
}
