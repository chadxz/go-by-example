package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	result := IntMin(2, -2)
	if result != -2 {
		// would normally use testify assertions for this in real life.
		// sticking with stdlib for now
		t.Errorf("IntMin(2, -2) = %d; want -2", result)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("%d,%d", testCase.a, testCase.b)
		t.Run(testName, func(t *testing.T) {
			result := IntMin(testCase.a, testCase.b)
			if result != testCase.want {
				t.Errorf("got %d, want %d", result, testCase.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
