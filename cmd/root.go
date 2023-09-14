package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golibs",
	Short: "go常用库",
	Long:  "go常用库",
}

func init() {
	log.Println("-----init------")
	rootCmd.AddCommand(gormCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
