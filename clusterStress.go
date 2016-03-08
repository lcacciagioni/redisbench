package main

import (
	"strconv"
	"time"

	redigocluster "github.com/chasex/redis-go-cluster"
)

// ClusterStressString is function to stress a redis cluster using simple set
// and strings
func ClusterStressString(c redigocluster.Cluster, minMsgSize, maxMsgSize, numOfMsg int) *StressResult {
	result := initResult()
	for i := 1; i <= numOfMsg; i++ {
		rediskey := "message" + strconv.Itoa(i)
		message := RandRangeString(minMsgSize, maxMsgSize)
		//set
		startSet := time.Now()
		c.Do("SET", rediskey, message)
		duration := time.Since(startSet)
		result.setTotalTime += duration
		if i == 1 || duration < result.minSetTime {
			result.minSetTime = duration
		}
		if i == 1 || duration > result.maxSetTime {
			result.maxSetTime = duration
		}
	}
	for j := 1; j <= numOfMsg; j++ {
		rediskey := "message" + strconv.Itoa(j)
		//get
		startGet := time.Now()
		_, err := redigocluster.String(c.Do("GET", rediskey))
		duration := time.Since(startGet)
		result.getTotalTime += duration
		if j == 1 || duration < result.minGetTime {
			result.minGetTime = duration
		}
		if j == 1 || duration > result.maxSetTime {
			result.maxGetTime = duration
		}
		if err != nil {
			result.lostKeys++
		}
	}
	return result
}

// ClusterStressBytes is function to stress a redis cluster using simple set
// and strings
func ClusterStressBytes(c redigocluster.Cluster, minMsgSize, maxMsgSize, numOfMsg int) *StressResult {
	result := initResult()
	for i := 0; i <= numOfMsg; i++ {
		rediskey := "message" + strconv.Itoa(i)
		message := RandRangeString(minMsgSize, maxMsgSize)
		//set
		startSet := time.Now()
		c.Do("SET", rediskey, message)
		duration := time.Since(startSet)
		result.setTotalTime += duration
		if i == 1 || duration < result.minSetTime {
			result.minSetTime = duration
		}
		if i == 1 || duration > result.maxSetTime {
			result.maxSetTime = duration
		}
	}
	for j := 0; j <= numOfMsg; j++ {
		rediskey := "message" + strconv.Itoa(j)
		//get
		startGet := time.Now()
		_, err := redigocluster.Bytes(c.Do("GET", rediskey))
		duration := time.Since(startGet)
		result.getTotalTime += duration
		if j == 1 || duration < result.minGetTime {
			result.minGetTime = duration
		}
		if j == 1 || duration > result.maxSetTime {
			result.maxGetTime = duration
		}
		if err != nil {
			result.lostKeys++
		}
	}
	return result
}
