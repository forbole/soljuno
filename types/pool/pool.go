package pool

type Pool interface {
	// DoAsync insert a task into the pool then execute the given task asynchronously
	// It returns the error channel to be able checked if the task returns error
	DoAsync(fun func() error) chan error

	// IsFree returns boolean showing if the pool is free or not
	IsFree() bool

	// IsEmpty returns boolean showing if the pool is empty or not
	IsEmpty() bool
}
