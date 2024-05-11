func isAnagram(s string, t string) bool {
	sMap := map[rune]int{}
	tMap := map[rune]int{}

	for _, value := range s {
		sMap[value] += 1
	}

	for _, value := range t {
		tMap[value] += 1
	}

	return reflect.DeepEqual(sMap, tMap)
}