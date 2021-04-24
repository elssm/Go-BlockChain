package BLC

import (
	"fmt"
	"testing"
)

func TestWallets_CreateWallet(t *testing.T) {
	wallets := NewWallets()
	wallets.CreateWallet()
	fmt.Printf("wallets : %v\n",wallets.Wallets)
}