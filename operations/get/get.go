package get

import (
	"operations/hash"
	"operations/state"
	"operations/structs"
	"slices"
)

func getIndex(key string) (value int, seg *[]structs.Segment) {
	for i, hashmap := range slices.Backward(*state.Hashmaps) {
		idx := hash.GetIndex(key, &hashmap)
		if idx != -1 {
			return idx, &(*state.Segments)[len(*state.Hashmaps)-1-i]
		}
	}
	return -1, nil
}

func Get(key string) (value string) {
	index, indices := getIndex(key)
	return (*indices)[index].Value
}
