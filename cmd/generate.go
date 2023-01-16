/*
Copyright © 2023 haormj <haormj@gmail.com>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "use templates and values generate target",
	Run: func(cmd *cobra.Command, args []string) {
		currentDir := "."
		currentDirAbs, err := filepath.Abs(currentDir)
		if err != nil {
			log.Fatalln(err)
		}
		// 解析values文件，并放到map数组中，后续作为template的变量放进去
		data := make(map[string]any)
		valuesFilePath := filepath.Join(currentDirAbs, "values.yaml")
		valuesFileBytes, err := os.ReadFile(valuesFilePath)
		if err != nil {
			log.Fatalln(err)
		}
		if err := yaml.Unmarshal(valuesFileBytes, data); err != nil {
			log.Fatalln(err)
		}
		templateDir := filepath.Join(currentDirAbs, "templates")
		targetDir := filepath.Join(currentDirAbs, "target")
		// 每次重新生成前，直接将之前的target目录删除，然后重新创建
		if err := os.RemoveAll(targetDir); err != nil {
			log.Fatalln(err)
		}
		if err := os.Mkdir(targetDir, 0755); err != nil {
			log.Fatalln(err)
		}
		// 遍历template目录，只便利该目录下的文件
		dirEntrys, err := os.ReadDir(templateDir)
		if err != nil {
			log.Fatalln(err)
		}
		for _, dirEntry := range dirEntrys {
			// do nothing
			if dirEntry.IsDir() {
				continue
			}
			templateFilePath := filepath.Join(templateDir, dirEntry.Name())
			targetFilePath := filepath.Join(targetDir, dirEntry.Name())
			// 读取template文件内容
			templateFileBytes, err := os.ReadFile(templateFilePath)
			if err != nil {
				log.Fatalln(err)
			}
			// 创建golang template，添加template函数库sprig
			t, err := template.New("code generator").Funcs(sprig.FuncMap()).Parse(string(templateFileBytes))
			if err != nil {
				log.Fatalln(err)
			}
			f, err := os.OpenFile(targetFilePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
			if err != nil {
				log.Fatalln(err)
			}
			// 替换模板中的内容，然后写到target目录中
			if err := t.Execute(f, data); err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
