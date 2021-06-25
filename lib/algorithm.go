package lib

import "sort"

func StringContains(data []string, e string) bool {
	sort.Strings(data)
	i := sort.Search(len(data), func(i int) bool {
		return data[i] == e
	})
	return i != len(data)
}
