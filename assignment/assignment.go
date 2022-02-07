package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	i := false
	sum64 := uint64(x) + uint64(y)
	sum := uint32(sum64)
	carryOut := uint32(sum64 >> 32)
	if carryOut == 1 {
		i = true
	}
	return sum, i
}

func CeilNumber(f float64) float64 {
	//res2 := float64(int64(f/0.05+0.25)) * 0.05
	res := math.Ceil(f/0.25) * 0.25
	return res

}

func AlphabetSoup(word string) string {

	s := []rune(word)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func StringMask(s string, n uint) string {
	rs := []rune(s)
	if len(s) <= 0 {
		return "*"
	} else if len(s) <= int(n) {
		for i := 0; i < len(rs); i++ {
			rs[i] = '*'
		}
		return string(rs)
	}
	for i := int(n); i < len(rs); i++ {
		rs[i] = '*'
	}
	return string(rs)
}

func WordSplit(arr [2]string) string {
	count := 0
	words := strings.Split(arr[1], ",")
	founded := ""
	s := arr[0]
	for i := 0; i < len(words); i++ {
		t := strings.Contains(s, words[i])
		if t {
			founded += words[i] + ","
			count++
		}
	}
	if count > 1 {
		return founded
	} else {
		return "not possible"
	}

}

func VariadicSet(i ...interface{}) []interface{} {
	keys := make(map[interface{}]struct{})
	list := []interface{}{}

	for _, v := range i {
		if _, value := keys[v]; !value {
			keys[v] = struct{}{}
			list = append(list, v)
		}
	}
	return list
}
