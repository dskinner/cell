package automata

import (
	"testing"
)

func Benchmark_len(b *testing.B) {
	b.StopTimer()
	var arr []bool = make([]bool, 50*50)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = 50*len(arr)+20
	}
}

func Benchmark_static(b *testing.B) {
	b.StopTimer()
	var arr []bool = make([]bool, 50*50)
	l := len(arr)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = 50*l+20
	}
}

func Benchmark_Step(b *testing.B) {
	b.StopTimer()
	l := New(1000, "0ms", 1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.Step()
	}
}
