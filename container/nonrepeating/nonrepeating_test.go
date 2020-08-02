package main

import "testing"

func TestNonRepeating(t *testing.T) {
	tests := []struct {
		s string
		a int
	}{
		// normal cases
		{"abcabcabc", 3},
		{"abcabcabcd", 4},
		{"fsdghjpo", 8},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"n", 1},
		{"bbbbb", 1},

		// Chinese support
		{"今天天气好好啊", 3},
		{"一二三四三二一", 4},
		{"黑化肥挥发发挥会花飞汇划费挥发发黑会飞花", 8},
	}

	for _, test := range tests {
		if actual := lengthOfNonRepeatingSubStr(test.s); actual != test.a {
			t.Errorf("lengthOfNonRepeatingSubStr(%s); got %d; excepted %d", test.s, actual, test.a)
		}
	}
}

func BenchmarkNonRepeating(b *testing.B) {
	s := "黑化肥挥发发挥会花飞汇划费挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	a := 8
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != a {
			b.Errorf("lengthOfNonRepeatingSubStr(%s); got %d; excepted %d", s, actual, a)
		}
	}
}

// 性能测试
// go test -bench . -cpuprofile cpu.out
// go tool pprof cpu.out
// web
