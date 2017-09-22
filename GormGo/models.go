package main

import (
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	gorm.Model
	CreditCard CreditCard                                    //一对一
	Emails     []Email                                       //一对多
	Languages  []Language `gorm:"many2many:user_languages;"` //多对多,需要说明中间表，好像可以只在一处生命
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

type Email struct {
	gorm.Model
	Address string
	UserID  uint
}

type Language struct {
	gorm.Model
	Name  string
	Users []User
}

var db *gorm.DB

func RegBD() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@/beeblog?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		log.Fatalf("open database error %s", err)
	}
	db.AutoMigrate(&Language{},&Email{}, &CreditCard{}, &User{})
}
