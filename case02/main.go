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

	_, err = conn.Do("del", "user1")
	if err != nil {
		log.Fatalln("Clean the key of user1")
	}

	_, err = conn.Do("hset", "user1", "name", "steven")
	if err != nil {
		log.Fatalln("store hash data Fail")
	}
	_, err = conn.Do("hset", "user1", "age", 24)
	if err != nil {
		log.Fatalln("store hash data Fail")
	}
	user1_name, err := redis.String(conn.Do("hget", "user1", "name"))
	if err != nil {
		log.Fatalln("we don't find the user1_name")
	}
	fmt.Println(user1_name)

	user1_age, err := redis.String(conn.Do("hget", "user1", "age"))
	if err != nil {
		log.Fatalln("we don't find the user1_age")
	}
	fmt.Println(user1_age)

	// 一次取出整個hash table => 要搭配 Strings 不是String
	user1, err := redis.Strings(conn.Do("hgetall", "user1"))
	if err != nil {
		log.Fatalln("we don't find the user1")
	}
	fmt.Printf("%T\n", user1)

	for i, v := range user1 {
		fmt.Print(v, " ")
		if i%2 == 1 {
			fmt.Println()
		}
	}
}
