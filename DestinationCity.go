package leetcode

func destCity(paths [][]string) string {
	var hash = make(map[string]string)

	for _, path := range paths {
		hash[path[0]] = path[1]
	}

	var start = paths[0][0]

	for hash[start] != "" {
		start = hash[start]
	}

	return start
}
