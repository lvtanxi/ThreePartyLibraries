package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"errors"
	"fmt"
	"os"
)

type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int    `xorm:"version"`
}

var x *xorm.Engine
var err error

func init() {
	x, err = xorm.NewEngine("mysql", "root:123456@/beeblog?charset=utf8")
	//记录日志(可以判断)
	f,err :=os.Create("sql.log")
	if err != nil {
		log.Fatalf("Fail to create engine:%v", err)
	}
	x.SetLogger(xorm.NewSimpleLogger(f))

	//显示sql
	x.ShowSQL(true)
	err = x.Sync(new(Account))
	if err != nil {
		log.Fatalf("Fail to Sync databass:%v", err)
	}
	//LRU缓存
	cacher :=xorm.NewLRUCacher(xorm.NewMemoryStore(),1000)
	x.SetDefaultCacher(cacher)
}

func newAccount(name string, balance float64) error {
	_, err = x.Insert(&Account{Name: name, Balance: balance})
	return err
}

func findAccountById(id int64) (*Account, error) {
	account := new(Account)
	has, err := x.Id(id).Get(account)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("account not found")
	}
	return account, err
}

func updateAccount(id int64, balance float64, isDeposit bool) (*Account, error) {
	account, err := findAccountById(id)
	if err != nil {
		log.Fatalf("Fail to updateAccount:%v", err)
		return nil, err
	}
	fmt.Println(account)
	if isDeposit {
		account.Balance += balance
	} else if account.Balance < balance {
		return nil, errors.New("not enough balance")
	} else {
		account.Balance -= balance
	}
	_, err = x.Id(id).Update(account)
	return account, err
}

func makeTransfer(fid, tid int64, balance float64) error {
	account, err := findAccountById(fid)
	if err != nil {
		log.Fatalf("Fail to makeTransfer:%v", err)
		return err
	}
	account2, err2 := findAccountById(tid)
	if err2 != nil {
		log.Fatalf("Fail to makeTransfer:%v", err)
		return err2
	}
	if account.Balance < balance {
		return errors.New("not enough balance")
	} else {
		account.Balance -= balance
		account2.Balance += balance
	}
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}
	_, err = sess.ID(fid).Update(account)
	if err != nil {
		log.Fatalf("Fail to makeTransfer:%v", err)
		sess.Rollback()
		return err
	}
	_, err = sess.ID(tid).Update(account2)
	if err != nil {
		log.Fatalf("Fail to makeTransfer:%v", err)
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func getAccountOderbyId() (as []*Account, err error) {
	err = x.Asc("id").Find(&as)
	return as, err
}

func deleteAccount(id int64) error {
	_, err = x.Delete(&Account{Id: id})
	return err
}

func findAccountByBalance(balance float64) (as []*Account, err error) {
	err = x.Where("balance < ?", balance).Find(&as)
	return as, err
}

func countAccount() ([]map[string][]byte, error) {
	return x.Query("SELECT count(1) as count FROM account")
}

func getAccountCount()(int64,error)  {
	return x.Count(new(Account))
}

//事件

func (a *Account)BeforeInsert()  {
	log.Printf("before inser:%s\n",a.Name)
}

func (a *Account)AfterInsert()  {
	log.Printf("After inser:%s\n",a.Name)
}



