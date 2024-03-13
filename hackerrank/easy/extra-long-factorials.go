package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int32) {
	var fact big.Int
	fact.MulRange(1, int64(n))
	fmt.Println(&fact)

}
