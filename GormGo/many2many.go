package main

import "log"

func addLanguage()  {
	u :=&User{
		CreditCard:CreditCard{Number:"测试"},
		Emails:[]Email{
			{Address:"lvtanxi@613.com"},
		},
		Languages:[]Language{
			{Name:"中文"},
			{Name:"English"},
		},
	}
	db.Create(u)
}

func many2manyQuery()  {
	u :=&User{}
	//多对多的时候需要要指明关联列名
	db.First(&u,3).Related(&u.Emails).Related(&u.CreditCard).Related(&u.Languages,"languages")

	log.Printf("this is user's id : %d ;this is CreditCard' number : %s ",u.ID,u.CreditCard.Number)
	for index,email:=range u.Emails{
		log.Printf("this is Email index : %d ; Address : %s ",index,email.Address)
	}
	for index,language:=range u.Languages{
		log.Printf("this is Language index : %d ; Name : %s ",index,language.Name)
	}
}
func many2manyPreloadQuery()  {
	u :=&User{}
	//多对多的时候需要要指明关联列名
	db.Preload("Languages").Preload("Languages").Preload("Languages").First(&u,3)

	log.Printf("many2manyPreloadQuery >> this is user's id : %d ;this is CreditCard' number : %s ",u.ID,u.CreditCard.Number)
	for index,email:=range u.Emails{
		log.Printf("many2manyPreloadQuery >> this is Email index : %d ; Address : %s ",index,email.Address)
	}
	for index,language:=range u.Languages{
		log.Printf("many2manyPreloadQuery >> this is Language index : %d ; Name : %s ",index,language.Name)
	}
}

func limitmany2manyQuery()  {
	users :=make([]*User,0)
	db.Preload("Languages").Find(&users)
	for index,user:=range users{
		log.Printf("index : %d ; user is : %d ; Languages : %d",index,user.ID,len(user.Languages))
	}
}

