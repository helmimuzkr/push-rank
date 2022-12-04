package hackerrank

import (
	"fmt"
	"testing"
)

// Clue
// Given an array of integers and a positive integer k,
// determine the number of (i,j) pairs where i < j and ar[i] + ar[j] is divisible by k.

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var x int32 // Jumlah operasi penjumlahan 2 elemen(bilangan bulat) ar yang masuk kriteria

	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			// Jika i < j dan hasil penjumlahan ar[i]+ar[j] habis dibagi k,
			// maka hasil dari ar[i] dan ar[j] sudah memenuhi kriteria x
			if i < j && ((ar[i]+ar[j])%k == 0) {
				x++
			}
		}
	}

	return x
}

func TestDivisibleSumPairs(t *testing.T) {
	var (
		n   int32 = 6
		k   int32 = 3
		arr       = []int32{1, 3, 2, 6, 1, 2}
	)
	fmt.Println(divisibleSumPairs(n, k, arr)) // 5
}
