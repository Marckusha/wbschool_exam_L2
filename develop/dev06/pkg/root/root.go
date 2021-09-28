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
	"wbschool_exam_L2/develop/dev06/pkg/utility"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "myGrep",
		Short: "",
		Long:  "",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("No such file or string")
				return
			}

			strs, err := ReadLines(args[0])

			if err != nil {
				fmt.Println("Not found file")
				return
			}

			config := utility.Config{}

			if ok, _ := cmd.Flags().GetBool("separated"); ok {
				config.IsSeparated = true
			}

			if val, _ := cmd.Flags().GetString("delimiter"); val != "" {
				config.Dilimeter = val
			}

			if val, _ := cmd.Flags().GetIntSlice("fields"); len(val) != 0 {
				config.Fields = val
			}

			util := utility.NewCutUtility(&config, strs)
			result := util.ExecuteUtility()

			fmt.Println(len(result))
			fmt.Println(result)

		},
	}
}

func SetFlags(c *cobra.Command) {
	c.Flags().StringP("delimiter", "d", "\t", "Delimiter")
	c.Flags().BoolP("separated", "s", false, "Separated")
	c.Flags().IntSliceP("fields", "f", nil, "Fields")
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
