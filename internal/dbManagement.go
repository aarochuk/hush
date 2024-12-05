package pass

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Pass struct {
	ID       int
	Name     string
	Username string
	Password string
}

var db *sql.DB

func dbInit() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not create a hush directory where your passwords would be stored.")
		return false
	}
	dir := homeDir + "/.hush"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("A hush directory does not exist for you yet, attempting to create one now.")
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Println("Error creating hush directory, ", err)
			return false
		}
		db_location := dir + "/.passwords"
		f, err := os.Create(db_location)
		if err != nil {
			fmt.Println("Error creating password database, ", err)
			return false
		}
		f.Close()
		db, err = sql.Open("sqlite3", db_location)
		if err != nil {
			fmt.Println("Error when opening database", err)
			return false
		}
		defer db.Close()
		createTable := `
CREATE TABLE IF NOT EXISTS passwords(
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
pName TEXT NOT NULL,
username TEXT,
password TEXT NOT NULL 
		)
		`
		_, err = db.Exec(createTable)
		if err != nil {
			fmt.Println("Error creating password table, ", err)
			return false
		}
		fmt.Println("Your database was successfully created")
		fmt.Println("Since you just created your database you will create a username and password to secure your passwords.")
		if createSuperUser() {
			fmt.Println("Super user successfully created. You can now use all the features of hush.")
		} else {
			fmt.Println("There was an error while creating the super user, please try again later.")
			return false
		}
	}
	return true
}

func saveSuperUser(username, password string) bool {
	//TODO: fix this according to changes made to user schema and all that jazz

	homeDir, err := os.UserHomeDir()
	dir := homeDir + "/.hush/.passwords"
	db, err = sql.Open("sqlite3", dir)
	if err != nil {
		fmt.Println("Could not open passwords database")
	}
	defer db.Close()

	_, e := db.Exec("INSERT INTO passwords(pName, username, password) VALUES(?, ?, ?);", "user 1", username, password)
	if e != nil {
		fmt.Println("Error when creating superuser, ", err)
		return false
	}
	return true
}

func ShowPasswords() {
	valid := dbInit()
	if valid {
		// TODO: Create a way to display all the passwords once you have passwords in the store
		fmt.Println("temporary holder")
	} else {
		fmt.Println("Could not find the hush_store folder or create one.")
	}
}

func SaveNewPassword(location, word string) bool {
	// TODO: work on this
	if dbInit() {
		homeDir, err := os.UserHomeDir()
		dir := homeDir + "/.hush/.passwords"
		db, err = sql.Open("sqlite3", dir)
		if err != nil {
			fmt.Println("Could not open passwords database")
		}
		defer db.Close()
		str := strings.Split(location, "/")
		pname := str[0]
		username := str[1]
		fmt.Println(username)
		_, e := db.Exec("INSERT INTO passwords(pName, username, password) VALUES(?, ?, ?);", pname, username, word)
		if e != nil {
			fmt.Println("Error when adding password, ", err)
			return false
		}
		fmt.Println("Successfully added password for", location)
	} else {
		fmt.Println("Sorry your database was not initialized and could not be initialized, please try again later.")
		return false
	}
	return true
}
