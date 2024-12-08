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
	if dbInit() {

		homeDir, err := os.UserHomeDir()
		dir := homeDir + "/.hush/.passwords"
		db, err = sql.Open("sqlite3", dir)
		if err != nil {
			fmt.Println("Could not open passwords database")
		}
		defer db.Close()
		rows, err := db.Query("SELECT * FROM passwords")
		if err != nil {
			fmt.Println("Error querying database")
			return
		}
		defer rows.Close()
		data := []Pass{}
		for rows.Next() {
			var pass Pass
			if err := rows.Scan(&pass.ID, &pass.Name, &pass.Username, &pass.Password); err != nil {
				fmt.Println("Unable to get data from database")
				return
			}
			data = append(data, pass)
		}
		if len(data) == 0 {
			fmt.Println("The database is empty right now.")
		} else {
			fmt.Printf("%20s|%20s|%20s\n", "Password Name", "Username", "Password")
			bottom_pipes := strings.Repeat("_", 65)
			fmt.Printf("%65s\n", bottom_pipes)
			for _, passw := range data {
				fmt.Printf("%20s|%20s|%20s\n", passw.Name, passw.Username, passw.Password)
			}
		}
	} else {
		fmt.Println("Database was not initialized and could not be initialized")
		return
	}
}

func ShowOnePassword(location string) {
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
		rows, err := db.Query("SELECT id, pName, username, password FROM passwords WHERE pName=? AND username=?", pname, username)
		if err != nil {
			fmt.Println("Error searching for password in database")
			return
		}
		defer rows.Close()
		data := []Pass{}
		for rows.Next() {
			var pass Pass
			if err := rows.Scan(&pass.ID, &pass.Name, &pass.Username, &pass.Password); err != nil {
				fmt.Println("Unable to get data from database")
				return
			}
			data = append(data, pass)
		}

		if len(data) == 0 {
			fmt.Println("The password you were looking for did not exist.")
		} else {
			fmt.Printf("The password to %s is %s", location, data[0].Password)
		}
	} else {
		fmt.Println("Database was not initialized and could not be initialized")
		return
	}
}

func SaveNewPassword(location, word string) bool {
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

		rows, err := db.Query("SELECT id, pName, username, password FROM passwords WHERE pName=?", pname)
		if err != nil {
			fmt.Println("Error searching for password in database")
			return false
		}
		defer rows.Close()
		data := []Pass{}
		for rows.Next() {
			var pass Pass
			if err := rows.Scan(&pass.ID, &pass.Name, &pass.Username, &pass.Password); err != nil {
				fmt.Println("Unable to get data from database")
				return false
			}
			data = append(data, pass)
		}

		if len(data) != 0 {
			fmt.Println("The password name already exist in the database sorry either edit the password or create a password with a unique name.")
			return false
		}
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

func RemovePassword(location string) bool {
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
		//username := str[1]

		fmt.Print("Are you sure you want to delete your password (y/n): ")
		var input string
		_, er := fmt.Scan(&input)
		if er != nil {
			fmt.Println("Could not get input to whether you wanted your password deleted", err)
			return false
		}
		if input == "y" {
			rows, err := db.Query("SELECT id, pName, username, password FROM passwords WHERE pName=?", pname)
			if err != nil {
				fmt.Println("Error searching for password in database")
				return false
			}
			defer rows.Close()
			data := []Pass{}
			for rows.Next() {
				var pass Pass
				if err := rows.Scan(&pass.ID, &pass.Name, &pass.Username, &pass.Password); err != nil {
					fmt.Println("Unable to get data from database")
					return false
				}
				data = append(data, pass)
			}
			if len(data) > 0 {
				_, err := db.Exec("DELETE FROM passwords WHERE pName = ?", pname)
				if err != nil {
					fmt.Println("Error deleting password from database")
					return false
				} else {
					fmt.Println("Successfully deleted password from database")
				}
			} else {
				fmt.Println("The Name/Username you just entered does not exist in the database.")
			}
		}
	} else {
		fmt.Println("Sorry your database was not initialized and could not be initialized, please try again later.")
		return false
	}
	return true
}

func EditPassword(location, pw string) bool {
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
		rows, err := db.Query("SELECT id, pName, username, password FROM passwords WHERE pName=? AND username=?", pname, username)
		if err != nil {
			fmt.Println("Error searching for password in database")
			return false
		}
		defer rows.Close()
		data := []Pass{}
		for rows.Next() {
			var pass Pass
			if err := rows.Scan(&pass.ID, &pass.Name, &pass.Username, &pass.Password); err != nil {
				fmt.Println("Unable to get data from database")
				return false
			}
			data = append(data, pass)
		}

		if len(data) == 0 {
			fmt.Println("The password you wanted to edit does not exist.")
		} else {
			_, err := db.Exec("UPDATE passwords SET password = ? WHERE pName = ? AND username = ?", pw, pname, username)
			if err != nil {
				fmt.Println("Could not update password in database, please try again later")
				return false
			}
			fmt.Printf("Successfully updated password for %s", location)
		}
	} else {
		fmt.Println("Database was not initialized and could not be initialized")
		return false
	}
	return true
}
