package list

import (
	"fmt"
)

type List struct {
	list []interface{} // 栈上
}

func New() *List {
	return &List{list: []interface{}{}}
}

func (l *List) Len() int {
	return len(l.list)
}

func (l *List) LAppend(node interface{}) {
	l.list = append([]interface{}{node}, l.list...)
}

func (l *List) RAppend(node interface{}) {
	l.list = append(l.list, node)
}

func (l *List) LPop() (node interface{}, err error) {
	if l.list == nil || l.Len() == 0 {
		return nil, fmt.Errorf("nil")
	}

	node = l.list[0]
	if len(l.list) >= 2 {
		l.list = l.list[1:]
	} else {
		l.list = l.list[0:0]
	}
	return node, nil
}

func (l *List) RPop() (node interface{}, err error) {
	if l.list == nil || l.Len() == 0 {
		return nil, fmt.Errorf("nil")
	}

	node = l.list[len(l.list)-1]
	if len(l.list) >= 2 {
		l.list = l.list[:len(l.list)-1]
	} else {
		l.list = l.list[0:0]
	}
	return node, nil
}

func (l *List) DelByIdx(idx int) {
	i := len(l.list)

	if i <= idx {
		return
	}

	if i == 0 {
		return
	}

	if i == 1 {
		if idx == 0 {
			l.list = l.list[0:0]
			return
		}
	}

	if idx == 0 {
		l.list = append(l.list[1:])
		return
	}
	l.list = append(l.list[:idx], l.list[idx+1:]...)
}

func (l *List) DelByString(key string) bool {
	if len(l.list) == 0 {
		return false
	}

	for k, v := range l.list {
		if v.(string) == key {
			if k == 0 {
				if len(l.list) == 1 {
					l.list = l.list[0:0]
					return true
				} else {
					l.list = l.list[1:]
					return true
				}
			}

			l.list = append(l.list[:k], l.list[k+1:]...)
			return true
		}
	}

	return false
}

func (l *List) MoveToFront(key string) {
	for i, v := range l.list {
		if v.(string) == key {
			l.DelByIdx(i)
			l.RAppend(key)
			break
		}
	}
}
