package structs

type Segment struct {
	Key   string
	Value string
}

type Index struct {
	Key   string
	Value int
}

//for our hashmap impl.

type Entry struct {
	Key   string
	Value string
}

type Hashmap struct {
	Buckets [][]Entry
	Size    int
}
