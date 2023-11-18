package main

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From, To string
	Sum      float64
}

func NewTransaction(from Account, to Account, sum float64) Transaction {
	return Transaction{
		From: from.Name,
		To:   to.Name,
		Sum:  sum,
	}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, t Transaction) Account {
	if t.From == a.Name {
		a.Balance -= t.Sum
	}
	if t.To == a.Name {
		a.Balance += t.Sum
	}
	return a
}

func ArraySum(numbers []int) int {
	sum := func(x, y int) int { return x + y }
	return Reduce[int](numbers, sum, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	accu := func(x, y []int) []int {
		x = append(x, ArraySum(y))
		return x
	}

	return Reduce[[]int](numbersToSum, accu, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {

	oper := func(initialValue []int, arr []int) []int {

		if len(arr) == 0 {
			initialValue = append(initialValue, 0)
		} else {
			initialValue = append(initialValue, ArraySum(arr[1:]))
		}

		return initialValue
	}

	return Reduce[[]int](numbersToSum, oper, []int{})
}

func BalanceFor(transactions []Transaction, name string) float64 {

	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			currentBalance -= t.Sum
		} else if t.To == name {
			currentBalance += t.Sum
		}
		return currentBalance
	}

	return Reduce[Transaction](transactions, adjustBalance, 0)
}

// accumulator (accumator variable, )
func Reduce[T, B any](collection []T, accumulator func(B, T) B, initialValue B) B {
	result := initialValue
	for _, i := range collection {
		result = accumulator(result, i)
	}
	return result
}
