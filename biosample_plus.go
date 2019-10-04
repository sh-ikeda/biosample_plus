package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func redis_connection(address string) redis.Conn {
	c, err := redis.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	return c
}

func redis_get(key string, c redis.Conn) string {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	return s
}

func main() {
	address := ""
	if len(os.Args) == 1 {
		address = "127.0.0.1:6379"
	} else {
		address = os.Args[1]
	}
	e := echo.New()
	cn := redis_connection(address)
	defer cn.Close()

	e.GET("/", func(c echo.Context) error {
		id := c.QueryParam("id")
		return c.String(http.StatusOK, redis_get(id, cn))
	})
	e.POST("/", func(c echo.Context) error {
		id := c.FormValue("id")
		return c.String(http.StatusOK, redis_get(id, cn))
	})
	e.Logger.Fatal(e.Start(":8080"))
}
