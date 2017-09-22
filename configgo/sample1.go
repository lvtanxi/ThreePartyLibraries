package main

import (
	"github.com/Unknwon/goconfig"
	"log"
)

func main() {
	//获取gocinfig对象

	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件:%s", err)
	}
	//获取默认分区的数据
	value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "key_default")
	if err != nil {
		log.Fatalf("无法获取key_default的值:%s", err)
	}
	log.Printf("%s > %s: %s", goconfig.DEFAULT_SECTION, "key_default", value)

	//插入新值
	isInsert :=cfg.SetValue(goconfig.DEFAULT_SECTION,"key_default","这是新的值")
	log.Printf("是否为插入操作: %v",isInsert)

	//注释读写操作
	//获取某个分区的注释
	comment :=cfg.GetSectionComments("super")
	log.Printf("获取super分区的注释为：%s",comment)

	//设置某个分区的注释(#是必须的)
	isInsert=cfg.SetSectionComments("super","#这只是我为super分区添加的测试注释")
	log.Printf("为super分区添加注释的结果为: %v",isInsert)

	//类型转换读取
	vInt,err :=cfg.Int("must","int")
	if err != nil {
		log.Fatalf("无法must分区中的int的值:%s", err)
	}
	log.Printf("获取到must分区中的int的值:%v", vInt)
	//Must系列方法
	aBool :=cfg.MustBool("must","bool")
	log.Printf("获取到must分区中的bool的值:%v", aBool)
	//删除指定键值
	ok:=cfg.DeleteKey("must","string")
	log.Printf("删除键值string是否成功：%v",ok)
	//保持配置文件
	err =goconfig.SaveConfigFile(cfg,"conf_save.ini")
	if err != nil {
		log.Fatalf("无法保持配置文件:%s", err)
	}
}
