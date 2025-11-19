package hash

type Entry struct {
	Key   string
	Value string
}

type Hashmap struct {
	Buckets [][]Entry
	Size    int
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

func AddToHashMap(key string, value string, hashmap *Hashmap, bucketcount int) {
	bucketIndex := Hash(key, bucketcount)
	hashmap.Buckets[bucketIndex] = append(hashmap.Buckets[bucketIndex], Entry{key, value})
}

func CreateHashMap(bucketCount int) *Hashmap {
	if bucketCount <= 0 {
		bucketCount = 16
	}

	return &Hashmap{
		Buckets: make([][]Entry, bucketCount),
		Size:    0,
	}
}
