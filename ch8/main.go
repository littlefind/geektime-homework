package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//连接redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn err = ", err)
		return
	} else {
		fmt.Println("conn success")
	}
	defer conn.Close()
	
	//生成key和value
	for i := 1; i <= 1000000; i++ {
		k := i + 100000000
		v := i + 100000000
		//写入进redis
		_, err = conn.Do("set", k, v)
		if err != nil {
			fmt.Println("set err = ", err)
			return
		} else {
			fmt.Printf("set success k = %v, v = %v\n", k, v)
		}
	}
}
