package main

import (
	"github.com/lwhile/utillib/safemap"
)

func main() {
	m := safemap.NewMap()
	m.Set("key", "value")
	m.Get("key")
	m.Len()
	m.Delete("key")
}
