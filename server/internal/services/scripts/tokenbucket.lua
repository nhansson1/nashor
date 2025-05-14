local key = KEYS[1]

local fastBucketTokens = tonumber(redis.call("HGET", key, "fast_bucket"))
local fastRefill = tonumber(redis.call("HGET", key, "fast_bucket_last_refill"))
local fastInterval = tonumber(redis.call("HGET", key, "fast_bucket_interval"))
local fastBucketMax = tonumber(redis.call("HGET", key, "fast_bucket_max"))

local slowBucketTokens = tonumber(redis.call("HGET", key, "slow_bucket"))
local slowRefill = tonumber(redis.call("HGET", key, "slow_bucket_last_refill"))
local slowInterval = tonumber(redis.call("HGET", key, "slow_bucket_interval"))
local slowBucketMax = tonumber(redis.call("HGET", key, "slow_bucket_max"))

local now = tonumber(redis.call("TIME")[1])

if fastRefill == 0 and slowRefill == 0 then
    fastRefill = now
    slowRefill = now
    redis.call("HSET", key, "fast_bucket_last_refill", fastRefill)
    redis.call("HSET", key, "slow_bucket_last_refill", slowRefill)
end

if now > fastRefill + fastInterval then
    fastRefill = now
    fastBucketTokens = fastBucketMax

    redis.call("HSET", key,
        "fast_bucket", fastBucketMax,
        "fast_bucket_last_refill", fastRefill
    )
end

if now > slowRefill + slowInterval then
    slowRefill = now
    slowBucketTokens = slowBucketMax

    redis.call("HSET", key,
        "slow_bucket", slowBucketMax,
        "slow_bucket_last_refill", slowRefill
    )
end

if slowBucketTokens <= 0 then
    return cjson.encode({
        status = "slow_limit",
        fast = fastBucketTokens,
        slow = slowBucketTokens
    })
end

if fastBucketTokens <= 0 then
    return cjson.encode({
        status = "fast_limit",
        fast = fastBucketTokens,
        slow = slowBucketTokens
    })
end

redis.call("HINCRBY", key, "fast_bucket", -1)
redis.call("HINCRBY", key, "slow_bucket", -1)
fastBucketTokens = fastBucketTokens - 1
slowBucketTokens = slowBucketTokens - 1

return cjson.encode({
    status = "ok",
    fast = fastBucketTokens,
    slow = slowBucketTokens
})
