package main

import (
	"fmt"

	"strconv"

	"github.com/lwhile/utillib/safemap"
)

func main() {
	m := safemap.NewMap()
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			m.Push(strconv.Itoa(i), i)
			fmt.Println(m.Pop(strconv.Itoa(i)))
		}(i)
	}
	for {
	}
}
