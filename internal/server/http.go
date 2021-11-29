package server

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func Echo()  {
	fmt.Println("aaaaaa")
	c, err := redis.Dial("tcp", "localhost:6379")
	if err == nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
}