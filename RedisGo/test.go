/**  
* Date: 2017-09-28 
* Time: 10:42 
* Description:
*/
package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)

func main() {
	//获取本地连接
	client, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connent to redis error", err)
		return
	}
	//这里写入的值永远不会过期
	_, err = client.Do("SET", "MYKEY", "SuperWan")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	//获取值
	userName, err := redis.String(client.Do("GET", "MYKEY"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get MYKEY: %v \n", userName)
	}

	//设置多个值
	_, err = client.Do("MSET", "test1", "test1", "test2", "test2", "test3", "test3", )
	if err != nil {
		fmt.Println("redis set many failed:", err)
	}
	//获取多个值
	reply, err := redis.Values(client.Do("MGET", "test1", "test2", "test3"))
	if err != nil {
		fmt.Println("redis get many failed:", err)
	}
	var value1, value2, value3 string
	if _, err := redis.Scan(reply, &value1, &value2, &value3); err == nil {
		fmt.Println(value1, value2, value3)
	}
	//插入 map ,注意，这里跟一般的set方法不一致
	params := map[string]string{"1": "xx", "2": "xxx", "3": "xxxx"}
	_, err = client.Do("HMSET", redis.Args{}.Add("param").AddFlat(params)...)
	if err != nil {
		fmt.Println("redis set map failed:", err)
	}
	v, err := redis.StringMap(client.Do("HGETALL", "param"))
	if err != nil {
		fmt.Println("redis get map failed:", err)
	}
	fmt.Println(v)

	//插入array
	arra := []string{"v1", "v2", "v3", "v4"}
	_, err = client.Do("HMSET", redis.Args{}.Add("testArray").AddFlat(arra)...)
	if err != nil {
		fmt.Println("redis set array failed:", err)
	}
	values, err := redis.Values(client.Do("HGETALL", "testArray"))
	if err != nil {
		fmt.Println("redis get array 1 failed:", err)
	}

	result := make([]string, 0)
	err = redis.ScanSlice(values, &result)
	if err != nil {
		fmt.Println("redis get array 2 failed:", err)
	}
	fmt.Println(result)

	//插入结构体

	var user, user2 struct {
		Name string
		Age  int
	}
	user.Name = "lvtanxi"
	user.Age = 27
	_, err = client.Do("HMSET", redis.Args{}.Add("user").AddFlat(&user)...)
	if err != nil {
		fmt.Println("redis set struct failed:", err)
	}
	values, err = redis.Values(client.Do("HGETALL", "user"))

	if err != nil {
		fmt.Println("redis get struct 1 failed:", err)
	}
	err = redis.ScanStruct(values, &user2)
	if err != nil {
		fmt.Println("redis get struct 2 failed:", err)
	}
	fmt.Println(user2)
	//LPUSH 插入
	_, err = client.Do("DEL", "albums")
	client.Do("LPUSH", "albums", "album1")
	client.Do("LPUSH", "albums", "album2")
	client.Do("LPUSH", "albums", "album3")
	values, err = redis.Values(client.Do("lrange", "albums", 0, 3))
	if err != nil {
		fmt.Println("redis get LPUSH 1 failed:", err)
	}
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

	//检测值是否存在
	is_key_exit, err := redis.Bool(client.Do("EXISTS", "user"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("key user exists : %v \n", is_key_exit)
	}
	//删除key
	deleteKey(client)

	//如何设置过期呢，可以使用SET的附加参数：(2秒过期)
	_, err = client.Do("SET", "mykey", "superWang", "EX", "2")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	userName, err = redis.String(client.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", userName)
	}

	//3秒后获取之前的只
	time.Sleep(3 * time.Second)
	userName, err = redis.String(client.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", userName)
	}
	defer client.Close()
}
func deleteKey(client redis.Conn) {
	_, err := client.Do("SET", "deleteMykey", "xsadfasdf")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(client.Do("GET", "deleteMykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get deleteMykey: %v \n", username)
	}

	_, err = client.Do("DEL", "deleteMykey")
	fmt.Println("redis delelte failed:", err)
	username, err = redis.String(client.Do("GET", "deleteMykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get deleteMykey: %v \n", username)
	}
}
