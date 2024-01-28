package main

import (
	"fmt"
	"os"
)

func readConfig(config string, dir string) {
}


func main() {
	// get $HOME
	home := os.Getenv("HOME")

    // set config as our configuration file
	config := home + "/.context"

	// check for our configuration file $HOME/.context
	if _, err := os.Stat(config); os.IsNotExist(err) {
		fmt.Println("Config file found: ", config) 
		os.Exit(1)
	}

    // get current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Couldn't get current working directory: ", err)
		os.Exit(3)
	}

    // get the command to run, first argument
    // read all the arguments into an array
	command := os.Args[1]
	args    := os.Args[2:]
	

	fmt.Println("hello world: ", dir)

    // split dir into components 
}
