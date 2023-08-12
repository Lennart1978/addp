package main

import (
	"fmt"
	"os"
	"strings"

	"example.com/addp/misc"
)

const (
	bashrcFilename = ".bashrc"
)

func main() {
	greetings()

	if len(os.Args) < 2 {
		fmt.Printf("\nUsage: addp PATH\n")
		os.Exit(1)
	}

	fmt.Println("Current $PATH:")
	printPath()

	pathToAdd := os.Args[1]
	if appendBashrc(pathToAdd) {
		fmt.Printf("\nSuccessfully added %s to $PATH!\n", pathToAdd)
	} else {
		fmt.Println("Something went wrong!")
	}
}

func greetings() {
	fmt.Printf("Hello %s, welcome to Lennart's addp.\n", misc.GetUserName())
	fmt.Println("-------------------------------------------")
}

func printPath() {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	fmt.Printf("%d entries\n", len(paths))

	for _, p := range paths {
		fmt.Println(p)
	}
}

func appendBashrc(path string) bool {
	bashrc := fmt.Sprintf("/home/%s/%s", strings.ToLower(misc.GetUserName()), bashrcFilename)
	f, err := os.OpenFile(bashrc, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer f.Close()

	newEntry := fmt.Sprintf("\n# Added by addp!\nexport PATH=$PATH:%s", path)
	if _, err := f.WriteString(newEntry); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
