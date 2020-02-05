package bsplus

import (
_	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/nitishm/go-rejson"
	"os"
)

func Redis_connection(address string) redis.Conn {
	c, err := redis.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	return c
}

func Redis_get(key string, c redis.Conn) string {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	return s
}

func Redis_set(key string, value string, c redis.Conn) {
	c.Do("SET", key, value)
}

func Redis_json_set(key string, value interface{}, rh *rejson.Handler) {
	rh.JSONSet(key, ".", value)
}

func Redis_json_get(key string, rh *rejson.Handler) string {
	s, err := redis.String(rh.JSONGet(key, "."))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(os.Stderr, s[0:1])
	return s
}
