package main

import "fmt"
import "flag"
import "strconv"
import "math/rand"
import "log"
import "github.com/garyburd/redigo/redis"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes takes the size of the message and generates a pseudo
// random string
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	minMsgSizePtr := flag.Int("minMsgSize", 100, "The minimun size of the massages to send in bytes")
	maxMsgSizePtr := flag.Int("maxMsgSize", 1000, "The maximun size of the massages to send in bytes")
	numOfMsgPtr := flag.Int("NumOfMsg", 10000, "The number of messages to send")
	if *minMsgSizePtr > *maxMsgSizePtr {
		log.Fatalln("minMsgSize (", *minMsgSizePtr, ")can't be bigger than maxMsgSize (", *maxMsgSizePtr, ").")
	}
	// declare the valid range
	validSizeRange := *maxMsgSizePtr - *minMsgSizePtr
	//INIT OMIT
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	i := 0
	for i <= *numOfMsgPtr {
		//set
		c.Do("SET", strconv.Itoa(i), "Hello World")
	}

	//get
	world, err := redis.String(c.Do("GET", "message1"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
	//ENDINIT OMIT
}
