package redisq

import (
	"github.com/redis/go-redis/v9"
)

type OptionFunc func(*Redisq)

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

func WithHandleFunc(f func() error) IOption {
	return newFuncOption(func(r *Redisq) {
		r.handleFunc = f
	})
}

func WithMaxConcurrency(c uint64) IOption {
	return newFuncOption(func(r *Redisq) {
		r.MaxConcurrency = c
	})
}

func WithAddTaskScript(addTaskScript string) IOption {
	return newFuncOption(func(r *Redisq) {
		r.addTaskScript = redis.NewScript(addTaskScript)
	})
}

func WithGetTaskScript(getTaskScript string) IOption {
	return newFuncOption(func(r *Redisq) {
		r.addTaskScript = redis.NewScript(getTaskScript)
	})
}

func WithMaxRetry(maxRetry uint64) IOption {
	return newFuncOption(func(r *Redisq) {
		r.MaxRetry = maxRetry
	})
}
