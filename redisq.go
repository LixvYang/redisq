package redisq

import (
	"github.com/lixvyang/redisq/constant"
	"github.com/lixvyang/redisq/scripts"
	"github.com/redis/go-redis/v9"
)

// redis queue
type Redisq struct {
	RedisOpt       *redis.Options
	MaxConcurrency uint64
	MaxRetry       uint64

	rdb           *redis.Client
	addTaskScript *redis.Script
	getTaskScript *redis.Script
}

func NewRedisq(redisOpt *redis.Options, opts ...IOption) *Redisq {
	o := &Redisq{
		RedisOpt:       redisOpt,
		MaxConcurrency: constant.DefaultConcurrency,
		MaxRetry:       constant.DefaultMaxRetry,
		addTaskScript:  redis.NewScript(scripts.AddTaskLua),
		getTaskScript:  redis.NewScript(scripts.GetTaskLua),
	}

	for _, opt := range opts {
		opt.apply(o)
	}

	o.initRedis()
	return o
}

func (rq *Redisq) initRedis() {
	rq.rdb = redis.NewClient(rq.RedisOpt)
}

func (rq *Redisq) AddTask() {
	
}

func (rq *Redisq) GetTask() {

}
