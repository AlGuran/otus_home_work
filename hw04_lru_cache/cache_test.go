package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("aaa", 100) // [aaa:100]
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200) // [bbb:200, aaa:100]
		require.False(t, wasInCache)

		wasInCache = c.Set("ccc", 300) // [ccc:300, bbb:200, aaa:100]
		require.False(t, wasInCache)

		wasInCache = c.Set("ddd", 400) // [ddd:400, ccc:300, bbb:200]
		require.False(t, wasInCache)

		val, ok := c.Get("aaa") // Удален
		require.False(t, ok)
		require.Equal(t, nil, val)

		val, ok = c.Get("bbb") // [bbb:200, ddd:400, ccc:300]
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("ccc", 999) // [ccc:999, bbb:200, ddd:400]
		require.True(t, wasInCache)

		val, ok = c.Get("ddd") // [ddd:400, ccc:999, bbb:200]
		require.True(t, ok)
		require.Equal(t, 400, val)

		wasInCache = c.Set("bbb", 111) // [bbb:111, ddd:400, ccc:999]
		require.True(t, wasInCache)

		val, ok = c.Get("ddd") // [ddd:400, bbb:111, ccc:999]
		require.True(t, ok)
		require.Equal(t, 400, val)

		val, ok = c.Get("bbb") // [bbb:111, ddd:400, ccc:999]
		require.True(t, ok)
		require.Equal(t, 111, val)

		wasInCache = c.Set("aaa", 101) // [aaa:101, bbb:111, ddd:400]
		require.False(t, wasInCache)

		val, ok = c.Get("ccc") // Удален
		require.Equal(t, nil, val)
		require.False(t, ok)
	})

	t.Run("clear cache", func(t *testing.T) {
		c := NewCache(2)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		c.Clear()

		val, ok := c.Get("aaa")
		require.Equal(t, nil, val)
		require.False(t, ok)

		val, ok = c.Get("bbb")
		require.Equal(t, nil, val)
		require.False(t, ok)

		wasInCache = c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
