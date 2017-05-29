package safemap

// 参考自云风的blog:
// http://blog.codingnow.com/2011/03/go_2.html
// 原作者拥有这段代码的一切权力.

type iMap interface {
	Push(key string, e interface{}) interface{}
	Pop(key string) interface{}
}

type mapPair struct {
	key   string
	value interface{}
}

type mapChan struct {
	pushReq chan *mapPair
	pushRep chan interface{}
	popReq  chan string
	popRep  chan interface{}
}

func (c *mapChan) Push(key string, value interface{}) interface{} {
	c.pushReq <- &mapPair{key, value}
	return <-c.pushRep
}

func (c *mapChan) Pop(key string) interface{} {
	c.popReq <- key
	return <-c.popRep
}

// NewMap return a iMap instance
func NewMap() iMap {
	c := mapChan{
		pushReq: make(chan *mapPair),
		pushRep: make(chan interface{}),
		popReq:  make(chan string),
		popRep:  make(chan interface{}),
	}
	m := make(map[string]interface{})
	go func() {
		for {
			select {
			case r := <-c.pushReq:
				if v, exist := m[r.key]; exist {
					c.pushRep <- v
				} else {
					m[r.key] = r.value
					c.pushRep <- nil
				}
			case k := <-c.popReq:
				if v, exist := m[k]; exist {
					delete(m, k)
					c.popRep <- v
				} else {
					c.popRep <- nil
				}
			}
		}
	}()
	return &c
}
