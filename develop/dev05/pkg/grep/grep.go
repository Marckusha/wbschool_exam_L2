package grep

import (
	"strconv"
	"strings"
	"wbschool_exam_L2/develop/dev05/pkg/models"
)

//Grep struct for grep utility
//strs - consist input arr strings
//indexes - consist indexes strs to input
//config - config for grep
type Grep struct {
	strs    []string
	indexes []int
	config  models.Configs
}

//SearchString main func; sets indexes
func (gs *Grep) SearchString(substr string) int {

	var (
		eqFunc func(s1, s2 string) bool
	)

	if gs.config.IsIgnore {
		substr = strings.ToLower(substr)
	}

	if gs.config.IsFixed {
		eqFunc = func(s1, s2 string) bool {
			return strings.Compare(s1, s2) == 0
		}
	} else {
		eqFunc = strings.Contains
	}

	//эквиваленция
	//если мы инвентируем результат, то IsInvert false и eqFunc должен вернуть false, для верного результата
	for i, elem := range gs.strs {
		if gs.config.IsIgnore {
			elem = strings.ToLower(elem)
		}
		if eqFunc(elem, substr) == !gs.config.IsInvert {
			gs.indexes = append(gs.indexes, i)
		}
	}

	return len(gs.indexes)
}

//GetStringResult build string based on indexes and FormatOUT
func (gs Grep) GetStringResult() []string {
	var (
		resStrs    []string
		nums       []int
		deltaLeft  int
		deltaRight int
	)

	if gs.config.FormatOut == models.A || gs.config.FormatOut == models.C {
		deltaRight = gs.config.FormatPos
	}
	if gs.config.FormatOut == models.B || gs.config.FormatOut == models.C {
		deltaLeft = gs.config.FormatPos
	}

	for i := 0; i < len(gs.indexes); i++ {
		left := gs.indexes[i] - deltaLeft
		right := gs.indexes[i] + deltaRight

		if left < 0 {
			left = 0
		}
		if right >= len(gs.strs) {
			right = len(gs.strs) - 1
		}

		if len(nums) != 0 && right <= nums[len(nums)-1] {
			continue
		} else if len(nums) != 0 && left <= nums[len(nums)-1] {
			left = nums[len(nums)-1] + 1
		}

		for j := left; j <= right; j++ {
			nums = append(nums, j)
		}
	}

	for i := 0; i < len(nums); i++ {
		if gs.config.IsNum {
			resStrs = append(resStrs, strconv.Itoa(nums[i]+1)+":"+gs.strs[nums[i]])
		} else {
			resStrs = append(resStrs, gs.strs[nums[i]])
		}
	}

	return resStrs
}

//NewGrep ...
func NewGrep(c models.Configs, s []string) *Grep {
	return &Grep{
		strs:   s,
		config: c,
	}
}
