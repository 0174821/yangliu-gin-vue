package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type Account struct {
	Name    string   `json:"name"`
	Pwd     string   `json:"pwd"`
	Balance int      `json:"balance"`
	Details []Detail `json:"details"`
	IsLogin bool     `json:"isLogin"`
}

type Detail struct {
	Kind    string `json:"kind"`
	Amounts int    `json:"amounts"`
	Message string `json:"message"`
}

func NewAccount(name, pwd string) *Account {
	return &Account{
		Name:    name,
		Pwd:     pwd,
		Balance: 0,
	}
}

func (ac *Account) Filepath() string {
	return path.Join("./accounts", fmt.Sprint(ac.Name, ".json"))
}

func (ac *Account) ShowBalance() int {
	return ac.Balance
}

func (ac *Account) ShowBalanceDetail() bool {
	if len(ac.Details) == 0 {
		fmt.Println("您当前没有收支记录")
		return false
	}
	fmt.Println("您的收支明细如下：")
	for index, d := range ac.Details {
		fmt.Printf("(%d)%s\t\t%d\t\t%s\n", index+1, d.Kind, d.Amounts, d.Message)
	}
	return true
}

func (ac *Account) UpBalance() string {
	var (
		amounts int
		message string
	)

	fmt.Print("请输入收入金额：")
	fmt.Scanln(&amounts)
	fmt.Print("请输入输入缘由：")
	fmt.Scanln(&message)

	ac.Balance = ac.Balance + amounts
	detail := Detail{
		Kind:    "收入",
		Amounts: amounts,
		Message: message,
	}
	ac.Details = append(ac.Details, detail)

	data, _ := json.Marshal(ac)
	os.WriteFile(ac.Filepath(), data, 0666)
	return "收入记录成功"
}

func (ac *Account) DownBalance() string {
	var (
		amounts int
		message string
	)

	fmt.Print("请输入支出金额：")
	fmt.Scanln(&amounts)
	fmt.Print("请输入支出缘由：")
	fmt.Scanln(&message)

	if amounts > ac.Balance {
		return "余额不足！"
	}

	ac.Balance = ac.Balance - amounts
	detail := Detail{
		Kind:    "支出",
		Amounts: amounts,
		Message: message,
	}
	ac.Details = append(ac.Details, detail)

	data, _ := json.Marshal(ac)
	os.WriteFile(ac.Filepath(), data, 0666)
	return "支出记录成功"
}

// 使用内存来临时保存数据
var allAccounts = make(map[string]Account, 100)
