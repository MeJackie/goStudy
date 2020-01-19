package distributedSystem

import (
	"fmt"
	"sync"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}

	return l
}

func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}

	return lockResult
}

func (l Lock) Unlock() {
	l.c <- struct{}{}
}

func Run() {
	var counter int
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if !l.Lock() {
				fmt.Println("lock failed")
				return
			}
			counter++
			fmt.Println("current counter", counter)
			l.Unlock()
		}()
	}

	wg.Wait()
}


// zookeeper
func TryLock()  {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)
	}
	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	l.Unlock()
}


