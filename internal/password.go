package pass

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GeneratePassword(location string, length int, special, letters, nus, edit bool) bool {
	pw := make([]byte, length)
	var types []int
	if !special {
		types = append(types, 0)
	}
	if !letters {
		types = append(types, 1)
	}
	if !nus {
		types = append(types, 2)
	}
	for i := range pw {
		curr := types[rand.Intn(len(types))]
		if curr == 0 {
			pw[i] = specialBytes[rand.Intn(len(specialBytes))]
		} else if curr == 1 {
			pw[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else if curr == 2 {
			pw[i] = numBytes[rand.Intn(len(numBytes))]
		}
	}
	if edit {
		return EditPassword(location, string(pw))
	}
	return SaveNewPassword(location, string(pw))
}

func CreatePassword(location string, multi, edit bool) bool {
	word := ""
	if multi {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter your password (press Ctrl+D or Ctrl+Z to end): ")
		for scanner.Scan() {
			word += scanner.Text() + "\n"
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error when getting multiline password, error message: ", err)
			return false
		}
	} else {
		fmt.Print("Enter your password: ")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error when getting password, error message: ", err)
			return false
		}
		word += input
	}
	if edit {
		return EditPassword(location, word)
	} else {
		return SaveNewPassword(location, word)
	}
}

func createSuperUser() bool {
	pword := ""
	uname := ""
	fmt.Print("Enter your username: ")
	var u_put string
	_, err := fmt.Scan(&u_put)
	if err != nil {
		fmt.Println("Error when getting username, error message: ", err)
		return false
	}
	uname += u_put

	fmt.Print("Enter your password: ")
	var p_put string
	_, e := fmt.Scan(&p_put)
	if e != nil {
		fmt.Println("Error when getting password, error message: ", e)
		return false
	}
	pword += p_put
	return saveSuperUser(uname, pword)
}
