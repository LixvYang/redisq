package redisq

const (
	defaultConcurrency = 30
)

// Redisq is a queue
type Redisq struct {
	Addr           string
	MaxConcurrency uint64
}

func NewRedisq(addr string, opts ...IOption) *Redisq {
	o := &Redisq{Addr: addr, MaxConcurrency: defaultConcurrency}

	for _, opt := range opts {
		opt.apply(o)
	}

	return o
}
