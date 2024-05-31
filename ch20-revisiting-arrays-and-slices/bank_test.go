package main

import "testing"

func TestBadBank(t *testing.T) {
	var(
		john = Account{Name: "John Doe", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		joe = Account{Name: "Joe", Balance: 200}
		
		transactions = []Transaction{
			NewTransaction(chris, john, 100),
			NewTransaction(joe, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(john), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(joe), 175)
}
