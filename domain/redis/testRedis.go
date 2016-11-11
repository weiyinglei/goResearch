package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	server   string = "45.32.45.32:6379"
	password string = "Debian66"
)

var pool *redis.Pool

func test(i int) {

	c := pool.Get()
	defer c.Close()

	t := strconv.Itoa(i)
	c.Do("SETEX", "foo"+t, 20, i)

	reply, err := redis.Int(c.Do("GET", "foo"+t))
	if err == nil {
		fmt.Print(reply)
	} else {
		fmt.Print(err)
	}
	//time.Sleep(1 * time.Second)
}

func poolInit() *redis.Pool {
	//redis pool
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			//if _, err := c.Do("SELECT",1); err != nil {
			// c.Close()
			// return nil, err
			//}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

//init connction pool
//do something
func main() {
	fmt.Println(time.Now())
	pool = poolInit()

	for i := 0; i < 10; i++ {
		test(i)
	}
}
