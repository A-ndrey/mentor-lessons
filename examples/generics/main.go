package main

import (
	"fmt"
	"reflect"
)

type PointsIDs = []int

func main() {
	// can't compare: slice, map, func
	fmt.Println(reflect.TypeOf(A(1)))
	fmt.Println(reflect.TypeOf(B(1)))
	fmt.Println(max(C(1), C(2)))
	fmt.Println(Point[int64]{x: int64(1), y: int64(2), z: int64(3)}.X())

}

func foo[T IC](a T) T {
	return a
}

func bar[K comparable, V any](k K, m map[K]V) V {
	return m[k]
}

func max[T Numeric](a, b T) T {
	if a > b {
		return a
	}

	return b
}

type IA interface {
	Do1()
}

type IB interface {
	Do2()
}

type IC interface {
	IA
	IB
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint32
}

type A = int64

type B int64

type C B

func (b B) do() {

}

type S[T comparable, V any] map[T]V

func (s S[T, V]) Contains(k T) bool {
	_, ok := s[k]
	return ok
}

type Point[T Numeric] struct {
	x T
	y T
	z T
}

func (p Point[T]) X() T {
	return p.x
}
