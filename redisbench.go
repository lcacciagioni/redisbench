package main

import "flag"
import "strconv"
import "math/rand"

import "time"
import "log"

import "strings"
import redigo "github.com/garyburd/redigo/redis"
import redigocluster "github.com/chasex/redis-go-cluster"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandStringBytesMaskImprSrc is a function highly optimized for generation of
// random strings in golang
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// RandRange returns a random string with using the min and max values as limit.
func RandRange(min, max int) string {
	randsize := rand.Intn(max - min)
	return RandStringBytesMaskImprSrc(min + randsize)
}

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
	log.Println("Starting SET")
	startSet := time.Now()
	for i := 0; i <= numOfMsg; i++ {
		rediskey := "message" + strconv.Itoa(i)
		message := RandRange(minMsgSize, maxMsgSize)
		//set
		c.Do("SET", rediskey, message)
	}
	log.Println("SET of ", strconv.Itoa(numOfMsg), " random messages has taken:",
		time.Since(startSet))

	log.Println("Starting GET")
	startGet := time.Now()
	for j := 0; j <= numOfMsg; j++ {
		rediskey := "message" + strconv.Itoa(j)
		//get
		_, err := redigo.String(c.Do("GET", rediskey))
		if err != nil {
			log.Println("key not found")
		}
	}
	log.Println("GET of ", strconv.Itoa(numOfMsg), " random messages has taken:", time.Since(startGet))
}

// StressNode is fuction that executes a series of stress tests in a sigle node
// of redis
func StressNode(host string, minMsgSize, maxMsgSize, numOfMsg int) {
	c, err := redigo.Dial("tcp", host)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()
	log.Println("Starting SET")
	startSet := time.Now()
	for i := 0; i <= numOfMsg; i++ {
		rediskey := "message" + strconv.Itoa(i)
		message := RandRange(minMsgSize, maxMsgSize)
		//set
		c.Do("SET", rediskey, message)
	}
	log.Println("SET of ", strconv.Itoa(numOfMsg), " random messages has taken:",
		time.Since(startSet))

	log.Println("Starting GET")
	startGet := time.Now()
	for j := 0; j <= numOfMsg; j++ {
		rediskey := "message" + strconv.Itoa(j)
		//get
		_, err := redigo.String(c.Do("GET", rediskey))
		if err != nil {
			log.Println("key not found")
		}
	}
	log.Println("GET of ", strconv.Itoa(numOfMsg), " random messages has taken:", time.Since(startGet))
}

func main() {
	minMsgSizePtr := flag.Int("minMsgSize", 100, "The minimun size of the massages to send in bytes")
	maxMsgSizePtr := flag.Int("maxMsgSize", 1000, "The maximun size of the massages to send in bytes")
	numOfMsgPtr := flag.Int("numOfMsg", 10000, "The number of messages to send")
	redisNodes := flag.String("redisNodes", "127.0.0.1:6379", "Cluster nodes declaration splitted by comma.")
	flag.Parse()
	if *minMsgSizePtr > *maxMsgSizePtr {
		log.Fatalln("minMsgSize (", *minMsgSizePtr, ")can't be bigger than maxMsgSize (", *maxMsgSizePtr, ").")
	}

	if len(strings.Split(*redisNodes, ",")) == 1 {
		StressNode(*redisNodes, *minMsgSizePtr, *maxMsgSizePtr, *numOfMsgPtr)
	} else {
		StressCluster(strings.Split(*redisNodes, ","), *minMsgSizePtr, *maxMsgSizePtr, *numOfMsgPtr)
	}

	//ENDINIT OMIT
}
