package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


func add() {
	user := &User{
		CreditCard: CreditCard{
			Number: "1234",
		},
	}
	db.Create(user)
}

func update()  {
	c :=&CreditCard{}
	db.First(c)
	c.Number="789x"
	db.Save(c)
}



//一对一查询
func query() {
	//这里是什么想要获取的东西
	u :=&User{}
	//指明外键关系
	db.First(u).Related(&u.CreditCard,"creditcard")
	log.Printf("user id is  %d and the CreditCard's number is %v", u.ID, u.CreditCard.Number)
}

//一对一查询
func queryOne() {
	//这里是什么想要获取的东西
	u :=&User{}
	//指明外键关系
	db.Preload("CreditCard").First(u)
	log.Printf("queryOne>>>user id is  %d and the CreditCard's number is %v", u.ID, u.CreditCard.Number)
}

func queryPreloadAll()  {
	users :=make([]*User,0)
	db.Preload("CreditCard").Find(&users)

	for index,user:=range users{
		log.Printf("index : %d ; user is : %d ; CreditCard : %s",index,user.ID,user.CreditCard.Number)
	}
}

func queryJoinAll()  {
	users :=make([]*User,0)
	db.Table("users").Select("users.id, credit_cards.number").Joins("left join credit_cards on credit_cards.user_id = users.id").Scan(&users)
	for index,user:=range users{
		log.Printf("index : %d ; user is : %d ; CreditCard : %s",index,user.ID,user.CreditCard.Number)
	}
}

func limitQuery()  {
	users :=make([]*User,0)
	db.Find(&users).Association("CreditCard")
	for index,user:=range users{
		log.Printf("index : %d ; user is : %d ; CreditCard : %s",index,user.ID,user.CreditCard.Number)
	}

}

