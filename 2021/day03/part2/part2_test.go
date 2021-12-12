package main

import "testing"

func TestPart2(t *testing.T) {
	g, e := doPart2()
	actual := g * e
	if actual != expectedResult {
		t.Logf("test failed; expected %d, got %d", expectedResult, actual)
		t.Fail()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doPart2()
	}
}
