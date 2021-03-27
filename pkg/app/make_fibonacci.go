// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

package app

import "fmt"

func Fibint(n int,

) int {
	if n <= 2 {
		fmt.Println("Trivial case, n = ", n)
		return 1
	}
	return Fibint(n-1) +
		Fibint(n-2)
}
func Fibint32(n int32,

) int32 {
	if n <= 2 {
		fmt.Println("Trivial case, n = ", n)
		return 1
	}
	return Fibint32(n-1) +
		Fibint32(n-2)
}
func Fibint64(n int64,

) int64 {
	if n <= 2 {
		fmt.Println("Trivial case, n = ", n)
		return 1
	}
	return Fibint64(n-1) +
		Fibint64(n-2)
}
func Fibuint(n uint,

) uint {
	if n <= 2 {
		fmt.Println("Trivial case, n = ", n)
		return 1
	}
	return Fibuint(n-1) +
		Fibuint(n-2)
}
func Fibuint32(n uint32,

) uint32 {
	if n <= 2 {
		fmt.Println("Trivial case, n = ", n)
		return 1
	}
	return Fibuint32(n-1) +
		Fibuint32(n-2)
}
func Fibuint64(n uint64,

) uint64 {
	if n <= 2 {
		fmt.Println("Trivial case, n = ", n)
		return 1
	}
	return Fibuint64(n-1) +
		Fibuint64(n-2)
}
