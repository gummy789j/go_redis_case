package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {

	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
}

func main() {

	conn := pool.Get()

	defer conn.Close()

	_, err := conn.Do("SET", "key1", "Steven")
	if err != nil {
		log.Println("set FAIL")
	}

	data, err := redis.String(conn.Do("GET", "key1"))
	if err != nil {
		log.Println("get FAIL")
	}

	fmt.Println(data)

	// pool close以後 不能再取出conn
	pool.Close()

}
