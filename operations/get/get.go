package get

import "os"

func Get(key string) string {
	file, err := os.ReadFile("/hashmap/hashmap.yaml")
	if err != nil {
		panic(err)
	}
	println(file)
	return "memes"
}
