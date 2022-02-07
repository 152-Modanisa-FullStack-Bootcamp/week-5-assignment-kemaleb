package assignment

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	type args struct {
		x uint32
		y uint32
	}
	tests := []struct {
		name  string
		args  args
		want  uint32
		want1 bool
	}{
		{"1 test", args{math.MaxUint32, 1}, 0, true},
		{"2 test", args{1, 1}, 2, false},
		{"3 test", args{42, 2701}, 2743, false},
		{"4 test", args{42, math.MaxUint32}, 41, true},
		{"5 test", args{4294967290, 5}, 4294967295, false},
		{"6 test", args{4294967290, 6}, 0, true},
		{"7 test", args{4294967290, 10}, 4, true},
	}
	for _, tt := range tests {
		i, j := AddUint32(tt.args.x, tt.args.y)
		assert.Equal(t, i, tt.want)
		assert.Equal(t, j, tt.want1)
	}
}

func TestCeilNumber(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		args args
		want float64
	}{
		{args{42.42}, 42.50},
		{args{42}, 42},
		{args{42.01}, 42.25},
		{args{42.24}, 42.25},
		{args{42.25}, 42.25},
		{args{42.26}, 42.50},
		{args{42.55}, 42.75},
		{args{42.75}, 42.75},
		{args{42.76}, 43},
		{args{42.99}, 43},
		{args{43.13}, 43.25},
	}
	for _, tt := range tests {
		i := CeilNumber(tt.args.f)
		assert.Equal(t, i, tt.want)

	}
}

func TestAlphabetSoup(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1 test", args{"hello"}, "ehllo"},
		{"2 test", args{""}, ""},
		{"3 test", args{"ab"}, "ab"},
		{"4 test", args{"ba"}, "ab"},
		{"5 test", args{"bac"}, "abc"},
		{"6 test", args{"cba"}, "abc"},
		{"7 test", args{"h"}, "h"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlphabetSoup(tt.args.s); got != tt.want {
				t.Errorf("AlphabetSoup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringMask(t *testing.T) {
	type args struct {
		s string
		n uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1 test", args{"!mysecret*", 2}, "!m********"},
		{"2 test", args{"", 2}, "*"},
		{"3 test", args{"a", 1}, "*"},
		{"4 test", args{"string", 0}, "******"},
		{"5 test", args{"string", 3}, "str***"},
		{"6 test", args{"string", 5}, "strin*"},
		{"7 test", args{"string", 6}, "******"},
		{"8 test", args{"string", 7}, "******"},
		{"9 test", args{"s*r*n*", 3}, "s*r***"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringMask(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("StringMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,hello,cat,bat,goodbye,yellow,why"
	type args struct {
		arr [2]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1 test", args{[2]string{"hellocat", words}}, "hello,cat,"},
		{"2 test", args{[2]string{"catbat", words}}, "cat,bat,"},
		{"3 test", args{[2]string{"yellowapple", words}}, "apple,yellow,"},
		{"4 test", args{[2]string{"", words}}, "not possible"},
		{"5 test", args{[2]string{"notcat", words}}, "not possible"},
		{"6 test", args{[2]string{"bootcamprocks!", words}}, "not possible"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordSplit(tt.args.arr); got != tt.want {
				t.Errorf("WordSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariadicSet(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VariadicSet(tt.args.i...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VariadicSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
