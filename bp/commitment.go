package bp

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

/*
	Pedersen Commitment

Given a values, we commit the value with different generators
for each element and for each randomness.
*/
func PCommit(value *big.Int, r *big.Int) ECPoint {

	commitment := EC.Zero()

	// r, err := rand.Int(rand.Reader, EC.N)
	// check(err)

	// modValue := new(big.Int).Mod(value, EC.N)

	// mG, rH
	// lhsX, lhsY := EC.C.ScalarMult(EC.BPG[0].X, EC.BPG[0].Y, modValue.Bytes())
	// rhsX, rhsY := EC.C.ScalarMult(EC.BPH[0].X, EC.BPH[0].Y, r.Bytes())
	// commitment = commitment.Add(ECPoint{lhsX, lhsY}).Add(ECPoint{rhsX, rhsY})

	commitment = EC.G.Mult(value).Add(EC.H.Mult(r))

	return commitment
}

/*
Vector Pedersen Commitment

Given an array of values, we commit the array with different generators
for each element and for each randomness.
*/
func VectorPCommit(value []*big.Int) (ECPoint, []*big.Int) {
	R := make([]*big.Int, EC.V)

	commitment := EC.Zero()

	for i := 0; i < EC.V; i++ {
		fmt.Println("i", i)
		fmt.Println("v", EC.V)
		fmt.Println("EN N: ", EC.N)

		r, err := rand.Int(rand.Reader, EC.N)
		check(err)

		R[i] = r

		modValue := new(big.Int).Mod(value[i], EC.N)

		// mG, rH
		lhsX, lhsY := EC.C.ScalarMult(EC.BPG[i].X, EC.BPG[i].Y, modValue.Bytes())
		rhsX, rhsY := EC.C.ScalarMult(EC.BPH[i].X, EC.BPH[i].Y, r.Bytes())

		commitment = commitment.Add(ECPoint{lhsX, lhsY}).Add(ECPoint{rhsX, rhsY})
	}
	fmt.Println("11")

	return commitment, R
}

/*
Two Vector P Commit

Given an array of values, we commit the array with different generators
for each element and for each randomness.
*/
func TwoVectorPCommit(a []*big.Int, b []*big.Int) ECPoint {
	if len(a) != len(b) {
		fmt.Println("TwoVectorPCommit: Uh oh! Arrays not of the same length")
		fmt.Printf("len(a): %d\n", len(a))
		fmt.Printf("len(b): %d\n", len(b))
	}

	commitment := EC.Zero()

	for i := 0; i < EC.V; i++ {
		commitment = commitment.Add(EC.BPG[i].Mult(a[i])).Add(EC.BPH[i].Mult(b[i]))
	}

	return commitment
}

/*
Vector Pedersen Commitment with Gens

Given an array of values, we commit the array with different generators
for each element and for each randomness.

We also pass in the Generators we want to use
*/
func TwoVectorPCommitWithGens(G, H []ECPoint, a, b []*big.Int) ECPoint {
	if len(G) != len(H) || len(G) != len(a) || len(a) != len(b) {
		fmt.Println("TwoVectorPCommitWithGens: Uh oh! Arrays not of the same length")
		fmt.Printf("len(G): %d\n", len(G))
		fmt.Printf("len(H): %d\n", len(H))
		fmt.Printf("len(a): %d\n", len(a))
		fmt.Printf("len(b): %d\n", len(b))
	}

	commitment := EC.Zero()

	for i := 0; i < len(G); i++ {
		modA := new(big.Int).Mod(a[i], EC.N)
		modB := new(big.Int).Mod(b[i], EC.N)

		commitment = commitment.Add(G[i].Mult(modA)).Add(H[i].Mult(modB))
	}

	return commitment
}
