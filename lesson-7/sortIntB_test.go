package lesson_7

import (
	"math/rand"
	"testing"
	"time"
)

func sampleData() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}
	return data
}

func BenchmarkSortInt(b *testing.B) {
	data := sampleData()
	for i := 0; i < b.N; i++ {
		SortInt(data)
	}
}

//goos: windows
//goarch: amd64
//pkg: github/hawk911
//cpu: AMD Ryzen 7 4800H
//BenchmarkSortInt
//BenchmarkSortInt-16           44          26835045 ns/op
//PASS
//
//Process finished with the exit code 0
