package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	sockMap := make(map[int32]int32)
	var total int32

	for i := int32(0); i < n; i++ {
		sockMap[ar[i]]++
	}

	for key, value := range sockMap {
		fmt.Println("key:", key, "value:", value)
		total += value / 2
	}

	return total
}

func main() {

}
