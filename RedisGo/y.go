/**  
* Date: 2017-09-30 
* Time: 09:39 
* Description:
*/
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) string() string {
	return fmt.Sprintf("name: %s - age: %d", p.name, p.age)
}

func main() {
	list := make([]interface{}, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"lvtanxi", 28}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a  Person and its value is %s\n", index, value.string())
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

	for index, element := range list {
		switch value :=element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a  Person and its value is %s\n", index, value.string())
		default:
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}


}
