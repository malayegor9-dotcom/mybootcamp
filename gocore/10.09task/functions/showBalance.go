package functions

import (
	"bootcamp/gocore/10.09task/methods"
	"fmt"
)

func ShowBalance(accounts map[int]methods.Account, id int) {
	account := accounts[id]
	fmt.Printf("Баланс счета: %.2f\n", account.Balance)
}
