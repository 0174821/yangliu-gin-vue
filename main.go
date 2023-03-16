package main

import (
	"fmt"
	"os"
	"yangliu-gin-vue/services"
)

func Login() {
	var (
		name string
		pwd  string
	)
	fmt.Printf("请输入用户名:")
	fmt.Scanln(&name)
	fmt.Printf("请输入密码:")
	fmt.Scanln(&pwd)

	message, _, _ := services.LoginService(name, pwd)
	fmt.Println(message)

}

func Register() {
	var (
		name     string
		pwd      string
		pwdAgain string
	)

	fmt.Printf("请输入用户名:")
	fmt.Scanln(&name)
	fmt.Printf("请输入密码:")
	fmt.Scanln(&pwd)
	fmt.Printf("请再次输入密码：")
	fmt.Scanln(&pwdAgain)

	if pwd != pwdAgain {
		fmt.Println("两次密码输入不一致")
		return
	}

	message, _, _ := services.RegisterService(name, pwd)
	fmt.Println(message)

}

func ShowBalance() {
	msg, amount, ok := services.ShowBalanceService()
	fmt.Println(msg)
	if ok {
		fmt.Println("您当前余额：", amount)
	}
}

func ShowBalanceDetail() {
	msg, _ := services.ShowBalanceDetailService()
	fmt.Println(msg)
}

func UpBalance() {
	msg := services.UpBalanceService()
	fmt.Println(msg)
}

func DownBalance() {
	msg := services.DownBalance()
	fmt.Println(msg)
}

func main() {
	fmt.Println("欢迎使用yangliu记账本软件")
	var choice int
	menuInfo := `
		1 登录
		2 注册
		3 余额
		4 明细
		5 收入
		6 支出
		7 退出
	`

	for {
		fmt.Println(menuInfo)
		fmt.Printf("请选择一个功能编号：")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Login()
		case 2:
			Register()
		case 3:
			ShowBalance()
		case 4:
			ShowBalanceDetail()
		case 5:
			UpBalance()
		case 6:
			DownBalance()
		case 7:
			os.Exit(0) // 0表示正常退出
		default:
			fmt.Println("请选择合法的功能编号")
		}
	}

}
