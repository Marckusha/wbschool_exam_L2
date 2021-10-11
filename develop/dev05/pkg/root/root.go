package root

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"wbschool_exam_L2/develop/dev05/pkg/grep"
	"wbschool_exam_L2/develop/dev05/pkg/models"

	"github.com/spf13/cobra"
)

//NewCommand create command for grep utility
func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "myGrep",
		Short: "",
		Long:  "",

		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 2 {
				log.Fatalf("No such file or string")
				return
			}

			searchStr := args[0]
			strs, err := ReadLines(args[1])

			if err != nil {
				log.Fatalf("Not found file: %v", err)
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
				gs := grep.NewGrep(config, strs)
				fmt.Println(gs.SearchString(searchStr))
				return
			}

			gs := grep.NewGrep(config, strs)
			gs.SearchString(searchStr)

			fmt.Println(gs.GetStringResult())
		},
	}
}

//SetFlags set flags for sort utility
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
