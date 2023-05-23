package scripts

const (
	// 获取到任务时就删除任务
	// ZRANGEBYSCORE KEYS[1] 0 ARGV[1] LIMIT 0 1
	GetTaskLua = `
		local queueKey = KEYS[1]  
		local taskScore = ARGV[1]  
		local result={}
		local items = redis.call("ZRANGEBYSCORE", queueKey, 0, taskScore, "LIMIT", "0", "1")
		
		if #items > 0 then
			result[1] = items[1] 
			redis.call("ZREM", queueKey, items[1])
		end
		
		return result
	`

	// ZADD KEYS[1] AGRV[2] KEYS[2]
	// SET KEYS[2] ARGV[1]
	AddTaskLua = `  
		local queueKey = KEYS[1]
		local taskKey = KEYS[2]
		local taskData = ARGV[1]
		local taskScore = ARGV[2]
		redis.call("ZADD", queueKey, taskScore, taskKey)	
		redis.call("SET", taskKey, taskData)
	`
)
