package safemap

// SafeMap interface
type SafeMap interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	GetAll() map[string]interface{}
	Delete(key string)
	Len() int
}

type mapPair struct {
	key   string
	value interface{}
}

type valuePair struct {
	value interface{}
	exist bool
}

type mapChan struct {
	m          map[string]interface{}
	setReqChan chan *mapPair
	setRepChan chan interface{}

	getReqChan chan string
	getRepChan chan *valuePair

	getAllReqChan chan struct{}
	getAllRepChan chan *mapPair

	delReqChan chan string
	delRepChan chan interface{}
}

func (c *mapChan) Set(key string, value interface{}) {
	c.setReqChan <- &mapPair{key, value}
	<-c.setRepChan
}

func (c *mapChan) Get(key string) (interface{}, bool) {
	c.getReqChan <- key
	vp := <-c.getRepChan
	return vp.value, vp.exist
}

func (c *mapChan) GetAll() map[string]interface{} {
	m := make(map[string]interface{})
	c.getAllReqChan <- struct{}{}
	done := make(chan struct{})
	var count int
	go func() {
		for {
			select {
			case vp := <-c.getAllRepChan:
				m[vp.key] = vp.value
				count++
				if count == c.Len() {
					close(done)
				}
			}
		}
	}()
	<-done
	return m
}

func (c *mapChan) Delete(key string) {
	c.delReqChan <- key
	<-c.delRepChan
}

func (c *mapChan) Len() int {
	return len(c.m)
}

// NewMap return a iMap instance
func NewMap() SafeMap {
	c := mapChan{
		m:             make(map[string]interface{}),
		setReqChan:    make(chan *mapPair),
		setRepChan:    make(chan interface{}),
		getReqChan:    make(chan string),
		getRepChan:    make(chan *valuePair),
		getAllReqChan: make(chan struct{}),
		getAllRepChan: make(chan *mapPair),
		delReqChan:    make(chan string),
		delRepChan:    make(chan interface{}),
	}

	go func() {
		for {
			select {
			case r := <-c.setReqChan:
				c.m[r.key] = r.value
				c.setRepChan <- nil
			case k := <-c.getReqChan:
				if v, exist := c.m[k]; exist {
					c.getRepChan <- &valuePair{v, true}
				} else {
					c.getRepChan <- &valuePair{nil, false}
				}
			case <-c.getAllReqChan:
				for k, v := range c.m {
					c.getAllRepChan <- &mapPair{k, v}
				}
			case k := <-c.delReqChan:
				delete(c.m, k)
				c.delRepChan <- nil
			}
		}
	}()
	return &c
}
