package unittesting

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetTimeDiffInMins(t *testing.T) {
	curTime := time.Now()
	var testCases = []struct {
		label string
		t1 time.Time
		t2 time.Time
		expected int64
	}{
		{"positiveDiff", curTime.Add(5*time.Minute), curTime, 5},
		{"negativeDiff", curTime, curTime.Add(5*time.Minute),  -5},
		{"noDiff", curTime, curTime, 0},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, GetTimeDiffInMins(tc.t1, tc.t2), tc.label)
	}
}