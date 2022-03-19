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
		{[]string{"آبی", "قرمز"}, []string{"آبی", "قرمز"}},
		{[]string{"آبی", "قرمز", "قرمز"}, []string{"آبی", "قرمز"}},
		{nil, []string{}},
	}

	for _, tt := range tests {
		got := String(tt.slice)
		require.Equal(t, tt.expect, got)
	}
}

func TestStringsAreUnique(t *testing.T) {
	tests := []struct {
		slice  []string
		expect bool
	}{
		{[]string{"a", "b", "ab", "ba", "ab\t", "a", "a1", "ab1", "ab\t"}, false},
		{[]string{"$\t", "12%87231NXAS"}, true},
		{[]string{"list\n\t"}, true},
		{[]string{"sal", "sal", "sal", "sal", "sal", "sal", "sal", "sal", "sal"}, false},
		{[]string{}, true},
		{[]string{"آبی", "قرمز"}, true},
		{[]string{"آبی", "قرمز", "قرمز"}, false},
		{nil, true},
	}

	for _, tt := range tests {
		got := StringsAreUnique(tt.slice)
		require.Equal(t, tt.expect, got)
	}

}
