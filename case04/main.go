package main

import (
	"log"
	"math/rand"
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

func Record(conn redis.Conn, key string, stuff string) error {

	_, err := conn.Do("LPUSH", key, stuff)

	return err

}

func CheckRecord(conn redis.Conn, key string, n int) (data []string, err error) {
	data, err = redis.Strings(conn.Do("LRANGE", key, 0, n-1))
	return
}

func RedisConnect(network string, address string) (conn redis.Conn, err error) {
	conn, err = redis.Dial(network, address)
	return
}

func main() {

	network := "tcp"

	address := ":6379"

	conn, err := RedisConnect(network, address)
	if err != nil {
		log.Printf("Redis disconnected %v", err)
	}

	defer conn.Close()

	commodities := []string{

		"soap",
		"face wash",
		"lotion",
		"hairbrush",
		"shaver",
		"folder chair",
		"bookcase",
		"comforter",
		"heater",
		"electric fan",
		"socket",
		"slippers",
		"pillowcase",
		"sheet",
		"mattress",
	}

	// we need a number as our experiment factor
	n := len(commodities)

	rand.Seed(int64(475))

	customer := "Mr.Lin"

	err = CleanByKey(conn, customer)
	if err != nil {
		log.Println("Clean list Fail")
	}

	for i := 0; i < n; i++ {
		stuff := commodities[rand.Intn(n)]
		err := Record(conn, customer, stuff)
		if err != nil {
			log.Println("Store list Fail")
		}
	}

	lastestTimes := 10

	data, err := CheckRecord(conn, customer, lastestTimes)
	if err != nil {
		log.Println("inquery browsing Record Fail")
	}

	for i, v := range data {
		log.Printf("Record[%d] = %s", i+1, v)
	}

}
