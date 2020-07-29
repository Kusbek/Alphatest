package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fin(t *testing.T) {
	testCases := []struct {
		num          int
		total        int
		size         int
		expectedPrev int
		expectedNext int
	}{
		{
			num:          1,
			total:        6,
			size:         20,
			expectedPrev: 0,
			expectedNext: 2,
		},
		{
			num:          1,
			total:        5,
			size:         20,
			expectedPrev: 0,
			expectedNext: 2,
		},
		{
			num:          4,
			total:        5,
			size:         20,
			expectedPrev: 3,
			expectedNext: 5,
		},
		{
			num:          5,
			total:        5,
			size:         20,
			expectedPrev: 4,
			expectedNext: 0,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {

			prev, next := findPrevNext(tc.total, tc.num, tc.size)
			assert.Equal(t, tc.expectedPrev, prev)
			assert.Equal(t, tc.expectedNext, next)
		})
	}
}
