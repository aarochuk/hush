package pass

import (
	"fmt"
	"os"
)

func checkHushStore() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not create a hush_store directory.")
		return false
	}
	dir := homeDir + "/.hush_store"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("A hush_store directory does not exist for you yet, attempting to create one now.")
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}

func ShowPasswords() {
	valid := checkHushStore()
	if valid {
		// TODO: Create a way to display all the passwords once you have passwords in the store
		fmt.Println("temporary holder")
	} else {
		fmt.Println("Could not find the hush_store folder or create one.")
	}
}

func SaveNewPassword(location, word string) bool {

}
