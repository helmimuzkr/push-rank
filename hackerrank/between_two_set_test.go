package hackerrank

import (
	"fmt"
	"testing"
)

// Function Description
// Complete the getTotalX function in the editor below. It should return the number of integers that are betwen the sets.
// getTotalX has the following parameter(s):
// - int a[n]: an array of integers
// - int b[m]: an array of integers

// Returns
// int: the number of integers that are between the sets

// Input Format
// The first line contains two space-separated integers, n and m, the number of elements in arrays a and b.
// The second line contains n  distinct space-separated integers a[i] where 0 <= i < n.
// The third line contains m distinct space-separated integers b[j] where 0 <= j < m.

// Example
// Input
// 2,  3		=> Input 2 integer, jumlah element baris 2(Deret A) dan baris 3(Deret B)
// 2, 4			=> input 2 integer, 2 bilangan bulat Deret A
// 16. 32, 96	=> input 3 integer, 3 bilangan bulat Deret B
// Output
// 3

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var (
		x    int32
		temp []int32 // Deret Temporary, menampung bilangan bulat sementara untuk dibandingkan
	)

	// Langkah pertama mencari deret sementara, yaitu bilangan bulat yang memiliki faktor dari Deret A
	// j adalah bilangan bulat yang dicari
	// dengan batas awal elemen pertama Deret A
	// dan batas akhir elemen pertama dari Deret B
	for j := a[0]; j <= b[0]; j++ {
		match := 0
		for i := range a {
			// Jika bilangan j % bilangan Deret a[i] == 0, maka a adalah faktor dari j.
			// Lalu, append ke Temporary variable
			if j%a[i] == 0 {
				match++
			}
		}

		// Kalau bilangan j match sebanyak jumlah bilangan pada Deret A, maka Deret A adalah faktor dari j
		if match == len(a) {
			temp = append(temp, j)
		}
	}

	fmt.Println("bilangan temporary", temp) // Melihat Deret temporary

	// Langkah terakhir mencocokan apakah Deret Temporary dengan Deret B
	// Apakah faktor Deret B adalah bilangan pada Deret Temporary?
	for _, bilTemp := range temp {
		match := 0
		for _, deretB := range b {
			// Jika bilangan temp adalah faktor dari Deret B, maka match++
			if deretB%bilTemp == 0 {
				match++
			}
		}

		// Kalau bilangan Temporary yang match sebanyak jumlah Deret B
		// atau bilangan Temporary adalah faktor dari Deret B,
		// maka sudah dipastikan bilangan Temprary tersebut merupakan bilangan dari Deret Baru (diantara Deret A dan Deret B)
		if match == len(b) {
			x++
		}
	}

	return x
}

func TestBetweenTwoSet(t *testing.T) {
	a := []int32{2, 4}
	b := []int32{16, 32, 96}

	x := getTotalX(a, b)
	fmt.Println(x) // 3
}

// Explain
// Pada kasus ini mencari Deret Baru, diantara Deret A dan Deret B
// Langkah
// 1. Bilangan bulat berapa yang mempunyai faktor dari Deret A, tetapi tidak melebihi nilai awal Deret B.
//		Berikut hasil pengecekan
//			- 4 % 2 == 0 	&&	 4 % 4 == 0
// 			- 8 % 2 == 0 	&&	 8 % 4 == 0
// 			- 12 % 2 == 0 	&&	 12 % 4 == 0
// 			- 16 % 2 == 0 	&&	 16 % 4 == 0
// 3. Hasil dari proses pencarian bilangan bulat pada nomor 1 adalah 4, 8, 12, dan 16. berikutnya disimpan ke dalam Variabel Temporary
// 4. Sekarang sudah mendaptkan Deret Temporary yang faktornya adalah Deret A.
// 5. Karena mencari Deret baru diantara Deret A dan B, berarti bilangan bulat Temporary adalah faktor dari Deret B.
// 		Berikut hasil pengecekan:
//			- 16 % 4 == 0	&&	16 % 8 == 0		&&	16 % 12 == 4	&&	16 % 16 == 0
//			- 32 % 4 == 0	&&	32 % 8 == 0		&&	32 % 12 == 8	&&	32 % 16 == 0
//			- 96 % 4 == 0	&&	96 % 8 == 0		&&	96 % 12 == 0	&&	96 % 16 == 0
// 6. Bisa dilihat hasil nomor 5, bahwa hanya 3 bilangan bulat dari Temporary yang merupakan faktor dari Deret B yaitu 4, 8, dan 16.
// 7. Maka Deret baru adalah 4, 8 , dan 16
// 8. X(jumlah Deret baru) adalah 3
