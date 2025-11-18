package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

const (
	MaxTransactionAmount = 10000.0 // Example limit for deposits/withdrawals
)

func main() {
	checkType(&Animals{
		Speed: 100,
		HP:    100,
	})
}

func checkType(m Moving) {
	animals, ok := m.(*Animals)
	if ok {
		animals.HP = 100
		fmt.Printf("%#v", animals)
		return
	}

	human, ok := m.(*Human)
	if ok {
		human.HP = 100
		fmt.Printf("%#v", animals)
		return
	}
}

type Moving interface {
	Jump(origin int) int
	MoveX(origin int) int
	MoveY(origin int) int
}

type Animals struct {
	Speed int
	HP    int
}

type Human struct {
	Power int
	HP    int
}

func (h *Human) Jump(from int) int {
	return 0
}

func (h *Human) MoveX(from int) int {
	return 0
}

func (h *Human) MoveY(from int) int {
	return 0
}

func (a *Animals) Jump(from int) int {
	return 0
}

func (a *Animals) MoveX(from int) int {
	return 0
}

func (a *Animals) MoveY(from int) int {
	return 0
}

type BankAccount struct {
	ID         string
	Owner      string
	Balance    float64
	MinBalance float64
}

func NewBankAccount(id, owner string, initialBalance, minBalance float64) (*BankAccount, error) {
	negativeError := &NegativeAmountError{}
	insufficientError := &InsufficientFundsError{}
	accountError := &AccountError{}
	if id == "" || owner == "" {
		accountError.ID = id
		accountError.Owner = owner
		return nil, accountError
	}
	if minBalance < 0 || initialBalance < 0 {
		negativeError.Amount = []float64{initialBalance, minBalance}
		return nil, negativeError
	}
	if initialBalance < minBalance {
		insufficientError.AvailableAmount = minBalance
		insufficientError.RequestedAmount = initialBalance
		return nil, insufficientError
	}
	bankAccount := &BankAccount{
		ID:         id,
		Owner:      owner,
		Balance:    initialBalance,
		MinBalance: minBalance,
	}
	return bankAccount, nil
}

func (ba *BankAccount) Deposit(amount float64) error {

	mu.Lock()
	negativeAmountErr := &NegativeAmountError{}
	exceedsLimitErr := &ExceedsLimitError{Limit: MaxTransactionAmount}
	if amount < 0 {
		negativeAmountErr.Amount = []float64{amount}
		return negativeAmountErr
	}

	if amount > MaxTransactionAmount {
		exceedsLimitErr.RequestedAmount = amount
		return exceedsLimitErr
	}

	ba.Balance += amount
	mu.Unlock()
	return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
	mu.Lock()
	negativeAmountErr := &NegativeAmountError{}
	exceedsLimitErr := &ExceedsLimitError{Limit: MaxTransactionAmount}
	insufficientFundsErr := &InsufficientFundsError{}
	if amount < 0 {
		negativeAmountErr.Amount = []float64{amount}
		return negativeAmountErr
	}

	if amount > MaxTransactionAmount {
		exceedsLimitErr.RequestedAmount = amount
		return exceedsLimitErr
	}

	currentBalance := ba.Balance

	if currentBalance < amount {
		insufficientFundsErr.AvailableAmount = currentBalance
		insufficientFundsErr.RequestedAmount = amount
		return insufficientFundsErr
	}

	decreased := ba.Balance - amount

	if decreased < ba.MinBalance {
		insufficientFundsErr.AvailableAmount = ba.MinBalance
		insufficientFundsErr.RequestedAmount = decreased
		return insufficientFundsErr
	}

	ba.Balance = decreased
	mu.Unlock()
	return nil
}

func (ba *BankAccount) Transfer(amount float64, target *BankAccount) error {
	if err := ba.Withdraw(amount); err != nil {
		return err
	}
	if err := target.Deposit(amount); err != nil {
		return err
	}

	return nil
}

type AccountError struct {
	Owner string
	ID    string
}

type InsufficientFundsError struct {
	AvailableAmount float64
	RequestedAmount float64
}

type NegativeAmountError struct {
	Amount []float64
}

type ExceedsLimitError struct {
	Limit           float64
	RequestedAmount float64
}

func (e *AccountError) Error() string {
	return fmt.Sprintf("account error: Owner=%v, ID=%v\n", e.Owner, e.ID)
}
func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient balance funds error: RequestedAmount=%v, AvailableAmount=%v\n", e.RequestedAmount, e.AvailableAmount)
}
func (e *NegativeAmountError) Error() string {
	return fmt.Sprintf("negative amount error: Amount=%v\n", e.Amount)
}
func (e *ExceedsLimitError) Error() string {
	return fmt.Sprintf("exceeds limit error: Limit=%v, RequestedAmount=%v\n", e.Limit, e.RequestedAmount)
}
