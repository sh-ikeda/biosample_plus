package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/garyburd/redigo/redis"
//	"net/http"
)

func redis_connection() redis.Conn {
	const IP_PORT = "127.0.0.1:6379"

	c, err := redis.Dial("tcp", IP_PORT)
	if err != nil {
		panic(err)
	}
	return c
}

func redis_set(key string, value string, c redis.Conn){
  c.Do("SET", key, value)
}

func main() {
	if len(os.Args)==1 {
		fmt.Fprintln(os.Stderr, "usage: insert_ttl turtle_file")
		os.Exit(1)
	}

	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fp.Close()

	c := redis_connection()
	defer c.Close()

	scanner := bufio.NewScanner(fp)
	header := ""
	body := ""
	idp := "ns3:identifier"
	id := ""
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if line[0] == '@' {
				header += line + "\n"
			} else {
				body += line + "\n"
				if strings.Contains(line, idp) {
					id = line[(strings.Index(line, "\"")+1):strings.LastIndex(line, "\"")]
				}
				if line[len(line)-1] == '.' {
					//fmt.Fprintln(os.Stdout, id, "\n", header, "\n", body)
					//fmt.Fprintln(os.Stdout, id, id)
					redis_set(id, header+"\n"+body, c)
					body = ""
				}
			}
		}
	}
}
