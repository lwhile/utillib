package safemap

import (
	"sync"
)

type lockMap struct {
	sync.Mutex
	m map[string]interface{}
}

func (lm *lockMap) Push(key string, value interface{}) interface{} {
	lm.Lock()
	defer lm.Unlock()
	if v, exist := lm.m[key]; exist {
		return v
	}
	lm.m[key] = value
	return nil
}

func (lm *lockMap) Pop(key string) interface{} {
	lm.Lock()
	defer lm.Unlock()
	if v, exist := lm.m[key]; exist {
		delete(lm.m, key)
		return v
	}
	return nil
}

func NewLockMap() *lockMap {
	lm := lockMap{}
	lm.m = make(map[string]interface{})
	return &lm
}
