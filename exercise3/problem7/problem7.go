package problem7

import "fmt"

type BankAccount struct {
	name    string
	balance int
}

type FedexAccount struct {
	name     string
	packages []string
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func withdrawMoney(amount int, accounts ...interface{}) {
	for _, account := range accounts {
		switch a := account.(type) {
		case *BankAccount:
			if a.balance >= amount {
				a.balance -= amount
			}
		case *KazPostAccount:
			if a.balance >= amount {
				a.balance -= amount
			}
		}
	}
}

func sendPackagesTo(recipient string, accounts ...interface{}) {
	for _, account := range accounts {
		switch a := account.(type) {
		case *FedexAccount:
			message := fmt.Sprintf("%s send package to %s", a.name, recipient)
			a.packages = append(a.packages, message)
		case *KazPostAccount:
			message := fmt.Sprintf("%s send package to %s", a.name, recipient)
			a.packages = append(a.packages, message)
		}
	}
}
