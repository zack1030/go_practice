package main

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

// Result : prevent optimization
var Result uint64

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	x = (x * h01) >> 56
	Result = x // FIX: change exported var to prevent optimization
	return x
}

func main() {
	popcnt(10)
}
