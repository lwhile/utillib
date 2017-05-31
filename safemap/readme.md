# SafeMap

safemap包实现了一个并发安全的map类型,相比常规的加锁方案,包里使用另外一种思路,即使用channel代替mutex.
经过测试,使用channel代替mutex可以带来3-4倍的性能提升.


## Benchmark

> cd safemap && go test -run= map_test.go map.go map1.go -bench=.


|                  | loop  |  ns/op |
|------------------|-------|--------|
|BenchmarkSafeMap-4|200000 |9623    |
|BenchmarkLockMap-4|1000000|3434    |


## Example

    import 	"github.com/lwhile/utillib/safemap"

    // setter,getter,delete,len这四类操作的表现特性与内置map类型一致
	
    // 注意Len方法没有实现并发安全

	m := safemap.NewMap()

	m.Set("key", "value") 

	value, exist := m.Get("key") // 第一个返回数据的类型为接口类型.若键值不存在则返回 nil,false

	m.Len() 	             // 返回map的大小,返回类型为int

	m.Delete("key")