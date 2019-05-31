package databasecenter

import (
	"database/sql"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var (
	dataBase DB
	db       *sql.DB
)

const (
	dataBaseName = "user.db"
	tableName    = "user"
)

func StartDatabase() bool {
	log.SetOutput(os.Stdout)
	return databaseConnection()
}

func databaseConnection() bool {
	if _, err := os.Stat(dataBaseName); !os.IsNotExist(err) {
		db = dataBase.Open(dataBaseName)
		//defer dataBase.Close(db)
		log.Println("DB is Open...")
		return true
	} else {
		db = dataBase.Open(dataBaseName)
		//defer dataBase.Close(db)
		createDatabse(db)
		log.Println("DB & Table is Created...")
		return true
	}
	return false
}

func createDatabse(db *sql.DB) {
	dataBase.CreateTable(db, tableName,
		"userid INTEGER PRIMARY KEY AUTOINCREMENT,"+
			"username varchar(20),"+
			"password varchar(65),"+
			"permission varchar(10)")
}

// GetDB is
func GetDB() *sql.DB {
	return db
}

// AddUser is
func AddUser(userName string, password string, permission string) {
	passwordCrypt, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	dataBase.InsertInto(db, tableName, "username,password,permission", userName, passwordCrypt, permission)
}

// CheckAuthentication is
func CheckAuthentication(username string, password string) bool {
	getData := dataBase.Find(db, tableName, "username", username)
	if getData != nil && getData["password"] != nil && len(getData["password"].([]uint8)) > 0 {
		if err := bcrypt.CompareHashAndPassword(getData["password"].([]uint8), []byte(password)); err == nil {
			return true
		}
	}
	return false
}

func CheckAdmin(username string, password string) bool {
	getData := dataBase.Find(db, tableName, "username", username)
	if getData != nil && getData["password"] != nil && len(getData["password"].([]uint8)) > 0 {
		if err := bcrypt.CompareHashAndPassword(getData["password"].([]uint8), []byte(password)); err == nil {
			return true
		}
	}
	return false
}
