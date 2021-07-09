package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gomodule/redigo/redis"
)

type Student struct {
	Name  string
	Age   int
	Skill string
}

func CleanByKey(conn redis.Conn, key string) (err error) {

	_, err = conn.Do("DEL", key)
	return
}

func SetHashStruct(s interface{}, conn redis.Conn, no string) (key string, err error) {

	rType := reflect.TypeOf(s)

	rVal := reflect.ValueOf(s)

	num := rVal.NumField()

	key = rType.Name() + no

	err = CleanByKey(conn, key)
	if err != nil {
		return
	}

	for i := 0; i < num; i++ {

		typeField := rType.Field(i).Name

		valField := rVal.Field(i)

		_, err = conn.Do("HSET", key, typeField, valField)

		if err != nil {
			return
		}
	}

	return
}

func GetHashStruct(conn redis.Conn, key string) ([]string, error) {

	data, err := redis.Strings(conn.Do("HGETALL", key))

	return data, err

}

func main() {

	stud01 := Student{
		Name:  "Steven",
		Age:   24,
		Skill: "breaking",
	}

	// 其實這個conn回傳值本身就是一個private的conn structure的pointer，但是把它包成public的Conn的interface
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Printf("Redis disconnected %v", err)
	}

	defer conn.Close()

	key, err := SetHashStruct(stud01, conn, "1")
	if err != nil {
		log.Println("Set Hash Fail", err)
	}

	data, err := GetHashStruct(conn, key)
	if err != nil {
		log.Println("Get Hash Fail", err)
	}

	for i, v := range data {
		fmt.Print(v, " ")
		if i%2 == 1 {
			fmt.Println()
		}
	}

}
