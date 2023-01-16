/*
Copyright © 2023 haormj <haormj@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init dir and other file",
	Run: func(cmd *cobra.Command, args []string) {
		// 创建values.yaml文件和templates目录，其他的应该也没有什么
		if _, err := os.Create("./values.yaml"); err != nil {
			log.Fatalln(err)
		}
		if err := os.Mkdir("./templates", 0755); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
