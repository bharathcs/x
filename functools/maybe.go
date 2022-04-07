package functools

import "fmt"

type Maybe interface {
}

type Just struct{}

func NewGenericFunc[age int64 | float64](myAge age) {
	fmt.Println(myAge)
}
