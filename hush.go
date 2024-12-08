package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	I "github.com/aarochuk/hush/internal"
)

func main() {
	// TODO: add as much error checking as possible for every
	// single error situation you can think of add details with each flag
	insertCmd := flag.NewFlagSet("insert", flag.ExitOnError)
	insertMulti := insertCmd.Bool("m", false, "Use this to make a multi line password")

	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	noSymbols := generateCmd.Bool("ns", false, "Use this flag if you dont want symbols in your password")
	noLetters := generateCmd.Bool("nl", false, "Use this flag if you dont want letters in your password")
	noNumbers := generateCmd.Bool("nn", false, "Use this flag if you dont want numbers in your password")

	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	genEdit := editCmd.Bool("g", false, "Use this flag if you want your new password to be generated rather than set by you")
	noS := editCmd.Bool("ns", false, "Use this flag if you dont want symbols in your password")
	noL := editCmd.Bool("nl", false, "Use this flag if you dont want letters in your password")
	noN := editCmd.Bool("nn", false, "Use this flag if you dont want numbers in your password")
	//removeCmd := flag.NewFlagSet("rm", flag.ExitOnError)
	if len(os.Args) == 1 {
		I.ShowPasswords()
	} else if len(os.Args) == 2 {
		I.ShowOnePassword(os.Args[1])
	} else {
		switch os.Args[1] {
		case "insert":
			if len(os.Args) < 3 {
				fmt.Println("To use this command you need to enter the name of the folder/file where you want to name this password.")
			} else {
				// TODO: perform all necessary validations on
				// the input including ensuring file doesnt exist
				// yet and it is a valid folder/filename
				insertCmd.Parse(os.Args[3:])
				I.CreatePassword(os.Args[2], *insertMulti, false)
			}
		case "generate":

			if len(os.Args) < 4 {
				fmt.Println("To use this command you need to enter the name of the folder/file where you want to name this password.")
			} else {
				// TODO: perform all necessary validations on
				generateCmd.Parse(os.Args[4:])
				plen, err := strconv.Atoi(os.Args[3])
				if err != nil {
					fmt.Println("Please enter an integer to be the length of the password")
				} else {
					I.GeneratePassword(os.Args[2], plen, *noSymbols, *noLetters, *noNumbers, false)
				}
			}

		case "delete":
			I.RemovePassword(os.Args[2])

		case "edit":
			editCmd.Parse(os.Args[3:])
			if *genEdit {
				//TODO: Fix gen password for edit it does not work right now
				plen, err := strconv.Atoi(os.Args[4])
				if err != nil {
					fmt.Println("Please enter an integer to be the length of the password")
				} else {
					I.GeneratePassword(os.Args[2], plen, *noS, *noL, *noN, true)
				}
			} else {
				I.CreatePassword(os.Args[2], *insertMulti, true)
			}

		default:
			fmt.Println("Please enter a legitimate command to use this program, use the -h flag to get help and details on all that you can do.")
		}
	}
}
