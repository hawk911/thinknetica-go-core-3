package lesson_7

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func sampleDataFloat() []float64 {
	rand.Seed(time.Now().UnixNano())
	var data []float64
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Float64())
	}
	return data
}

func BenchmarkSortFloat(b *testing.B) {
	data := sampleDataFloat()
	for i := 0; i < b.N; i++ {
		sort.Float64s(data)
	}
}

//goos: windows
//goarch: amd64
//pkg: github/hawk911
//cpu: AMD Ryzen 7 4800H
//BenchmarkSortFloat
//BenchmarkSortFloat-16                 21          52677929 ns/op
//PASS
//
//Process finished with the exit code 0
