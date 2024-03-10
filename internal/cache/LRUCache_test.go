package cache

import (
	"container/list"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMyCache_Add(t *testing.T) {
	req := require.New(t)
	t.Run("add new elem", func(t *testing.T) {
		cache := MyCache{
			capacity: 2,
			cache:    map[string]*list.Element{},
			list:     list.New(),
		}
		got := cache.Add("1", "testcase1")
		req.Equal(got, true)
	})
	t.Run("add exist elem", func(t *testing.T) {
		cache := MyCache{
			capacity: 2,
			cache:    map[string]*list.Element{},
			list:     list.New(),
		}
		got := cache.Add("1", "testcase1")
		got = cache.Add("1", "testcase1")
		req.Equal(got, false)
	})
	t.Run("add elem to full cache", func(t *testing.T) {
		cache := MyCache{
			capacity: 2,
			cache:    map[string]*list.Element{},
			list:     list.New(),
		}
		got := cache.Add("1", "testcase1")
		got = cache.Add("2", "testcase2")
		got = cache.Add("3", "testcase3")
		req.Equal(got, true)
	})
}

func TestMyCache_Get(t *testing.T) {
	req := require.New(t)
	t.Run("get not exist elem", func(t *testing.T) {
		cache := MyCache{
			capacity: 2,
			cache:    map[string]*list.Element{},
			list:     list.New(),
		}
		got, err := cache.Get("1")
		req.Equal(err, false)
		req.Equal(got, "")
	})
	t.Run("get exist elem", func(t *testing.T) {
		cache := MyCache{
			capacity: 2,
			cache:    map[string]*list.Element{},
			list:     list.New(),
		}
		err := cache.Add("1", "testcase2")
		req.Equal(err, true)
		got, err := cache.Get("1")
		req.Equal(err, true)
		req.Equal(got, "testcase2")
	})
}

func TestMyCache_Remove(t *testing.T) {
	req := require.New(t)
	t.Run("remove exist elem", func(t *testing.T) {
		cache := MyCache{
			capacity: 2,
			cache:    map[string]*list.Element{},
			list:     list.New(),
		}
		err := cache.Add("1", "testcase2")
		req.Equal(err, true)
		err = cache.Remove("1")
		req.Equal(err, true)
	})
	t.Run("remove exist elem", func(t *testing.T) {
		cache := MyCache{
			capacity: 2,
			cache:    map[string]*list.Element{},
			list:     list.New(),
		}
		err := cache.Remove("1")
		req.Equal(err, false)
	})
}

func TestNewLRUCache(t *testing.T) {
	//req := require.New(t)
	t.Run("create new lru cache", func(t *testing.T) {
		NewLRUCache(2)
	})
}
