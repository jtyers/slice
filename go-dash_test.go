package main

//go:generate ./slice -out go-dash_generated_test.go -package main -type string -dir .
//go:generate ./slice -out go-dash_generated_ptr_test.go -package main -type *string -dir .
//go:generate ./slice -out go-dash_generated_custom_test.go -package main -type CustomType -import github.com/jtyers/go-dash-slice -dir .

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringConcat(t *testing.T) {
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

func TestStringDrop(t *testing.T) {
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

func TestStringDropRight(t *testing.T) {
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

func TestStringFilter(t *testing.T) {
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

func TestStringFirst(t *testing.T) {
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

func TestStringLast(t *testing.T) {
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

func TestStringMap(t *testing.T) {
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

func TestStringReduce(t *testing.T) {
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

func TestStringReverse(t *testing.T) {
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

func TestStringUniq(t *testing.T) {
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

func stringPtr(s string) *string {
	return &s
}

func stringPtrSlice(s []string) []*string {
	result := []*string{}
	for _, v := range s {
		result = append(result, stringPtr(v))
	}
	return result
}

func TestStringPtrConcat(t *testing.T) {
	var tests = []struct {
		name   string
		input  []*string
		add    []*string
		output []*string
	}{
		{
			"should concat a new item",
			stringPtrSlice([]string{"first", "second", "third"}),
			stringPtrSlice([]string{"fourth", "fifth"}),
			stringPtrSlice([]string{"first", "second", "third", "fourth", "fifth"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.Concat(test.add)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestStringPtrDrop(t *testing.T) {
	var tests = []struct {
		name   string
		input  []*string
		i      int
		output []*string
	}{
		{
			"should drop items from the left",
			stringPtrSlice([]string{"first", "second", "third"}),
			2,
			stringPtrSlice([]string{"third"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.Drop(test.i)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestStringPtrDropRight(t *testing.T) {
	var tests = []struct {
		name   string
		input  []*string
		i      int
		output []*string
	}{
		{
			"should drop items from the right",
			stringPtrSlice([]string{"first", "second", "third"}),
			2,
			stringPtrSlice([]string{"first"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.DropRight(test.i)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestStringPtrFilter(t *testing.T) {
	var tests = []struct {
		name   string
		input  []*string
		f      func(*string, int) bool
		output []*string
	}{
		{
			"should run items via filter function",
			stringPtrSlice([]string{"first", "second", "third"}),
			func(s *string, i int) bool {
				ss := *s
				return ss[len(ss)-1:len(ss)] == "d"
			},
			stringPtrSlice([]string{"second", "third"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.Filter(test.f)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestStringPtrFirst(t *testing.T) {
	var tests = []struct {
		name   string
		input  []*string
		output []*string
	}{
		{
			"should return first item",
			stringPtrSlice([]string{"first", "second", "third"}),
			stringPtrSlice([]string{"first"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.First()

		require.Equal(t, got.Value(), test.output)
	}
}

func TestStringPtrLast(t *testing.T) {
	var tests = []struct {
		name   string
		input  []*string
		output []*string
	}{
		{
			"should return last item",
			stringPtrSlice([]string{"first", "second", "third"}),
			stringPtrSlice([]string{"third"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.Last()

		require.Equal(t, got.Value(), test.output)
	}
}

func TestStringPtrMap(t *testing.T) {
	var tests = []struct {
		name    string
		input   []*string
		mapFunc func(*string, int) *string
		output  []*string
	}{
		{
			"should map input to output for simple function",
			stringPtrSlice([]string{"first", "second", "third"}),
			func(s *string, i int) *string {
				result := strings.ToUpper(*s)
				return &result
			},
			stringPtrSlice([]string{"FIRST", "SECOND", "THIRD"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.Map(test.mapFunc)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestStringPtrReduce(t *testing.T) {
	var tests = []struct {
		name    string
		input   []*string
		initial *string
		f       func(*string, *string, int) *string
		output  []*string
	}{
		{
			"should reduce input to output via reduce function",
			stringPtrSlice([]string{"first", "second", "third"}),
			stringPtr("initial"),
			func(acc *string, val *string, i int) *string {
				vv := *acc + "-" + strings.ToUpper(*val)
				return &vv
			},
			stringPtrSlice([]string{"initial-FIRST-SECOND-THIRD"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.Reduce(test.f, test.initial)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestStringPtrReverse(t *testing.T) {
	var tests = []struct {
		name   string
		input  []*string
		output []*string
	}{
		{
			"should reverse slice order",
			stringPtrSlice([]string{"first", "second", "third"}),
			stringPtrSlice([]string{"third", "second", "first"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.Reverse()

		require.Equal(t, got.Value(), test.output)
	}
}

func TestStringPtrUniq(t *testing.T) {
	// currently Uniq() doesn't work with pointers, since it cannot directly compare values
	t.Skip()

	var tests = []struct {
		name   string
		input  []*string
		output []*string
	}{
		{
			"should filter out duplicates where duplicates are present",
			stringPtrSlice([]string{"first", "second", "third", "second", "first"}),
			stringPtrSlice([]string{"first", "second", "third"}),
		},
		{
			"should filter out duplicates where no duplicates are present",
			stringPtrSlice([]string{"first", "second", "third"}),
			stringPtrSlice([]string{"first", "second", "third"}),
		},
	}

	for _, test := range tests {
		c := NewStringPtrSlice(test.input)

		got := c.Uniq()

		require.Equal(t, got.Value(), test.output)
	}
}

func ct(name string) CustomType {
	return CustomType{name}
}

func TestCustomTypeConcat(t *testing.T) {
	var tests = []struct {
		name   string
		input  []CustomType
		add    []CustomType
		output []CustomType
	}{
		{
			"should concat a new item",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			[]CustomType{ct("fourth"), ct("fifth")},
			[]CustomType{ct("first"), ct("second"), ct("third"), ct("fourth"), ct("fifth")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.Concat(test.add)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestCustomTypeDrop(t *testing.T) {
	var tests = []struct {
		name   string
		input  []CustomType
		i      int
		output []CustomType
	}{
		{
			"should drop items from the left",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			2,
			[]CustomType{ct("third")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.Drop(test.i)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestCustomTypeDropRight(t *testing.T) {
	var tests = []struct {
		name   string
		input  []CustomType
		i      int
		output []CustomType
	}{
		{
			"should drop items from the right",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			2,
			[]CustomType{ct("first")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.DropRight(test.i)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestCustomTypeFilter(t *testing.T) {
	var tests = []struct {
		name   string
		input  []CustomType
		f      func(CustomType, int) bool
		output []CustomType
	}{
		{
			"should run items via filter function",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			func(s CustomType, i int) bool {
				return s.name[len(s.name)-1:len(s.name)] == "d"
			},
			[]CustomType{ct("second"), ct("third")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.Filter(test.f)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestCustomTypeFirst(t *testing.T) {
	var tests = []struct {
		name   string
		input  []CustomType
		output []CustomType
	}{
		{
			"should return first item",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			[]CustomType{ct("first")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.First()

		require.Equal(t, got.Value(), test.output)
	}
}

func TestCustomTypeLast(t *testing.T) {
	var tests = []struct {
		name   string
		input  []CustomType
		output []CustomType
	}{
		{
			"should return last item",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			[]CustomType{ct("third")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.Last()

		require.Equal(t, got.Value(), test.output)
	}
}

func TestCustomTypeMap(t *testing.T) {
	var tests = []struct {
		name    string
		input   []CustomType
		mapFunc func(CustomType, int) CustomType
		output  []CustomType
	}{
		{
			"should map input to output for simple function",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			func(s CustomType, i int) CustomType {
				result := strings.ToUpper(s.name)
				return CustomType{result}
			},
			[]CustomType{ct("FIRST"), ct("SECOND"), ct("THIRD")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.Map(test.mapFunc)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestCustomTypeReduce(t *testing.T) {
	var tests = []struct {
		name    string
		input   []CustomType
		initial CustomType
		f       func(CustomType, CustomType, int) CustomType
		output  []CustomType
	}{
		{
			"should reduce input to output via reduce function",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			ct("initial"),
			func(acc CustomType, val CustomType, i int) CustomType {
				vv := acc.name + "-" + strings.ToUpper(val.name)
				return ct(vv)
			},
			[]CustomType{ct("initial-FIRST-SECOND-THIRD")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.Reduce(test.f, test.initial)

		require.Equal(t, got.Value(), test.output)
	}
}

func TestCustomTypeReverse(t *testing.T) {
	var tests = []struct {
		name   string
		input  []CustomType
		output []CustomType
	}{
		{
			"should reverse slice order",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			[]CustomType{ct("third"), ct("second"), ct("first")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.Reverse()

		require.Equal(t, got.Value(), test.output)
	}
}

func TestCustomTypeUniq(t *testing.T) {
	var tests = []struct {
		name   string
		input  []CustomType
		output []CustomType
	}{
		{
			"should filter out duplicates where duplicates are present",
			[]CustomType{ct("first"), ct("second"), ct("third"), ct("second"), ct("first")},
			[]CustomType{ct("first"), ct("second"), ct("third")},
		},
		{
			"should filter out duplicates where no duplicates are present",
			[]CustomType{ct("first"), ct("second"), ct("third")},
			[]CustomType{ct("first"), ct("second"), ct("third")},
		},
	}

	for _, test := range tests {
		c := NewCustomTypeSlice(test.input)

		got := c.Uniq()

		require.Equal(t, got.Value(), test.output)
	}
}
