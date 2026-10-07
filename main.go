package main

import (
	"log"
	"os"

	"github.com/Akene-Uzezi/sprout/internal/util"

	"github.com/spf13/cobra"
)

var (
	moduleName string
	language   string
)

var rootCmd = &cobra.Command{
	Use:     "sprout",
	Short:   "A CLI tool meant to scaffold new Go projects",
	Version: "1.0.12",
	Long: `

  _____________________________ ________   ____ ______________
 /   _____/\______   \______   \\_____  \ |    |   \__    ___/
 \_____  \  |     ___/|       _/ /   |   \|    |   / |    |
 /        \ |    |    |    |   \/    |    \    |  /  |    |
/_______  / |____|    |____|_  /\_______  /______/   |____|
        \/                   \/         \/

This is sprout. A CLI tool meant to scaffold new Go projects.`,
}

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Scaffold new go projects",
	Long:  "This is the command used to create new go projects at the specified path. the '.' character is used for the current directory, any other path should be specified",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		if path == "." {
			wd, err := os.Getwd()
			if err != nil {
				log.Printf("Unable to get wd: %s", err)
				os.Exit(1)
			}
			if moduleName == "" {
				util.NoModuleName(wd)
			} else {
				util.WithModuleName(moduleName)
			}
		} else {
			err := os.Chdir(path)
			if err != nil {
				log.Printf("error: %s", err)
				os.Exit(1)
			}
			wd, err := os.Getwd()
			if err != nil {
				log.Printf("Unable to get wd: %s", err)
				os.Exit(1)
			}
			if moduleName == "" {
				util.NoModuleName(wd)
			} else {
				util.WithModuleName(moduleName)
			}
		}
	},
}

func main() {
	initCmd.Flags().StringVarP(&moduleName, "module", "m", "", "The module name to be created. Defaults to the name of the directory")
	initCmd.Flags().StringVarP(&language, "lang", "l", "", "The language of the project to be scaffolded. Defaults to Go")
	rootCmd.AddCommand(initCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Printf("An Error occured: %s", err)
		os.Exit(1)
	}
}
