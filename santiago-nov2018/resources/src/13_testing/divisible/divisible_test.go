package divisible_test

import (
	"testing"

	"github.com/ifreddyrondon/go-workshop/santiago-nov2018/resources/src/13_testing/divisible"
)

func TestSum(t *testing.T) {
	tt := []struct {
		name   string
		top    int
		by     []int
		expect int
	}{
		{
			name:   "when top 0 expected 0",
			top:    0,
			by:     []int{3},
			expect: 0,
		},
		{
			name:   "when top 3 expected 3",
			top:    3,
			by:     []int{3},
			expect: 3,
		},
		{
			name:   "when top 15 expected 45",
			top:    15,
			by:     []int{3},
			expect: 45,
		},
		{
			name:   "when top 15 and by 3, and 5 expected 75",
			top:    15,
			by:     []int{3, 5},
			expect: 75,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := divisible.Sum(tc.top, tc.by...)
			if tc.expect != result {
				t.Errorf("test err: expected %v but got %v", tc.expect, result)
			}
		})
	}
}
