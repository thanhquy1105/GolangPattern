package facade

import "fmt"

type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit Notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit Notification")
}
