// Package util for the basic utility functions this will be using
package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func NoModuleName(wd string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to get home dir: %s", err)
		os.Exit(1)
	}
	trimdir := strings.TrimLeft(wd, homeDir)
	splittrimdir := strings.SplitN(trimdir, "/", 2)
	moduleName := splittrimdir[1]
	cmd := exec.Command("go", "mod", "init", moduleName)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("error executing go mod init: %s", err)
		os.Exit(1)
	}
	CreateDirAndFiles()
	log.Println(string(output))
}

func WithModuleName(moduleName string) {
	cmd := exec.Command("go", "mod", "init", moduleName)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("error executing go mod init: %s", err)
		os.Exit(1)
	}
	CreateDirAndFiles()
	log.Printf(string(output))
}

func CreateDirAndFiles() {
	err := os.Mkdir("cmd", 0o777)
	if err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
	err = os.Mkdir("internal", 0o777)
	if err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
	err = os.Chdir("cmd")
	if err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
	mainFileInput := `
		package main
		
		import "fmt"
		
		func main() {
			fmt.Println("hello world")
		}
	`
	err = os.WriteFile("main.go", []byte(mainFileInput), 0o666)
	if err != nil {
		log.Printf("error writing to file: %s", err)
		os.Exit(1)
	}
}
