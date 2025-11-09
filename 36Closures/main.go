package main

import "fmt"

// activate100RsGiftCard credits a 100 Rs giftcard to the user and returns a function.
func activate100RsGiftCard() func(int) int {
	amount := 100
	// debitFunc takes some amount from the user an subtracts it from the giftcard amount the user has
	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
	return debitFunc
}

func main() {
	dF := activate100RsGiftCard()
	fmt.Println("Bought chocolate for 25 Rs, the giftcard amout left is", dF(25))
	fmt.Println("Bought chips for 15 Rs, the giftcard amout left is", dF(15))
	fmt.Println("Bought toffees for 10 Rs, the giftcard amout left is", dF(10))
}