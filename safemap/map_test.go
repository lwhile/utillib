package safemap

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestSafeMap(t *testing.T) {
	m := NewMap()
	size := 10000
	wg := sync.WaitGroup{}
	// test set item to map
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			m.Set(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	// test length of map
	if m.Len() != size {
		t.Errorf("%d != %d\n", m.Len(), size)
	}

	// test get item of map
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			v, ok := m.Get(strconv.Itoa(i))
			if !ok {
				t.Errorf("No exist key %d\n", i)
			}
			if vv, ok := v.(int); ok {
				if vv != i {
					t.Errorf("%d != %d\n", vv, i)
				}
			} else {
				t.Errorf("Assert v to int fail")
			}
			wg.Done()
		}(i)
	}

	// test get item which no exist
	_, ok := m.Get("test")
	if ok {
		t.Fail()
	}

	// test delete item from map
	for i := 0; i < size/2; i++ {
		wg.Add(1)
		go func(i int) {
			m.Delete(strconv.Itoa(i))
			wg.Done()
		}(i)
	}
	wg.Wait()

	// test length of map after deleteing map
	if m.Len() != size/2 {
		t.Errorf("%d != %d\n", m.Len(), size/2)
	}
}

func TestLockMap(t *testing.T) {
	m := newLockMap()
	size := 10000
	wg := sync.WaitGroup{}
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			m.Set(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			v, ok := m.Get(strconv.Itoa(i))
			if !ok {
				t.Errorf("No exist key %d\n", strconv.Itoa(i))
			} else {
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
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {

		}(i)
	}
}

func BenchmarkSafeMap(b *testing.B) {
	m := NewMap()
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			m.Set(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			m.Get(strconv.Itoa(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkLockMap(b *testing.B) {
	m := newLockMap()
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			m.Set(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			m.Get(strconv.Itoa(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
