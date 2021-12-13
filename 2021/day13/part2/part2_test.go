package main

import "testing"

func TestPart2(t *testing.T) {
	actual := doPart2()
	if actual != expectedResult {
		t.Logf("test failed; expected \"%s\", got \"%s\"", expectedResult, actual)
		t.FailNow()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doPart2()
	}
}
