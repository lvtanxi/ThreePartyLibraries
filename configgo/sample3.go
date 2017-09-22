package main

import (
	"github.com/Unknwon/goconfig"
	"log"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("conf_work.ini")
	if err != nil{
		log.Fatalf("this config load faild %s",err)
		return
	}

	result,err:=cfg.GetSection(goconfig.DEFAULT_SECTION)
	if err != nil{
		log.Fatalf("GetSection %s load faild %s","goconfig.DEFAULT_SECTION",err)
		return
	}
	for key,vaule :=range result{
		log.Printf("%s : %s",key,vaule)
	}

	result,err=cfg.GetSection("courses")
	if err != nil{
		log.Fatalf("GetSection %s load faild %s","courses",err)
		return
	}
	for key,vaule :=range result{
		log.Printf("%s : %s",key,vaule)
	}
	cfg.SetKeyComments("courses","#3","https://github.com/Unknwon/go-rock-libraries-showcases")
	goconfig.SaveConfigFile(cfg,"output.ini")

	log.Printf("this is value %v",cfg.MustValue("dir.Go名库讲解.01-goconfig","name",""))

}
