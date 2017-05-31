package main

import (
	"github.com/lwhile/utillib/safemap"
	"github.com/lwhile/utillib/strs"
)

func main() {
	m := safemap.NewMap()
	m.Set("key", "value")
	m.Get("key")
	m.Len()
	m.Delete("key")

	strs.WildcardMatch()
}
