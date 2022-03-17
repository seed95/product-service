package unique

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		slice  []int
		expect []int
	}{
		{[]int{1, 2, 4, 1, 4, 7, 2}, []int{1, 2, 4, 7}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{1}},
		{[]int{}, []int{}},
		{nil, []int{}},
	}

	for _, tt := range tests {
		got := Int(tt.slice)
		require.Equal(t, tt.expect, got)
	}

}

func TestString(t *testing.T) {
	tests := []struct {
		slice  []string
		expect []string
	}{
		{[]string{"a", "b", "ab", "ba", "ab\t", "a", "a1", "ab1", "ab\t"}, []string{"a", "b", "ab", "ba", "ab\t", "a1", "ab1"}},
		{[]string{"$\t", "12%87231NXAS"}, []string{"$\t", "12%87231NXAS"}},
		{[]string{"list\n\t"}, []string{"list\n\t"}},
		{[]string{"sal", "sal", "sal", "sal", "sal", "sal", "sal", "sal", "sal"}, []string{"sal"}},
		{[]string{}, []string{}},
		{nil, []string{}},
	}

	for _, tt := range tests {
		got := String(tt.slice)
		require.Equal(t, tt.expect, got)
	}

}
