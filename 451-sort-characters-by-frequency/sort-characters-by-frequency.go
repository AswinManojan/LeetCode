func frequencySort(s string) string {
    result:=""
	mp := make(map[rune]int)
	for _, x := range s {
		mp[x]++
	}
	sortedKeys := sortByValue(mp)
	for _, x := range sortedKeys {
		r := strings.Repeat(string(x), mp[x])
		result+=r
	}
    return result
}

func sortByValue(mp map[rune]int) []rune {
	keys := []rune{}
	for key := range mp {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return mp[keys[j]] < mp[keys[i]]
	})
	return keys
}