package main

func createHugeString() {}

func main() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}
