/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package root

import (
	"bufio"
	"fmt"
	"os"
	"wbschool_exam_L2/develop/dev05/pkg/grep"
	"wbschool_exam_L2/develop/dev05/pkg/models"

	"github.com/spf13/cobra"
)

//var cfgFile string

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "myGrep",
		Short: "",
		Long:  "",

		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 2 {
				fmt.Println("No such file or string")
				return
			}

			searchStr := args[0]
			strs, err := ReadLines(args[1])

			if err != nil {
				fmt.Println("Not found file")
				return
			}

			config := models.Configs{}

			if val, _ := cmd.Flags().GetInt("context"); val > 0 {
				config.FormatPos = val
				config.FormatOut = models.C
			}
			if val, _ := cmd.Flags().GetInt("after"); val > 0 {
				config.FormatPos = val
				config.FormatOut = models.A
			}
			if val, _ := cmd.Flags().GetInt("before"); val > 0 {
				config.FormatPos = val
				config.FormatOut = models.B
			}

			if ok, _ := cmd.Flags().GetBool("ignore"); ok {
				config.IsIgnore = true
			}
			if ok, _ := cmd.Flags().GetBool("invert"); ok {
				config.IsInvert = true
			}
			if ok, _ := cmd.Flags().GetBool("fixed"); ok {
				config.IsFixed = true
			}
			if ok, _ := cmd.Flags().GetBool("num"); ok {
				config.IsNum = true
			}
			if ok, _ := cmd.Flags().GetBool("count"); ok {
				gs := grep.NewGrepString(config, strs)
				fmt.Println(gs.SearchString(searchStr))
				return
			}

			gs := grep.NewGrepString(config, strs)
			gs.SearchString(searchStr)

			fmt.Println(gs.GetStringResult())
		},
	}
}

func SetFlags(c *cobra.Command) {
	var (
		flagA int
		flagB int
		flagC int
	)

	c.Flags().IntVarP(&flagA, "after", "A", 0, "Print +N after string match")
	c.Flags().IntVarP(&flagB, "before", "B", 0, "Before")
	c.Flags().IntVarP(&flagC, "context", "C", 0, "Context")
	c.Flags().BoolP("count", "c", false, "Count")
	c.Flags().BoolP("ignore", "i", false, "Ignore-case")
	c.Flags().BoolP("invert", "v", false, "Invert")
	c.Flags().BoolP("fixed", "F", false, "Fixed")
	c.Flags().BoolP("num", "n", false, "Line num")
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sLines = append(sLines, scanner.Text())
	}

	return sLines, scanner.Err()
}
