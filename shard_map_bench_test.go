package cmap_test

import (
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/LyricTian/cmap"
)

func BenchmarkShardMapSet(b *testing.B) {
	m := cmap.NewShardMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(i), i)
	}
}

func BenchmarkParallelShardMapSet(b *testing.B) {
	m := cmap.NewShardMap()
	b.ResetTimer()

	var i int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v := strconv.Itoa(int(atomic.AddInt64(&i, 1)))
			m.Set(v, v)
		}
	})
}

func BenchmarkShardMapGet(b *testing.B) {
	m := cmap.NewShardMap()
	m.Set("foo", "bar")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get("foo")
	}
}

func BenchmarkParallelShardMapGet(b *testing.B) {
	m := cmap.NewShardMap()
	m.Set("foo", "bar")
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Get("foo")
		}
	})
}
