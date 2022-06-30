package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Ethereum(10))

		assertBalance(t, wallet, Ethereum(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Ethereum(20)}
		err := wallet.Withdraw(Ethereum(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Ethereum(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Ethereum(20)}
		err := wallet.Withdraw(Ethereum(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Ethereum(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Ethereum) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
