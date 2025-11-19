package get

import (
	"fmt"
	"operations/hash"
	"operations/state"
)

func Get(key string) (string, error) {
	index := hash.Hash(key, state.KeyDict.Size)
	keydict := state.KeyDict
	for _, entry := range keydict.Buckets[index] {
		if entry.Key == key {
			return entry.Value, nil
		}
	}
	return "", fmt.Errorf("key doesn't exist")
}
