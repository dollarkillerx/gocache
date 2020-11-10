package lru

import (
	"fmt"
	"log"
	"sync"

	"github.com/dollarkillerx/gocache/list"
)

// LRU 数据idx 存储 双向队列中O(1), 数据存储在MAP中O(1)
type LRU struct {
	maxTotal int64
	total    int64 //  slice 可用存储大小 1*100

	list  *list.List
	cache map[string]interface{}
	mu sync.Mutex
}

func New(maxTotal int64) *LRU {
	if maxTotal == 0 {
		maxTotal = 2000
	} else {
		maxTotal = maxTotal * 200
	}

	return &LRU{
		maxTotal: maxTotal,
		list:     list.New(),
		cache:    map[string]interface{}{},
	}
}

func (l *LRU) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.list.Len()
}

func (l *LRU) lenSync() int {
	return l.list.Len()
}

func (l *LRU) Get(key string) (val interface{}, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if ele, ok := l.cache[key]; ok {
		// 更新LRU队列
		l.list.MoveToFront(key)
		return ele, nil
	}
	return nil, fmt.Errorf("404")
}

func (l *LRU) Set(key string, value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.cache[key]; ok {
		// 更新LRU队列
		l.list.MoveToFront(key)
		l.cache[key] = value
		return
	}
	if l.total < l.maxTotal {
		l.total++
		l.list.RAppend(key)
		l.cache[key] = value
		return
	}
	// 开始清缓存
	pop, err := l.list.LPop()
	if err != nil {
		log.Println(err)
		return
	}
	delete(l.cache, pop.(string))

	l.list.RAppend(key)
	l.cache[key] = value
}

func (l *LRU) Del(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.lenSync() == 0 {
		return
	}

	if l.list.DelByString(key) {
		l.total--
		delete(l.cache, key)
	}
}
