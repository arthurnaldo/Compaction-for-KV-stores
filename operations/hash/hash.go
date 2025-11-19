package hash

type entry struct {
	key   int
	value string
}

type hashmap struct {
	buckets [][]entry
	size    int
}

func Hash(key string, bucketCount int) int {
	var h uint64 = 1469598103934665603
	const prime64 = 1099511628211

	for i := 0; i < len(key); i++ {
		h ^= uint64(key[i])
		h *= prime64
	}

	return int(h % uint64(bucketCount))

}

func CreateHashMap(bucketCount int) *hashmap {
	if bucketCount <= 0 {
		bucketCount = 16
	}

	return &hashmap{
		buckets: make([][]entry, bucketCount),
		size:    0,
	}
}
