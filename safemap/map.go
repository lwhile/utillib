package safemap

type iMap interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
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
	m      map[string]interface{}
	setReq chan *mapPair
	setRep chan interface{}

	getReq chan string
	getRep chan *valuePair

	delReq chan string
	delRep chan interface{}
}

func (c *mapChan) Set(key string, value interface{}) {
	c.setReq <- &mapPair{key, value}
	<-c.setRep
}

func (c *mapChan) Get(key string) (interface{}, bool) {
	c.getReq <- key
	vp := <-c.getRep
	return vp.value, vp.exist
}

func (c *mapChan) Delete(key string) {
	c.delReq <- key
	<-c.delRep
}

func (c *mapChan) Len() int {
	return len(c.m)
}

// NewMap return a iMap instance
func NewMap() iMap {
	c := mapChan{
		m:      make(map[string]interface{}),
		setReq: make(chan *mapPair),
		setRep: make(chan interface{}),
		getReq: make(chan string),
		getRep: make(chan *valuePair),
		delReq: make(chan string),
		delRep: make(chan interface{}),
	}

	go func() {
		for {
			select {
			case r := <-c.setReq:
				c.m[r.key] = r.value
				c.setRep <- nil
			case k := <-c.getReq:
				if v, exist := c.m[k]; exist {
					c.getRep <- &valuePair{v, true}
				} else {
					c.getRep <- &valuePair{nil, false}
				}
			case k := <-c.delReq:
				delete(c.m, k)
				c.delRep <- nil
			}
		}
	}()
	return &c
}
