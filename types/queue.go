package types

// SlotQueue is a simple type alias for a (buffered) channel of block heights.
type SlotQueue chan uint64

func NewQueue(size int) SlotQueue {
	return make(chan uint64, size)
}

// BankTaskQueue is a simple type alias for a (buffered) channel of bank tasks.
type BankTaskQueue chan func()

func NewBankTaskQueue(size int) BankTaskQueue {
	return make(chan func(), size)
}
