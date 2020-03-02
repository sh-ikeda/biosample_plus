package main

import (
	"bufio"
	"github.com/labstack/echo"
	"net/http"
	"os"

	"bsplus"
)

type Params struct {
	Ids []string `json:"id" form:"id"`
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

	e.GET("/", func(c echo.Context) error {
		id := c.QueryParam("id")
		bsp_entry , err := bsplus.Redis_get(id, cn)
		if err != nil {
			return c.String(http.StatusNotFound, "ID \"" + id + "\" is not found on BioSamplePlus.")
		}
		return c.String(http.StatusOK, bsp_entry)
	})

	e.POST("/api/", func(c echo.Context) error {
		ids := new(Params)
		if err := c.Bind(ids); err != nil {
			return err
		}
		resp := "["
		for k := range ids.Ids {
			bsp_entry, err := bsplus.Redis_get(ids.Ids[k], cn)
			if err == nil {
			if resp != "[" {
				resp += ","
			}
				resp += bsp_entry
			}
		}
		resp += "]"
		return c.String(http.StatusOK, resp)
	})

	e.POST("/list/", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		scanner := bufio.NewScanner(src)
		resp := "["
		for scanner.Scan() {
			id := scanner.Text()
			bsp_entry, err := bsplus.Redis_get(id, cn)
			if err == nil {
				if resp != "[" {
					resp += ","
				}
				resp += bsp_entry
			}
		}
		resp += "]"
		return c.String(http.StatusOK, resp)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
