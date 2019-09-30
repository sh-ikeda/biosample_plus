package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"net/http"
)

func redis_connection() redis.Conn {
	const IP_PORT = "127.0.0.1:6379"

	c, err := redis.Dial("tcp", IP_PORT)
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
	e := echo.New()
	cn := redis_connection()
	defer cn.Close()

	e.GET("/", func(c echo.Context) error {
		id := c.QueryParam("id")
		return c.String(http.StatusOK, redis_get(id, cn))
	})
	e.POST("/", func(c echo.Context) error {
		id := c.FormValue("id")
		return c.String(http.StatusOK, redis_get(id,cn))
	})
	e.Logger.Fatal(e.Start(":8080"))
}
