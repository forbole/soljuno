package pool

type Pool interface {
	DoAsync(fun func() error) chan error
	IsFree() bool
	IsStopped() bool
}
