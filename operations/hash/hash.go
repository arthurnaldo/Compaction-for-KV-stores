package hash

import "operations/structs"

func Hash(key string, bucketCount int) int {
	var h uint64 = 1469598103934665603
	const prime64 = 1099511628211

	for i := 0; i < len(key); i++ {
		h ^= uint64(key[i])
		h *= prime64
	}

	return int(h % uint64(bucketCount))

}

func GetIndex(key string, hashmap *structs.Hashmap) int {
	idx := Hash(key, hashmap.Size)
	for i, entry := range hashmap.Buckets[idx] {
		if entry.Key == key {
			return i
		}
	}
	return -1
}

func AddToHashMap(key string, value string, hashmap *structs.Hashmap) {
	bucketIndex := Hash(key, hashmap.Size)
	bucket := hashmap.Buckets[bucketIndex]
	for i := range bucket {
		if bucket[i].Key == key {
			bucket[i].Value = value
			return
		}
	}
	hashmap.Buckets[bucketIndex] = append(hashmap.Buckets[bucketIndex], structs.Entry{Key: key, Value: value})
}

func DeleteFromHashMap(key string, hashmap *structs.Hashmap) {

}

// func CreateHashMap(bucketCount int) *Hashmap {
// 	if bucketCount <= 0 {
// 		bucketCount = 16
// 	}

// 	return &Hashmap{
// 		Buckets: make([][]Entry, bucketCount),
// 		Size:    0,
// 	}
// }
