package safemap

import (
	"sync"
)

type lockMap struct {
	sync.Mutex
	m map[string]interface{}
}

func (lm *lockMap) Set(key string, value interface{}) {
	lm.Lock()
	defer lm.Unlock()
	lm.m[key] = value
}

func (lm *lockMap) Get(key string) (interface{}, bool) {
	lm.Lock()
	defer lm.Unlock()
	v, exist := lm.m[key]
	return v, exist
}

func (lm *lockMap) Delete(key string) {
	lm.Lock()
	defer lm.Unlock()
	delete(lm.m, key)
}

func (lm *lockMap) Len() int {
	return len(lm.m)
}

func newLockMap() iMap {
	lm := lockMap{}
	lm.m = make(map[string]interface{})
	return &lm
}
