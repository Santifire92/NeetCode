package arraysandhashing

import "sort"

func groupAnagrams(strs []string) [][]string {
	hashingTable := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		if anagramsBag, ok := hashingTable[sortedStr]; ok {
			hashingTable[sortedStr] = append(anagramsBag, str)
		} else {
			hashingTable[sortedStr] = []string{str}
		}
	}
	result := make([][]string, 0, len(hashingTable))
	for _, anagramsBag := range hashingTable {
		result = append(result, anagramsBag)
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}
