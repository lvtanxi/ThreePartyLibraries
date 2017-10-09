/**  
* Date: 2017-10-09 
* Time: 14:55 
* Description:
*/
package main

import (
	"encoding/xml"
	"os"
	"io/ioutil"
	"fmt"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}



func main() {
	readXml()
	writeXml()
}

func writeXml() {
	r :=&Recurlyservers{Version:"2"}
	r.Svs=append(r.Svs,server{ServerName:"chengdu_vpn",ServerIP:"192.168.0.1"})
	r.Svs=append(r.Svs,server{ServerName:"chongqing_vpn",ServerIP:"192.168.0.9"})
	outout,err:=xml.MarshalIndent(r," ","   ")
	CheckErr(err)
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(outout)
}

func readXml() {
	file, err := os.Open("test.xml")
	defer file.Close()
	CheckErr(err)
	data, err := ioutil.ReadAll(file)
	CheckErr(err)
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	CheckErr(err)
	fmt.Println(v.Version, ">>>>>>>>>>>>>>")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
