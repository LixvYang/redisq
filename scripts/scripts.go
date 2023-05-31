package scripts

const (
	GetTaskLua = `--KEYS 
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

	AddTaskLua = `--KEYS
		local queueKey = KEYS[1]
		local taskKey = KEYS[2]
		local taskData = ARGV[1]
		local taskScore = ARGV[2]
		redis.call("ZADD", queueKey, taskScore, taskKey)	
		redis.call("SET", taskKey, taskData)
	`


)
