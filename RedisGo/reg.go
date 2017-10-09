/**  
* Date: 2017-10-09 
* Time: 16:01 
* Description:
*/
package main

import (
	"regexp"
	"fmt"
)

func main() {
	fmt.Println(isIp("192.168.0.1z"))
}

func isIp(ip string) bool {
	m, err := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$", ip)
	if err !=nil{
		fmt.Println(err)
		return false
	}
	return m
}
