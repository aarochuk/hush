package pass

import (
	"fmt"
	"log"
	"os"
)

func checkHushStore() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
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
