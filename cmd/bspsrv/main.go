package main

import (
_	"fmt"
_	"github.com/garyburd/redigo/redis"
_	"github.com/nitishm/go-rejson"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"bsplus"
)

type Params struct {
	Ids []string `json:"id"`
	//Ids string `json:"id"`
}

func main() {
	address := ""
	if len(os.Args) == 1 {
		address = "127.0.0.1:6379"
	} else {
		address = os.Args[1]
	}
	e := echo.New()
	cn := bsplus.Redis_connection(address)
	defer cn.Close()
	// rh := rejson.NewReJSONHandler()
	// rh.SetRedigoClient(cn)

	e.GET("/", func(c echo.Context) error {
		id := c.QueryParam("id")
		return c.String(http.StatusOK, bsplus.Redis_get(id, cn))
		//return c.String(http.StatusOK, bsplus.Redis_json_get(id, rh))
	})

	e.POST("/", func(c echo.Context) error {
		id := c.FormValue("id")
		return c.String(http.StatusOK, bsplus.Redis_get(id, cn))
	})

	e.POST("/api/", func(c echo.Context) error {
		id := new(Params)
		if err := c.Bind(id); err != nil {
			return err
		}
		resp := "["
		for k := range id.Ids {
			if resp != "[" {
				resp += ","
			}
			resp += bsplus.Redis_get(id.Ids[k], cn)
		}
		resp += "]"
		return c.String(http.StatusOK, resp)
		//return c.String(http.StatusOK, redis_get(id.Ids[0], cn))
		//return c.String(http.StatusOK, redis_get(id.Ids, cn))
		//return c.JSON(http.StatusOK, id)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
