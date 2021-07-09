package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Redis disconnected %v", err)
	}

	defer conn.Close()

	conn.Do("select", 0)
	conn.Do("set", "key1", "steven")
	key1, err := redis.String(conn.Do("get", "key1")) // redis有許多好用的轉換api
	if err != nil {
		log.Fatalln("we don't find the key1")
	}
	fmt.Println(key1)
}
