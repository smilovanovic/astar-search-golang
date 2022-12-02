package main

import (
	"strings"
	"testing"
)

var result Result

func BenchmarkAstar(b *testing.B) {
	var s Result
	for i := 0; i < b.N; i++ {
		niz := strings.Split("grbrgbbrggbr", "")
		s = astar(&niz)
	}
	result = s
}
