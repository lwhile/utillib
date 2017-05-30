package safemap

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestSafeMap(t *testing.T) {
	m := NewMap()
	size := 100000
	wg := sync.WaitGroup{}
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			m.Push(strconv.Itoa(i), i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			v := m.Pop(strconv.Itoa(i))
			if v != strconv.Itoa(i) {
				if vv, ok := v.(int); ok {
					if vv != i {
						fmt.Println(vv, "!=", i)
						t.Fail()
					}
				} else {
					fmt.Println("assert v to int fail")
					t.Fail()
				}
			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
