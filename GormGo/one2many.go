package main

import "log"

func addEmail()  {
	u :=&User{
		CreditCard:CreditCard{Number:"测试"},
		Emails:[]Email{
			{Address:"lvtanxi@613.com"},
		},
	}
	db.Create(u)
}

func oneToManyQuery(){
	u :=&User{
		CreditCard:CreditCard{},
		Emails:[]Email{},
	}
	db.First(&u,4).Related(&u.CreditCard,"creditcard").Related(&u.Emails,"emails")

	log.Printf("this is user's id : %d ;this is CreditCard' number : %s ",u.ID,u.CreditCard.Number)
	for index,email:=range u.Emails{
		log.Printf("this is Email index : %d ; Address : %s ",index,email.Address)
	}
}

func oneToManyPreloadQuery(){
	u :=&User{}
	//多个外键需要连续Preload
	db.Preload("CreditCard").Preload("Emails").First(&u,3)

	log.Printf("oneToManyPreloadQuery  >>> this is user's id : %d ;this is CreditCard' number : %s ",u.ID,u.CreditCard.Number)
	for index,email:=range u.Emails{
		log.Printf("oneToManyPreloadQuery  >>> this is Email index : %d ; Address : %s ",index,email.Address)
	}
}

func oneToManyAllQuery(){
	users :=make([]*User,0)

	//多个外键需要连续Preload
	db.Preload("CreditCard").Preload("Emails").Find(&users)
	for index,user:=range users{
		log.Printf("oneToManyAllQuery  >>> index : %d ; user is : %d ; CreditCard : %s",index,user.ID,user.CreditCard.Number)
		for index,email:=range user.Emails{
			log.Printf("oneToManyAllQuery  >>> this is Email index : %d ; Address : %s ",index,email.Address)
		}
	}

}