package root

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"wbschool_exam_L2/develop/dev03/pkg/ArrayStrings"

	"github.com/spf13/cobra"
)

var cfgFile string

//var LinesStr ArrayStrings.ArrayStrings
var TestString string

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "mySort",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:

		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				fmt.Println("No such file")
				return
			}

			str, err := readLines(args[0])

			if err != nil {
				fmt.Println("Not found file")
				return
			}

			LinesStr := ArrayStrings.NewArrayStrings(str)

			//если флаг не выбран, то сортируем стандартной сортировкой
			f := LinesStr.StandartSort

			if val, _ := cmd.Flags().GetInt("column"); val > 0 {
				LinesStr.SetSortColumn(val)
			}

			if ok, _ := cmd.Flags().GetBool("unique"); ok {
				LinesStr.Unique()
				f = LinesStr.StandartSort
			}

			if ok, _ := cmd.Flags().GetBool("ignore"); ok {
				LinesStr.IgnoreSpace()
			}

			if ok, _ := cmd.Flags().GetBool("number"); ok {
				f = LinesStr.NumberSort
			}

			if ok, _ := cmd.Flags().GetBool("month"); ok {
				f = LinesStr.MonthSort
			}

			if ok, _ := cmd.Flags().GetBool("suffix"); ok {
				//todo
				fmt.Println("is suffix")
			}

			//sorts
			sort.SliceStable(LinesStr, f)

			if ok, _ := cmd.Flags().GetBool("reverse"); ok {
				LinesStr.Reverse()
			}

			if ok, _ := cmd.Flags().GetBool("check"); ok {

				cpSlice := ArrayStrings.NewArrayStrings(str)

				fmt.Println(LinesStr.Equal(cpSlice))

				return
			}

			var b strings.Builder
			for i := 0; i < len(LinesStr); i++ {
				b.WriteString(LinesStr[i].Value)
				if i == len(LinesStr)-1 {
					break
				}
				b.WriteString("\r\n")
			}
			TestString = b.String()
		},
	}
}

func SetFlags(c *cobra.Command) {
	var count int

	c.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mySort.yaml)")

	c.Flags().IntVarP(&count, "column", "k", 0, "Sorts by column")
	c.Flags().BoolP("reverse", "r", false, "Revers sorts")
	c.Flags().BoolP("number", "n", false, "Sort by number")
	c.Flags().BoolP("unique", "u", false, "Unique values sort")
	c.Flags().BoolP("month", "M", false, "Sort month")
	c.Flags().BoolP("ignore", "b", false, "Ignore tailing space")
	c.Flags().BoolP("check", "c", false, "Check sort")

	//TODO
	c.Flags().BoolP("suffix", "H", false, "Check suffix")
}

func readLines(path string) ([]string, error) {
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
