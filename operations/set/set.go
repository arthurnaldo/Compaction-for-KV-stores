package set

import (
	"operations/hash"
	"operations/state"
	"operations/structs"
)

func Set(key string, value string, delete bool) bool {
	if delete {
		value = "TOMBSTONE"
	}
	mostRecentSegment := &(*state.Segments)[len(*state.Segments)-1]
	*mostRecentSegment = append(*mostRecentSegment, structs.Segment{Key: key, Value: value})

	mostRecentHashmap := &(*state.Hashmaps)[len(*state.Hashmaps)-1]
	hash.AddToHashMap(key, value, mostRecentHashmap)

	if len(*mostRecentSegment) > 100 {
		*state.Segments = append(*state.Segments, []structs.Segment{})
		*state.Hashmaps = append(*state.Hashmaps, structs.Hashmap{})
	}

	return true
}
