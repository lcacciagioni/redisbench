package main

import (
	"fmt"
	"log"
	"time"

	redigocluster "github.com/chasex/redis-go-cluster"
	redigo "github.com/garyburd/redigo/redis"
)

// StressCluster is a simple function to stress test a redis cluster
func StressCluster(hosts []string, minMsgSize, maxMsgSize, numOfMsg int) {
	c, err := redigocluster.NewCluster(&redigocluster.Options{
		StartNodes:   hosts,
		ConnTimeout:  50 * time.Millisecond,
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive:    16,
		AliveTime:    60 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("=====STRINGS=====")
	printResult(ClusterStressString(*c, minMsgSize, maxMsgSize, numOfMsg), numOfMsg)
	fmt.Println("=====BYTES=====")
	printResult(ClusterStressBytes(*c, minMsgSize, maxMsgSize, numOfMsg), numOfMsg)
}

// StressNode is fuction that executes a series of stress tests in a sigle node
// of redis
func StressNode(host string, minMsgSize, maxMsgSize, numOfMsg int) {
	c, err := redigo.Dial("tcp", host)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()
	fmt.Println("=====STRINGS=====")
	printResult(NodeStressString(c, numOfMsg, minMsgSize, maxMsgSize), numOfMsg)
	fmt.Println("=====BYTES=====")
	printResult(NodeStressBytes(c, numOfMsg, minMsgSize, maxMsgSize), numOfMsg)
}
