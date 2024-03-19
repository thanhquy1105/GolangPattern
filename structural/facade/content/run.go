package facade

import (
	"fmt"
	"log"
)

func Run() {
	fmt.Println()
	walletFacade := NewWalletFacade("abc", 123)
	fmt.Println()
	err := walletFacade.AddMoneyToWallet("abc", 123, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletFacade.deductMoneyToWallet("abc", 123, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

}
