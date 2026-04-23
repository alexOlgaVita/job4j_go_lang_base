package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Lru_Cache(t *testing.T) {
	t.Parallel()

	t.Run("{Cache has size 0. Try to put one element once "+
		"and read value by an key once.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(0)
		cache.Put("hello", "world")
		res := cache.Get("hello")
		assert.Empty(t, res)
	})

	t.Run("{Cache has size 1. Put one element once "+
		"and read value by an existing key once.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(1)
		cache.Put("hello", "world")
		res := cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 1. Put one element once "+
		"and try to read value by an not existing key once.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(1)
		cache.Put("hello", "world")
		res := cache.Get("hi")
		assert.Empty(t, res)
	})

	t.Run("{Cache has size 1. Put one element once "+
		"and read value by an existing key twice.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(1)
		cache.Put("hello", "world")
		res := cache.Get("hello")
		res = cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 1. Put the same element twice "+
		"and read value by an existing key once.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(1)
		cache.Put("hello", "world")
		cache.Put("hello", "world")
		res := cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 1. Put two dfferent elements "+
		"and read value by the first key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(1)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		res := cache.Get("hello")
		assert.Empty(t, res)
	})

	t.Run("{Cache has size 1. Put two dfferent elements "+
		"and read value by the second key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(1)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		res := cache.Get("hi")
		exp := "buddy"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 1. Put two dfferent elements "+
		"and read value by the second key twice.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(1)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		res := cache.Get("hi")
		res = cache.Get("hi")
		exp := "buddy"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put the onr element "+
		"and read value by an existing key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		res := cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put the onr element "+
		"and read value by an existing key twice.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		res := cache.Get("hello")
		res = cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put the onr element "+
		"and read value by an existing key thrice.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		res := cache.Get("hello")
		res = cache.Get("hello")
		res = cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put the same element twice "+
		"and read value by an existing key once.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		cache.Put("hello", "world")
		res := cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put the same element twice "+
		"and read value by an existing key twice.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		cache.Put("hello", "world")
		res := cache.Get("hello")
		res = cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put two different elements "+
		"and read value by the first key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		res := cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put two different elements "+
		"and read value by the second key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		res := cache.Get("hi")
		exp := "buddy"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put two different elements "+
		"and read value by the first key twice.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		res := cache.Get("hello")
		res = cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put two different elements "+
		"and read value by the second key twice.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		res := cache.Get("hi")
		res = cache.Get("hi")
		exp := "buddy"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put two different elements "+
		"and read value by the first key twice, then by the second key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		res := cache.Get("hello")
		res = cache.Get("hello")
		res = cache.Get("hi")
		exp := "buddy"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 2. Put two different elements "+
		"and read value by the second key twice, then by the first key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		res := cache.Get("hi")
		res = cache.Get("hi")
		res = cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 3. Put three different elements "+
		"and read value by the second key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		cache.Put("salut", "monsieur")
		res := cache.Get("hi")
		exp := "buddy"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 3. Put three different elements "+
		"and read value by the second key twice then by the first key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		cache.Put("salut", "monsieur")
		res := cache.Get("hi")
		res = cache.Get("hi")
		res = cache.Get("hi")
		res = cache.Get("hello")
		exp := "world"
		assert.Equal(t, exp, *res)
	})

	t.Run("{Cache has size 3. Put three different elements "+
		"and read value by the second key twice then by the third key.", func(t *testing.T) {
		t.Parallel()
		cache := base.NewLruCache(3)
		cache.Put("hello", "world")
		cache.Put("hi", "buddy")
		cache.Put("salut", "monsieur")
		res := cache.Get("hi")
		res = cache.Get("hi")
		res = cache.Get("hi")
		res = cache.Get("salut")
		exp := "monsieur"
		assert.Equal(t, exp, *res)
	})
}
