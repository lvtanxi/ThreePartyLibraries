package main

import (
	"fmt"
)

var printFn = func(index int, bean interface{}) error {
	fmt.Printf("%d: %#v\n", index, bean.(*Account))
	return nil
}

const prompt = `please enter number if operation:
1. create new Account
2. show detail of Account
3. Deposit
4. Withdraw
5. Make transfer
6. List Account by Id
7. List Account by balance
8. Delete Account
9. Exit
0. Count
`

func main() {
	fmt.Println("Welcome bank if xorm!")
EXIT:
	for {
		fmt.Println(prompt)
		var num int
		fmt.Scanf("%d\n", &num)
		switch num {
		case 1:
			add()
		case 2:
			showDetailById()
		case 3:
			chooseUpdateAccout("deposit", true)
		case 4:
			chooseUpdateAccout("withdraw", false)
		case 5:
			transfer()
		case 6:
			findAllAccount()
		case 7:
			inputBalance()
		case 8:
			deleteMethod()
		case 0:
			count()
		case 9:
			break EXIT
		default:
			iteration()
			rowsQuery()
			colsQuery()
			omitQuery()
			limitQuery()
		}
	}
}

func limitQuery() {
	x.Limit(2, 2).Iterate(new(Account), printFn)
}

//查询特定字段
func colsQuery() {
	x.Cols("name").Iterate(new(Account), printFn)
}

//排出特定字段
func omitQuery() {
	x.Omit("name").Iterate(new(Account), printFn)
}

func rowsQuery() {
	a := new(Account)
	rows, err := x.Rows(a)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(a); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%#v\n", a)
		}
	}
}

//迭代
func iteration() {
	fmt.Println("Query all records:")
	x.Iterate(new(Account), printFn)
}

func count() {
	results, err := countAccount()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(">>>>>>>>>", string(results[0]["count"]))
	}

	count, err := getAccountCount()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(">>>>>>>>>", count)
	}
}
func inputBalance() {
	fmt.Println("Please enter<balance>")
	var balance float64
	fmt.Scanf("%f\n", &balance)
	accounts, err := findAccountByBalance(balance)
	if err != nil {
		fmt.Println(err)
	} else {
		for index, acc := range accounts {
			fmt.Printf("%d:%v\n", index, acc)
		}
	}
}

func deleteMethod() {
	fmt.Println("Please Enter <id>")
	var id int64
	fmt.Scanf("%d\n", &id)
	err := deleteAccount(id)
	if err != nil {
		fmt.Println(err)
	}
}
func findAllAccount() {
	accounts, err := getAccountOderbyId()
	if err != nil {
		fmt.Println(err)
	} else {
		for index, acc := range accounts {
			fmt.Printf("%d:%v\n", index, acc)
		}
	}
}
func transfer() {
	fmt.Println("Please Enter <fid> <tid> <balance>")
	var fid, tid int64
	var balance float64
	fmt.Scanf("%d %d %f\n", &fid, &tid, &balance)
	if err := makeTransfer(fid, tid, balance); err != nil {
		fmt.Println(err)
	}
}

func chooseUpdateAccout(name string, isDeposit bool) {
	fmt.Println("Please Enter <id> " + name + "<balance>")
	var id int64
	var balance float64
	fmt.Scanf("%d %f\n", &id, &balance)
	account, err := updateAccount(id, balance, isDeposit)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", account)
	}
}

func showDetailById() {
	fmt.Println("Please Enter <id>")
	var id int64
	fmt.Scanf("%d\n", &id)
	account, err := findAccountById(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", account)
	}
}

func add() {
	fmt.Println("Please enter <name> <balance>")
	var name string
	var balance float64
	fmt.Scanf("%s %f\n", &name, &balance)
	if err := newAccount(name, balance); err != nil {
		fmt.Println(err)
	}
}
