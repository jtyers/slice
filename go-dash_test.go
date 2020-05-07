package main

//go:generate ./go-dash-slice -out go-dash_generated_test.go -package main -type string -dir .

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcat(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		add    []string
		output []string
	}{
		{
			"should concat a new item",
			[]string{"first", "second", "third"},
			[]string{"fourth", "fifth"},
			[]string{"first", "second", "third", "fourth", "fifth"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.Concat(test.add)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestDrop(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		i      int
		output []string
	}{
		{
			"should drop items from the left",
			[]string{"first", "second", "third"},
			2,
			[]string{"third"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.Drop(test.i)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestDropRight(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		i      int
		output []string
	}{
		{
			"should drop items from the right",
			[]string{"first", "second", "third"},
			2,
			[]string{"first"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.DropRight(test.i)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestFilter(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		f      func(string, int) bool
		output []string
	}{
		{
			"should run items via filter function",
			[]string{"first", "second", "third"},
			func(s string, i int) bool {
				return s[len(s)-1:len(s)] == "d"
			},
			[]string{"second", "third"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.Filter(test.f)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestFirst(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		output []string
	}{
		{
			"should return first item",
			[]string{"first", "second", "third"},
			[]string{"first"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.First()

		require.Equal(t, got.Value(), test.output)
	}
}

func TestLast(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		output []string
	}{
		{
			"should return last item",
			[]string{"first", "second", "third"},
			[]string{"third"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.Last()

		require.Equal(t, got.Value(), test.output)
	}
}

func TestMap(t *testing.T) {
	var tests = []struct {
		name    string
		input   []string
		mapFunc func(string, int) string
		output  []string
	}{
		{
			"should map input to output for simple function",
			[]string{"first", "second", "third"},
			func(s string, i int) string {
				return strings.ToUpper(s)
			},
			[]string{"FIRST", "SECOND", "THIRD"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.Map(test.mapFunc)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestReduce(t *testing.T) {
	var tests = []struct {
		name    string
		input   []string
		initial string
		f       func(string, string, int) string
		output  []string
	}{
		{
			"should reduce input to output via reduce function",
			[]string{"first", "second", "third"},
			"initial",
			func(acc string, val string, i int) string {
				return acc + "-" + strings.ToUpper(val)
			},
			[]string{"initial-FIRST-SECOND-THIRD"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.Reduce(test.f, test.initial)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		output []string
	}{
		{
			"should reverse slice order",
			[]string{"first", "second", "third"},
			[]string{"third", "second", "first"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.Reverse()

		require.Equal(t, got.Value(), test.output)
	}
}

func TestUniq(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		output []string
	}{
		{
			"should filter out duplicates where duplicates are present",
			[]string{"first", "second", "third", "second", "first"},
			[]string{"first", "second", "third"},
		},
		{
			"should filter out duplicates where no duplicates are present",
			[]string{"first", "second", "third"},
			[]string{"first", "second", "third"},
		},
	}

	for _, test := range tests {
		c := NewStringSlice(test.input)

		got := c.Uniq()

		require.Equal(t, got.Value(), test.output)
	}
}
