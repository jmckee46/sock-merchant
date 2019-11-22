package main

import "testing"

func TestSockMerchant1(t *testing.T) {
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	socks := sockMerchant(int32(len(ar)), ar)

	if socks != 3 {
		t.Errorf("got %d instead of 3", socks)
	}

}

// func sockMerchant(n int32, ar []int32) int32 {
