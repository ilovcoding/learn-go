package gktime_getting_started

import (
	"errors"
	"testing"
	"time"
)

type ReusableObject struct {
}
type ObjectPool struct {
	bufChan chan *ReusableObject
}

func NewObjectPool(num int) *ObjectPool {
	objectPool := ObjectPool{}
	objectPool.bufChan = make(chan *ReusableObject, num)
	for i := 0; i < num; i++ {
		objectPool.bufChan <- &ReusableObject{}
	}
	return &objectPool
}

func (p *ObjectPool) GetObject(duration time.Duration) (*ReusableObject, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(duration):
		return nil, errors.New("time out")
	}
}

func (p *ObjectPool) ReleaseObject(object *ReusableObject) error {
	select {
	case p.bufChan <- object:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjectPool(t *testing.T) {
	pool := NewObjectPool(10)
	// 满的对象池多放一个变量会有溢出错误
	if err := pool.ReleaseObject(&ReusableObject{}); err != nil {
		t.Error(err)
	}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObject(time.Second); err != nil {
			t.Error(err)
		} else {
			t.Logf("%p", &v)
			// 使用完了释放对象
			//if err := pool.ReleaseObject(v); err != nil {
			//	t.Error(err)
			//}
		}
	}
}
