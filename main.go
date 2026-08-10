package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sprout",
	Short: "A CLI tool meant to scaffold new Go projects",
	Long: `

  _____________________________ ________   ____ ______________
 /   _____/\______   \______   \\_____  \ |    |   \__    ___/
 \_____  \  |     ___/|       _/ /   |   \|    |   / |    |
 /        \ |    |    |    |   \/    |    \    |  /  |    |
/_______  / |____|    |____|_  /\_______  /______/   |____|
        \/                   \/         \/

This is sprout. A CLI tool meant to scaffold new Go projects.`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("An Error occured: %s", err)
		os.Exit(1)
	}
}
