package utility

import (
	"sort"
	"strings"
)

type Config struct {
	Dilimeter   string
	IsSeparated bool
	Fields      []int
}

type CutUtility struct {
	config Config
	strs   []string
}

func (cu *CutUtility) ExecuteUtility() []string {
	res := make([]string, 0)
	sort.Slice(cu.config.Fields, func(i, j int) bool {
		return cu.config.Fields[i] < cu.config.Fields[j]
	})

	for i := 0; i < len(cu.strs); i++ {
		if cu.config.IsSeparated && !strings.Contains(cu.strs[i], cu.config.Dilimeter) {
			continue
		}

		splitStr := strings.Split(cu.strs[i], cu.config.Dilimeter)

		var newstr []string
		for _, numColumn := range cu.config.Fields {
			numColumn--
			if numColumn < len(splitStr) && numColumn >= 0 {
				newstr = append(newstr, splitStr[numColumn])
			}
		}

		res = append(res, strings.Join(newstr, cu.config.Dilimeter))
	}

	return res
}

func NewCutUtility(conf *Config, s []string) *CutUtility {
	return &CutUtility{
		config: *conf,
		strs:   s,
	}
}
