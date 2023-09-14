package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var gormCmd = &cobra.Command{
	Use:   "gorm",
	Short: "gorm lib 测试",
	Long:  "gorm lib 测试",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("-------gormCmd-----------")
		gormRun()
	},
}

func gormRun() {

}
