package main

import kv "github.com/secoba/golang-kv"

func main() {
	packetBf := kv.Badger(".tucache")
	err := packetBf.Set([]byte("test"), []byte("test"))
	if err != nil {
		panic(err)
	}
}
