// This code is generated by https://github.com/jtyers/slice
// DO NOT EDIT!

// nfn: NewCustomTypeSlice

package main


import (
  . "github.com/jtyers/slice/customtype"
)

type chainCustomType struct {
  isPtr bool
	value []CustomType
}

func NewCustomTypeSlice(slice []CustomType) *chainCustomType {
	return &chainCustomType{
		value: slice,
		
	}
}

func (c *chainCustomType) Value() []CustomType {
	return c.value
}

func ConcatCustomType(slice []CustomType, slice2 []CustomType) (res []CustomType) {
	res = make([]CustomType, 0, len(slice) + len(slice2))
	for _, entry := range slice {
		res = append(res, entry)
	}
	for _, entry := range slice2 {
		res = append(res, entry)
	}
	return
}

func (c *chainCustomType) Concat(slice2 []CustomType) *chainCustomType {
	return &chainCustomType{value: ConcatCustomType(c.value, slice2)}
}

func DropCustomType(slice []CustomType, n int) (res []CustomType) {
	l := len(slice) - n
	if l < 0 {
		l = 0
	}
	res = make([]CustomType, 0, l)
	for _, entry := range slice[len(slice) - l:] {
		res = append(res, entry)
	}
	return
}

func (c *chainCustomType) Drop(n int) *chainCustomType {
	return &chainCustomType{value: DropCustomType(c.value, n)}
}

func DropRightCustomType(slice []CustomType, n int) (res []CustomType) {
	l := len(slice) - n
	if l < 0 {
		l = 0
	}
	res = make([]CustomType, 0, l)
	for _, entry := range slice[:l] {
		res = append(res, entry)
	}
	return
}

func (c *chainCustomType) DropRight(n int) *chainCustomType {
	return &chainCustomType{value: DropRightCustomType(c.value, n)}
}

func FilterCustomType(slice []CustomType, fn func(CustomType,int)bool) (res []CustomType) {
	res = make([]CustomType, 0, len(slice))
	for index, entry := range slice {
		if fn(entry, index) {
			res = append(res, entry)
		}
	}
	return
}

func (c *chainCustomType) Filter(fn func(CustomType,int)bool) *chainCustomType {
	return &chainCustomType{value: FilterCustomType(c.value, fn)}
}

func FirstCustomType(slice []CustomType) (res CustomType) {
	if len(slice) == 0 {
		return
	}
	res = slice[0]
	return
}

func (c *chainCustomType) First() *chainCustomType {
	return &chainCustomType{value: []CustomType{FirstCustomType(c.value)}}
}

func LastCustomType(slice []CustomType) (res CustomType) {
	if len(slice) == 0 {
		return
	}
	res = slice[len(slice) - 1]
	return
}

func (c *chainCustomType) Last() *chainCustomType {
	return &chainCustomType{value: []CustomType{LastCustomType(c.value)}}
}

func MapCustomType(slice []CustomType, fn func(CustomType,int)CustomType) (res []CustomType) {
	res = make([]CustomType, 0, len(slice))
	for index, entry := range slice {
		res = append(res, fn(entry, index))
	}
	return
}

func (c *chainCustomType) Map(fn func(CustomType,int)CustomType) *chainCustomType {
	return &chainCustomType{value: MapCustomType(c.value, fn)}
}


func ReduceCustomType(slice []CustomType, fn func(CustomType,CustomType,int)CustomType, initial CustomType) (res CustomType) {
	res = initial
	for index, entry := range slice {
		res = fn(res, entry, index)
	}
	return
}

func (c *chainCustomType) Reduce(fn func(CustomType,CustomType,int)CustomType, initial CustomType) *chainCustomType {
	return &chainCustomType{value: []CustomType{ReduceCustomType(c.value, fn, initial)}}
}

func ReverseCustomType(slice []CustomType) (res []CustomType) {
	res = make([]CustomType, len(slice))
	for index, entry := range slice {
		res[len(slice)-1-index] = entry
	}
	return
}

func (c *chainCustomType) Reverse() *chainCustomType {
	return &chainCustomType{value: ReverseCustomType(c.value)}
}

func UniqCustomType(slice []CustomType) (res []CustomType) {
	seen := make(map[CustomType]bool)
	res = []CustomType{}
	for _, entry := range slice {
		if _, found := seen[entry]; !found {
			seen[entry] = true
			res = append(res, entry)
		}
	}
	return
}

func (c *chainCustomType) Uniq() *chainCustomType {
	if c.isPtr {
		panic("Uniq() does not currently support pointers")
	}
	return &chainCustomType{value: UniqCustomType(c.value)}
}
