package main

import "fmt"
import "flag"
import "strconv"
import "math/rand"

import "time"
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

// RandRange returns a random string with using the min and max values as limit.
func RandRange(min, max int) string {
	randsize := rand.Intn(max - min)
	return RandStringBytes(min + randsize)
}

func main() {
	minMsgSizePtr := flag.Int("minMsgSize", 100, "The minimun size of the massages to send in bytes")
	maxMsgSizePtr := flag.Int("maxMsgSize", 1000, "The maximun size of the massages to send in bytes")
	numOfMsgPtr := flag.Int("numOfMsg", 10000, "The number of messages to send")
	redisHostPtr := flag.String("redisHost", "localhost", "The host where redis is located.")
	redisPortPtr := flag.Int("redisPort", 6379, "The port where redis is listening")
	flag.Parse()
	if *minMsgSizePtr > *maxMsgSizePtr {
		log.Fatalln("minMsgSize (", *minMsgSizePtr, ")can't be bigger than maxMsgSize (", *maxMsgSizePtr, ").")
	}
	//INIT OMIT
	c, err := redis.Dial("tcp", *redisHostPtr+":"+strconv.Itoa(*redisPortPtr))
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	fmt.Println("Starting SET")
	startSet := time.Now()
	for i := 0; i <= *numOfMsgPtr; i++ {
		rediskey := "message" + strconv.Itoa(i)
		message := RandRange(*minMsgSizePtr, *maxMsgSizePtr)
		//set
		c.Do("SET", rediskey, message)
	}
	fmt.Println("SET of ", strconv.Itoa(*numOfMsgPtr), " random messages has taken:", time.Since(startSet))

	fmt.Println("Starting GET")
	startGet := time.Now()
	for j := 0; j <= *numOfMsgPtr; j++ {
		rediskey := "message" + strconv.Itoa(j)
		//get
		_, err := redis.String(c.Do("GET", rediskey))
		if err != nil {
			fmt.Println("key not found")
		}
	}
	fmt.Println("GET of ", strconv.Itoa(*numOfMsgPtr), " random messages has taken:", time.Since(startGet))
	//ENDINIT OMIT
}
