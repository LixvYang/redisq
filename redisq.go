package redisq

// Redisq is a queue
type Redisq struct {
	Addr           string
	MaxConcurrency uint64
}

type IOption interface {
	apply(*Redisq)
}

type funcOption struct {
	f func(*Redisq)
}

func (fo *funcOption) apply(r *Redisq) {
	fo.f(r)
}

func newFuncOption(f func(*Redisq)) IOption {
	return &funcOption{
		f: f,
	}
}

func WithMaxConcurrency(c uint64) IOption {
	return newFuncOption(func(r *Redisq) {
		r.MaxConcurrency = c
	})
}

type OptionFunc func(*Redisq)

func NewRedisq(addr string, opts ...OptionFunc) *Redisq {
	o := &Redisq{Addr: addr}

	for _, opt := range opts {
		opt(o)
	}

	return o
}
