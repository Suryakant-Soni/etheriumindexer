package main

import (
	"etheriumindexer/internal/indexer"
	"etheriumindexer/internal/setup"
)

func init() {
	setup.Setup()
}
func main() {
	indexer.Run()
}
