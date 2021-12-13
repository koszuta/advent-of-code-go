package main

import "testing"

func TestPart2(t *testing.T) {
	pos, depth := doPart2()
	actual := pos * depth
	if actual != expectedResult {
		t.Logf("test failed; expected %d, got %d", expectedResult, actual)
		t.FailNow()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doPart2()
	}
}
