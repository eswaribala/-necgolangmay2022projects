package main

import (
	"fmt"
	"github.com/teivah/bitvector"
)

func main() {

	var bv1 bitvector.Ascii
	var bv2 bitvector.Ascii
	bv1 = bv1.Set(10, true)
	bv2 = bv2.Set(15, true)
	bv3 := bv1.Xor(bv2)
	fmt.Println(bv3)

}
