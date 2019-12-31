package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, value := range tests {
		if result := caclcTringle(value.a, value.b); result != value.c {
			t.Errorf("caclcTringle(%d, %d) = %d; but the result is %d", value.a, value.b, value.c, result)
		}
	}
}
