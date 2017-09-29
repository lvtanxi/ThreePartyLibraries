/**  
* Date: 2017-09-28 
* Time: 10:18 
* Description:
*/
package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	err = client.Set("username", "value", 0).Err()
	if err!=nil{
		fmt.Println(err)
		return
	}
	val, err:=client.Get("username").Result()
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("result %s : %s\n","username",val)
	val, err =client.Get("xxx").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val)
	}
	fmt.Printf("result %s : %s","xxx",val)
	defer client.Close()
}
