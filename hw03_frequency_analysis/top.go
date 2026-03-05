package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

type Item struct {
	Key   string
	Value int
}

func Top10(str string) []string {
	sl := strings.Fields(str)
	mp := make(map[string]int)
	for _, s := range sl {
		clean := cleanWord(s)
		if clean == "-" {
			continue
		}
		mp[clean]++
	}

	slItems := make([]Item, 0, len(mp))
	isExists := make(map[string]struct{})
	for _, s := range sl {
		cl := cleanWord(s)
		if k, ok := mp[cl]; ok {
			if _, exists := isExists[cl]; exists {
				continue
			}
			isExists[cl] = struct{}{}
			slItems = append(slItems, Item{Key: cl, Value: k})
		}

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

func cleanWord(s string) string {
	trimmed := strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-'
	})
	return strings.ToLower(trimmed)
}
