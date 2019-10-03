package cmprobin

import (
	"math/rand"
	"testing"
	"time"
	"unsafe"

	"github.com/glaslos/cmprobin/lewuathe"
	"github.com/glaslos/cmprobin/petermattis"
)

const benchSize = 1 << 20

func BenchmarkPeterMattis(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	keys := make([]uint64, benchSize)
	for i := range keys {
		keys[i] = uint64(rng.Intn(1 << 20))
	}
	v := unsafe.Pointer(new(int))
	b.ResetTimer()

	var m *petermattis.RobinHoodMap
	for i, j := 0, 0; i < b.N; i, j = i+1, j+1 {
		if m == nil || j == len(keys) {
			b.StopTimer()
			m = petermattis.NewRobinHoodMap(len(keys))
			j = 0
			b.StartTimer()
		}
		m.Put(keys[j], v)
	}
}

func BenchmarkLewuathe(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	keys := make([]string, benchSize)
	for i := range keys {
		keys[i] = string(uint64(rng.Intn(1 << 20)))
	}
	v := unsafe.Pointer(new(int))
	b.ResetTimer()

	var m *lewuathe.RobinHood
	for i, j := 0, 0; i < b.N; i, j = i+1, j+1 {
		if m == nil || j == len(keys) {
			b.StopTimer()
			m = lewuathe.NewRobinHood(len(keys))
			j = 0
			b.StartTimer()
		}
		m.Put(keys[j], v)
	}
}
