package main

import (
	"bitset"
	"fmt"
)

func main() {
	fmt.Print(21321321)
	var a = bitset.Bitset{}
	a.SetLen(64 * 30)
	a.Print_()
	a.SetBit(0)
	a.Print_()
	a.SetLen(0)
	a.Print_()

	a.SetLen(64 * 70)
	a.Print_()
	a.SetBit(64 * 30)
	a.Print_()
	a.SetBit(64*30 + 1)
	a.Print_()
	a.SetBit(64*30 - 1)
	a.Print_()
	a.SetLen(64 * 30)
	a.Print_()

	a.SetLen(65)
	a.SetBit(62)
	a.SetBit(61)
	a.SetBit(0)
	a.SetBit(1)
	a.Print_()
	a.SetLen(64)
	a.Print_()
	a.SetLen(63)
	a.Print_()
	a.SetLen(62)
	a.Print_()

	a.Print_()

}
