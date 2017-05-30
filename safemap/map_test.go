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
			if vv, ok := v.(int); ok {
				if vv != i {
					fmt.Println(vv, "!=", i)
					t.Fail()
				}
			} else {
				fmt.Println("assert v to int fail")
				t.Fail()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestLockMap(t *testing.T) {
	m := NewLockMap()
	size := 10000
	wg := sync.WaitGroup{}
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			m.Push(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			v := m.Pop(strconv.Itoa(i))
			if vv, ok := v.(int); ok {
				if vv != i {
					fmt.Println(vv, "!=", i)
					t.Fail()
				}
			} else {
				fmt.Println("assert v to int fail")
				t.Fail()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkSafeMap(b *testing.B) {
	m := NewMap()
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			m.Push(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			m.Pop(strconv.Itoa(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkLockMap(b *testing.B) {
	m := NewLockMap()
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			m.Push(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			m.Pop(strconv.Itoa(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
