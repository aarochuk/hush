package main

import (
	"flag"
	"fmt"
	"os"

	I "github.com/aarochuk/hush/internal"
)

func main() {
	// TODO: add as much error checking as possible for every
	// single error situation you can think of add details with each flag
	insertCmd := flag.NewFlagSet("insert", flag.ExitOnError)
	insertMulti := insertCmd.Bool("m", false, "Use this to make a multi line password")
	switch os.Args[1] {
	case "insert":
		if len(os.Args) < 3 {
			fmt.Println("To use this command you need to enter the name of the folder/file where you want to name this password.")
		} else {
			// TODO: perform all necessary validations on
			// the input including ensuring file doesnt exist
			// yet and it is a valid folder/filename
			I.CreatePassword(os.Args[2], *insertMulti)
		}
		I.CreatePassword()
	}
}
