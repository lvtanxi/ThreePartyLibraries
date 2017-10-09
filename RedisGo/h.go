/**  
* Date: 2017-09-30 
* Time: 14:50 
* Description:
*/
package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"regexp"
)

func main() {
	a := [...]int{5, 4: 1, 0, 2: 3, 2, 1: 4}

	fmt.Println(a)

	m,err :=regexp.MatchString("^[0-9]+$","123456")
	if err !=nil {
		fmt.Println(err)
	}else {
		fmt.Println(m)
	}

	http.HandleFunc("/", sayHelloName)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for key, value := range r.Form {
		fmt.Println("key:", key)
		fmt.Println("val:", strings.Join(value, ""))
	}
	fmt.Fprint(w, "hell ad")

}
