package cache

import "container/list"

type LRUCache interface {
	Add(key, value string) bool
	Get(key string) (value string, ok bool)
	Remove(key string) (ok bool)
}

type Entry struct {
	key   string
	value string
}

type MyCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
}

func NewLRUCache(n int) LRUCache {
	return MyCache{
		capacity: n,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (m MyCache) Add(key, value string) bool {
	if elem, ok := m.cache[key]; ok {
		m.list.MoveToFront(elem)
		elem.Value.(*Entry).value = value
		return false
	} else {
		if len(m.cache) >= m.capacity {
			delete(m.cache, m.list.Back().Value.(*Entry).key)
			m.list.Remove(m.list.Back())
		}
		entry := &Entry{key: key, value: value}
		newElem := m.list.PushFront(entry)
		m.cache[key] = newElem
		return true
	}
}

func (m MyCache) Get(key string) (value string, ok bool) {
	if elem, ok := m.cache[key]; ok {
		m.list.MoveToFront(elem)
		return elem.Value.(*Entry).value, true
	}
	return "", false
}

func (m MyCache) Remove(key string) (ok bool) {
	if elem, ok := m.cache[key]; ok {
		delete(m.cache, key)
		m.list.Remove(elem)
		return true
	}
	return false
}
