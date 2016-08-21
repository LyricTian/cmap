package cmap_test

import (
	"testing"

	"sync/atomic"

	"strconv"

	"github.com/LyricTian/cmap"
)

func BenchmarkMapSet(b *testing.B) {
	m := cmap.NewMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(i), i)
	}
}

func BenchmarkParallelMapSet(b *testing.B) {
	m := cmap.NewMap()
	b.ResetTimer()

	var i int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v := strconv.Itoa(int(atomic.AddInt64(&i, 1)))
			m.Set(v, v)
		}
	})
}

func BenchmarkMapGet(b *testing.B) {
	m := cmap.NewMap()
	m.Set("foo", "bar")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get("foo")
	}
}

func BenchmarkParallelMapGet(b *testing.B) {
	m := cmap.NewMap()
	m.Set("foo", "bar")
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Get("foo")
		}
	})
}
