package groupanagram

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mS := [26]rune{}
	for _, value := range s {
		mS[value-'a'] += 1
	}
	for _, value := range t {
		mS[value-'a'] -= 1
		if mS[value-'a'] < 0 {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	groupAnagramString := [][]string{}
	groupMap := make(map[string][]string)
	for _, value := range strs {
		isPresent := false
		for indexValue, _ := range groupMap {
			if isAnagram(indexValue, value) {
				groupMap[indexValue] = append(groupMap[indexValue], value)
				isPresent = true
			}
		}
		if !isPresent {
			groupMap[value] = append(groupMap[value], value)
		}

	}
	for indexValue, _ := range groupMap {
		groupAnagramString = append(groupAnagramString, groupMap[indexValue])
	}

	return groupAnagramString
}

func TestGroupAnagrams() {
	strS := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strS))
}
