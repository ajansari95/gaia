package main

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	a := "cosmos1nz6g54huekp4qs45shttf79czvg93nn7plarr9"
	addr, err := sdk.AccAddressFromBech32(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)
}
