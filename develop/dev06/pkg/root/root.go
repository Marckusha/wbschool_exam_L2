package root

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"wbschool_exam_L2/develop/dev06/pkg/utility"

	"github.com/spf13/cobra"
)

//NewCommand create command for grep utility
func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "myGrep",
		Short: "",
		Long:  "",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalf("No such file or string")
				return
			}

			strs, err := ReadLines(args[0])

			if err != nil {
				log.Fatalf("Not found file: %v", err)
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

//SetFlags set flags for sort utility
func SetFlags(c *cobra.Command) {
	c.Flags().StringP("delimiter", "d", "\t", "Delimiter")
	c.Flags().BoolP("separated", "s", false, "Separated")
	c.Flags().IntSliceP("fields", "f", nil, "Fields")
}

//ReadLines return input array string
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
