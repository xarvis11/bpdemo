package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"xaults.com/bpdemo/bp"
)

func main() {

	bp.EC = bp.NewECPrimeGroupKey(8)

	// Define input values and output values
	inputValues := []*big.Int{big.NewInt(15), big.NewInt(15)} // Example: two inputs, each 15
	outputValues := []*big.Int{big.NewInt(29), big.NewInt(1)} // Example: two outputs, 29 and 1

	sumInputValues := big.NewInt(0)
	sumOutputValues := big.NewInt(0)

	inputCommitments := make([]bp.ECPoint, len(inputValues))
	inputBlindings := make([]*big.Int, len(inputValues))
	outputCommitments := make([]bp.ECPoint, len(outputValues))
	outputBlindings := make([]*big.Int, len(outputValues))

	// calculate input blindings
	for i, _ := range inputValues {

		r, err := rand.Int(rand.Reader, bp.EC.N)
		if err != nil {
			panic(err)
		}
		inputBlindings[i] = r
	}
	fmt.Printf("\ninput blindings: %s", inputBlindings)

	// calculate output blindings
	for i, _ := range outputValues {

		r, err := rand.Int(rand.Reader, bp.EC.N)
		if err != nil {
			panic(err)
		}
		outputBlindings[i] = r
	}
	fmt.Printf("\noutput blindings:  %s", outputBlindings)

	// Adjust the blinding factors to ensure the sum of inputs equals the sum of outputs
	inputBlindingSum := big.NewInt(0)
	for _, r := range inputBlindings {
		inputBlindingSum.Add(inputBlindingSum, r)
	}
	outputBlindingSum := big.NewInt(0)
	for _, r := range outputBlindings {
		outputBlindingSum.Add(outputBlindingSum, r)
	}
	// Adjust the last output blinding factor to balance the sums
	diff := new(big.Int).Sub(inputBlindingSum, outputBlindingSum)
	outputBlindings[len(outputBlindings)-1].Add(outputBlindings[len(outputBlindings)-1], diff)
	// outputBlindings[len(outputBlindings)-1].Mod(outputBlindings[len(outputBlindings)-1], bp.EC.N)

	// calculate input commitments
	for i, input := range inputValues {

		comm := bp.PCommit(input, inputBlindings[i])
		inputCommitments[i] = comm
		sumInputValues.Add(sumInputValues, input)
	}
	fmt.Printf("\ninput commitments:  %s", inputCommitments)
	fmt.Printf("\nsum of input values:  %s", sumInputValues)

	// calculate output commitments
	for i, output := range outputValues {

		comm := bp.PCommit(output, inputBlindings[i])
		outputCommitments[i] = comm
		sumOutputValues.Add(sumOutputValues, output)
	}
	fmt.Printf("\noutput commitments:  %s", outputCommitments)
	fmt.Printf("\nsum of output values:  %s", sumOutputValues)

	// Sum input and output commitments
	sumInputCommitments := bp.EC.Zero()
	for _, inputCommitment := range inputCommitments {
		sumInputCommitments = inputCommitment.Add(sumInputCommitments)
	}
	fmt.Printf("\nsum input commitments:  %s", sumInputCommitments)

	sumOutoutCommitments := bp.EC.Zero()
	for _, outputCommitment := range outputCommitments {
		sumOutoutCommitments = sumOutoutCommitments.Add(outputCommitment)
	}
	fmt.Printf("\nsum output commitments:  %s", sumOutoutCommitments)

	// Verify if the summed commitments for inputs equal the summed commitments for outputs
	if sumInputCommitments.Equal(sumOutoutCommitments) {
		fmt.Println()
		fmt.Println("#################")
		fmt.Println("\nThe sum of the input commitments equals the sum of the output commitments.")
		fmt.Println("#################")
	} else {
		fmt.Println("#################")
		fmt.Println("\nThe sum of the input commitments does NOT equal the sum of the output commitments.")
		fmt.Println("#################")
	}

	// Range Proof

	// bp.EC = bp.NewECPrimeGroupKey(1)
	// value := new(big.Int)
	// value.SetString("255", 10)

	// v := make([]*big.Int, 1)
	// v[0] = value
	// curve, r := bp.VectorPCommit(v)

	// fmt.Printf("curve:", curve)
	// fmt.Println("pc:", r)

	// bp.EC = bp.NewECPrimeGroupKey(8)
	// proof := bp.CommRPProve(sumInputValues, inputBlindingSum, sumInputCommitments)
	proof := bp.CommRPProve(inputValues[0], inputBlindings[0], inputCommitments[0])
	fmt.Println("Prrof:", proof)
	if bp.RPVerify(proof) {
		fmt.Println("#################")
		fmt.Println("Range Proof Verification works")
		fmt.Println("#################")
	} else {
		fmt.Println("#################")
		fmt.Println("*****Range Proof FAILURE")
		fmt.Println("#################")
	}

}
