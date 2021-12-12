package main

import "testing"

func TestPart1(t *testing.T) {
	pos, depth := doPart1()
	actual := pos * depth
	if actual != expectedResult {
		t.Logf("test failed; expected %d, got %d", expectedResult, actual)
		t.Fail()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doPart1()
	}
}
