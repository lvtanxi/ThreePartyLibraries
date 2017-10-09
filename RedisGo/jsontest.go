/**  
* Date: 2017-10-09 
* Time: 15:35 
* Description:
*/
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Users struct {
	Users []User `json:"users"`
}

type Student struct {
	User
	ServerIp string `json:"serverIp,omitempty"` // 如果 ServerIP 为空，则不输出到JSON串中
}

func main() {
	users :=Users{Users:[]User{{"测试1",1},{"测试2",2}}}
	data,err:=json.Marshal(users)
	checkErrx(err)
	newU :=Users{}
	json.Unmarshal(data,&newU)
	fmt.Println(string(data),newU)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	f :=make(map[string]interface{},0)
	json.Unmarshal(b,&f)
	fmt.Println(f)
	for key, value := range f {
		switch vv :=value.(type) {
		case string:
			fmt.Println(key, "is string", vv)
		case int:
			fmt.Println(key, "is int", vv)
		case []interface{}:
			fmt.Println(key, "is an array")
			for k, v := range vv {
				fmt.Println(k, v)
			}
		default:
			fmt.Println(key, "is of a type I don't know how to handle",vv)
		}
	}

	stu :=Student{User:User{"测试1",1}}
	data,err =json.Marshal(stu)
	checkErrx(err)
	fmt.Println(string(data))
}

func checkErrx(err error) {
	if err != nil {
		panic(err)
	}
}
