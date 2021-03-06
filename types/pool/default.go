package pool

import "github.com/panjf2000/ants/v2"

func NewDefaultPool(poolSize int) (Pool, error) {
	pool, err := ants.NewPool(poolSize)
	return &defaultPool{pool: pool}, err
}

type defaultPool struct {
	pool *ants.Pool
}

func (p *defaultPool) DoAsync(fun func() error) (chan error, error) {
	errCh := make(chan error, 1)
	err := p.pool.Submit(func() {
		errCh <- fun()
	})
	return errCh, err
}

func (p *defaultPool) IsFree() bool {
	return p.pool.Free() != 0
}

func (p *defaultPool) IsEmpty() bool {
	return p.pool.Running() == 0
}
