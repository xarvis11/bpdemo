package main

import (
	"fmt"
	"math/big"

	"xaults.com/bpdemo/bp"
)

func main() {
	// fmt.Println("Hello, World!")
	// bp.test()
	// v := make([]*big.Int, 4)

	// bp.EC = bp.NewECPrimeGroupKey(3)
	// v := make([]*big.Int, 3)
	// // for j := range v {
	// // 	v[j] = big.NewInt(2)
	// // }
	// v[0] = big.NewInt(2)
	// v[1] = big.NewInt(1)
	// v[2] = big.NewInt(1)
	// // v[3] = big.NewInt(1)

	// _, r := bp.VectorPCommit(v)

	// // fmt.Printf("curve:", output)
	// fmt.Println("pc:", r)

	// // bp.EC = bp.NewECPrimeGroupKey(64)
	// result := new(big.Int).Sub(r[0], new(big.Int).Add(r[1], r[2]))
	// result2 := r[0].Sub(r[0], r[0].Add(r[1], r[2]))

	// if r[0].Cmp(new(big.Int).Add(r[1], r[2])) != 0 {
	// 	fmt.Println("Inputs not equal to outputs")
	// }

	// fmt.Println("#### sum check: ", result, result2)

	bp.EC = bp.NewECPrimeGroupKey(64)
	if bp.RPVerify(bp.RPProve(big.NewInt(6))) {
		fmt.Println("Range Proof Verification works")
	} else {
		fmt.Println("*****Range Proof FAILURE")
	}

}
