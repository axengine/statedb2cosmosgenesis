package main

import (
	"fmt"
	"testing"
)

func TestHexAddr2CosmosAddr(t *testing.T) {
	globalInitCosmosConfig("evmos")
	addr := evmAddressToCosmosAddress("0x0F508F143E77b39F8e20DD9d2C1e515f0f527D9F")
	fmt.Println(addr)
}
