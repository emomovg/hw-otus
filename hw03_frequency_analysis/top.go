package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Item struct {
	Key   string
	Value int
}

func Top10(str string) []string {
	sl := strings.Fields(str)
	mp := make(map[string]int)
	for _, s := range sl {
		mp[s]++
	}

	slItems := make([]Item, 0, len(mp))
	for k, v := range mp {
		slItems = append(slItems, Item{Key: k, Value: v})
	}

	sort.Slice(slItems, func(i, j int) bool {
		if slItems[i].Value == slItems[j].Value {
			return slItems[i].Key < slItems[j].Key
		}
		return slItems[i].Value > slItems[j].Value
	})

	resLen := min(10, len(slItems))
	result := make([]string, 0, resLen)
	for _, item := range slItems[:resLen] {
		result = append(result, item.Key)
	}

	return result
}
