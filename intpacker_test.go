package intpacker

import (
	"math"
	"sync"
	"sync/atomic"
	"testing"
)

func TestUint32(t *testing.T) {
	tt := []struct {
		n1 uint32
		n2 uint32
	}{
		{n1: 10, n2: 11},
		{n1: math.MaxUint32, n2: 0},
		{n1: 0, n2: math.MaxUint32},
		{n1: math.MaxUint32, n2: math.MaxUint32},
	}
	for _, args := range tt {
		i := NewUint32(args.n1, args.n2)
		got1, got2 := i.Unpack()
		if args.n1 != got1 || args.n2 != got2 {
			t.Errorf("Did not unpack, n1=%v got1=%v n2=%v got2=%v", args.n1, got1, args.n2, got2)
		}
	}
}

func TestUint32WorksWithAtomic(t *testing.T) {
	type val struct {
		n1 uint32
		n2 uint32
	}
	args := []struct {
		arg  val
		want val
	}{
		{arg: val{n1: 1, n2: 2}, want: val{n1: 11, n2: 13}},
		// Subtract from the first number. Subtract from the second is unsupported
		{arg: val{n1: ^uint32(0), n2: 2}, want: val{n1: 9, n2: 13}},
	}
	for _, tt := range args {
		i := NewUint32(10, 11)
		x := NewUint32(tt.arg.n1, tt.arg.n2)
		atomic.AddUint64(i.Ptr(), x.Uint64())
		got1, got2 := i.Unpack()
		if tt.want.n1 != got1 || tt.want.n2 != got2 {
			t.Errorf("Did not add correctly want1=%v got1=%v want2=%v got2=%v", tt.want.n1, got1, tt.want.n2, got2)
		}
	}
}

func BenchmarkUint32AtomicAdd(b *testing.B) {
	o := NewUint32(1, 2)
	incr := uint32(0)
	for x := 0; x < b.N; x++ {
		atomic.AddUint64(o.Ptr(), NewUint32(incr, incr).Uint64())
		incr++
	}
}

func BenchmarkUint32WithMutex(b *testing.B) {
	o1 := uint32(1)
	o2 := uint32(2)
	incr := uint32(0)
	var mu sync.Mutex
	for x := 0; x < b.N; x++ {
		mu.Lock()
		o1 += incr
		o2 += incr
		mu.Unlock()
		incr++
	}
}

func BenchmarkUint32AtomicAddParallel(b *testing.B) {
	o := NewUint32(1, 2)
	b.SetParallelism(3)
	b.RunParallel(func(pb *testing.PB) {
		incr := uint32(0)
		for pb.Next() {
			atomic.AddUint64(o.Ptr(), NewUint32(incr, incr).Uint64())
			incr++
		}
	})
}

func BenchmarkUint32WithMutexParallel(b *testing.B) {
	o1 := uint32(1)
	o2 := uint32(2)
	var mu sync.Mutex
	b.SetParallelism(3)
	b.RunParallel(func(pb *testing.PB) {
		incr := uint32(0)
		for pb.Next() {
			mu.Lock()
			o1 += incr
			o2 += incr
			mu.Unlock()
			incr++
		}
	})
}
