package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"os"
	"strings"
)

func redis_connection() redis.Conn {
	const IP_PORT = "127.0.0.1:6379"

	c, err := redis.Dial("tcp", IP_PORT)
	if err != nil {
		panic(err)
	}
	return c
}

func redis_set(key string, value string, c redis.Conn) {
	c.Do("SET", key, value)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: insert_json json_file")
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Loading json file...")
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Done.")

	fmt.Fprintln(os.Stderr, "Parsing...")
	var decode_data interface{}
	if err := json.Unmarshal(bytes, &decode_data); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Done. (Total %d entries)\n", len(decode_data.([]interface{})))

	c := redis_connection()
	defer c.Close()

	fmt.Fprintln(os.Stderr, "Inserting data into redis DB...")
	for _, data := range decode_data.([]interface{}) {
		d := data.(map[string]interface{})
		id := strings.Trim(d["@id"].(string), "biosample:")
		s, _ := json.Marshal(d)
		redis_set(id, string(s), c)
		//fmt.Fprintln(os.Stdout, id)
		//fmt.Fprintln(os.Stdout, string(s))
	}
	fmt.Fprintln(os.Stderr, "Done.")
}
