package main

import (
	"strconv"
	"time"

	redigo "github.com/garyburd/redigo/redis"
)

// NodeStressString stress tests a single node with strings
func NodeStressString(c redigo.Conn, numOfMsg, minMsgSize, maxMsgSize int) *StressResult {
	result := initResult()
	// this first part will gonna set strings as values
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
		_, err := redigo.String(c.Do("GET", rediskey))
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

// NodeStressBytes stress tests a single node with bytes
func NodeStressBytes(c redigo.Conn, numOfMsg, minMsgSize, maxMsgSize int) *StressResult {
	result := initResult()
	// this first part will gonna set strings as values
	for i := 0; i <= numOfMsg; i++ {
		rediskey := "message" + strconv.Itoa(i)
		message := RandRangeBytes(minMsgSize, maxMsgSize)
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
		_, err := redigo.Bytes(c.Do("GET", rediskey))
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
