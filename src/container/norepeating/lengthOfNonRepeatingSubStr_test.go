package main

import "testing"

func TestLengthOfNonRepeatingSubStr(t *testing.T) {
	tests := []struct {
		s   string
		num int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},

		{"", 0},
		{"b", 1},
		{"bbbbbbbbbbbbb", 1},
		{"abcabcabcd", 4},

		{"这是Golang学习代码", 12},
		{"一二三一二", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, value := range tests {
		if result := lengthOfNonRepeatingSubStr(value.s); result != value.num {
			t.Errorf("lengthOfNonRepeatingSubStr(%s)=%d;but result is %d", value.s, value.num, result)
		}
	}
}

func BenchmarkLengthOfNonRepeatingSubStr(b *testing.B) {
	tests := []struct {
		s   string
		num int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},

		{"", 0},
		{"b", 1},
		{"bbbbbbbbbbbbb", 1},
		{"abcabcabcd", 4},

		{"这是Golang学习代码", 12},
		{"一二三一二", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}
	b.ResetTimer() //重置时间，抛除之前数据准备所花费的时间

	for i := 0; i < b.N; i++ {
		for _, value := range tests {
			if result := lengthOfNonRepeatingSubStr(value.s); result != value.num {
				b.Errorf("lengthOfNonRepeatingSubStr(%s)=%d;but result is %d", value.s, value.num, result)
			}
		}
	}
}
