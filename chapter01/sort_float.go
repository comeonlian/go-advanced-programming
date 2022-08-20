package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}

func SortFloat64FastV1(a []float64) {
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	sort.Ints(c)
}

func main() {
	SortFloat64FastV1(a)
	fmt.Printf("%v", a)
}

