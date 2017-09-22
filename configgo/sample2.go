package main

import (
	"log"
	"github.com/Unknwon/goconfig"
)

func main() {
	//获取gocinfig对象
	cfg, err := goconfig.LoadConfigFile("conf1.ini", "conf2.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件:%s", err)
	}
	//获取多个文件中的某一个key的值，后面的文件会覆盖前面文件的值
	value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "key_default")
	if err != nil {
		log.Fatalf("无法获取键值(%s): %s", "key_default", err)
	}
	log.Printf("%s > %s : %v ", goconfig.DEFAULT_SECTION, "key_default", value)

	//多文件覆盖加载
	err = cfg.AppendFiles("conf3.ini")
	if err != nil {
		log.Fatalf("无法追加文件:%s", "conf3.ini")
	}
	//配置文件重载
	//cfg.Reload()
	//为must系列方法设置缺省值
	vBool := cfg.MustBool("must","boo404",true)
	log.Printf("%s > %s : %v ", "must", "boo404", vBool)
	//递归读取值
	value,err = cfg.GetValue(goconfig.DEFAULT_SECTION,"search")
	if err != nil {
		log.Fatalf("无法获取键值(%s): %s", "search", err)
	}
	log.Printf("%s > %s : %v ", goconfig.DEFAULT_SECTION, "search", value)
	//子孙分区读取
	value,err = cfg.GetValue("parent.child","age")
	if err != nil {
		log.Fatalf("无法获取键值(%s): %s", "age", err)
	}
	log.Printf("%s > %s : %v ", "parent.child", "age", value)

	//子分区没有的属性会向上去找
	value,err = cfg.GetValue("parent.child","sex")
	if err != nil {
		log.Fatalf("无法获取键值(%s): %s", "sex", err)
	}
	log.Printf("%s > %s : %v ", "parent.child", "sex", value)
	//获取整个分区
	//自增键名获取
	sec,err := cfg.GetSection("auto increment")
	if err != nil {
		log.Fatalf("无法获取分区(%s): %s", "auto increment", err)
	}
	log.Printf("%s > %s : %v ", "auto increment", "auto increment", sec)

}
