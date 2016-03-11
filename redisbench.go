package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// StressResult is used to return common values to print later
type StressResult struct {
	setTotalTime, getTotalTime, maxSetTime, minSetTime, maxGetTime, minGetTime time.Duration
	lostKeys                                                                   uint
}

func initResult() *StressResult {
	result := &StressResult{
		setTotalTime: time.Duration(0),
		getTotalTime: time.Duration(0),
		maxSetTime:   time.Duration(0),
		minSetTime:   time.Duration(0),
		maxGetTime:   time.Duration(0),
		minGetTime:   time.Duration(0),
		lostKeys:     uint(0),
	}
	return result
}

func printResult(result *StressResult, numOfMsg int) {
	fmt.Println("SET")
	fmt.Println("Total:", result.setTotalTime)
	fmt.Println("Min:", result.minSetTime)
	fmt.Println("Max:", result.maxSetTime)
	fmt.Println("Avg:", result.setTotalTime/time.Duration(numOfMsg))
	fmt.Println("GET")
	fmt.Println("Total:", result.getTotalTime)
	fmt.Println("Min:", result.minGetTime)
	fmt.Println("Max:", result.maxGetTime)
	fmt.Println("Keys Missed / Errors:", result.lostKeys)
	fmt.Println("Avg:", result.setTotalTime/time.Duration(numOfMsg))
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: redisbench [options]\n\n")
		flag.PrintDefaults()
	}
	minMsgSizePtr := flag.Int("minMsgSize", 100, "The minimun size of the massages to send in bytes")
	maxMsgSizePtr := flag.Int("maxMsgSize", 1000, "The maximun size of the massages to send in bytes")
	numOfMsgPtr := flag.Int("numOfMsg", 10000, "The number of messages to send")
	redisNodes := flag.String("redisNodes", "127.0.0.1:6379", "Cluster nodes declaration splitted by comma.")
	help := flag.Bool("help", false, "This help.")
	flag.Parse()
	if *minMsgSizePtr > *maxMsgSizePtr {
		log.Fatalln("minMsgSize (", *minMsgSizePtr, ")can't be bigger than maxMsgSize (", *maxMsgSizePtr, ").")
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if len(strings.Split(*redisNodes, ",")) == 1 {
		err := StressNode(*redisNodes, *minMsgSizePtr, *maxMsgSizePtr, *numOfMsgPtr)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := StressCluster(strings.Split(*redisNodes, ","), *minMsgSizePtr, *maxMsgSizePtr, *numOfMsgPtr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
