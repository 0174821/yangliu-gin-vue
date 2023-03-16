package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"yangliu-gin-vue/models"
	"yangliu-gin-vue/utils"
)

var currentAccount *models.Account

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
)

func init() {
	log.SetPrefix("[YANGLIU]")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
	filepath := path.Join("logs", "yangliu.log")
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(file)

	// 设置日志级别
	Info = log.New(file, "[info]", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
	Warning = log.New(file, "[Warning]", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
	Error = log.New(file, "[Error]", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
	Fatal = log.New(file, "[Fatal]", log.LstdFlags|log.Lmicroseconds|log.Llongfile)

}

func LoginService(name, pwd string) (string, bool, *models.Account) {
	filepath, ok := utils.IsExist(name)
	if !ok {
		fmt.Println("当前用户名不存在")
		return "用户名不存在", false, nil
	}

	// 判断name是否已经存在
	// filepath := "./accounts/" + name + ".json"
	// _, err := os.Stat(filepath)
	// if os.IsNotExist(err) {
	// 	// fmt.Println(os.IsNotExist(err))
	// 	fmt.Println("当前用户名不存在")
	// 	return
	// }

	// 打开json文件，发序列化处账号信息
	accountjson, _ := ioutil.ReadFile(filepath)
	var account models.Account
	json.Unmarshal(accountjson, &account)
	fmt.Println(account)

	// 判断name是不是已经注册了
	// _, ok := allAccounts[name]
	// if !ok {
	// 	fmt.Println("没有这个用户")
	// }

	// 再次判断用户输入的密码和注册时保存的密码是否一致
	// account := allAccounts[name]
	if account.Pwd != utils.Md5Salt(pwd) {
		return "密码错误", false, nil
	} else {
		account.IsLogin = true
		currentAccount = &account
		Info.Printf("%s登陆成功", currentAccount.Name)
		return "登录成功", true, currentAccount
	}
}

func RegisterService(name, pwd string) (string, bool, *models.Account) {

	// 如何判断当前路径下是否存在一个文件？
	filepath, ok := utils.IsExist(name)
	if ok {
		fmt.Println("当前用户名已经被注册")
		return "当前用户名已经被注册", false, nil
	}

	// 如何判断当前路径下是否存在一个文件？
	// filepath := "./accounts/" + name + ".json"
	// _, err := os.Stat(filepath)
	// if !os.IsNotExist(err) {
	// 	fmt.Println("当前用户名已经被注册")
	// 	return
	// }

	// 判断用户名是否已经被注册
	// _, ok := allAccounts[name]
	// fmt.Println(ok)
	// if ok {
	// 	fmt.Println("当前用户名已经被注册")
	// 	return
	// }

	// 可以注册了
	// newAccount := Account{
	// 	Name:    name,
	// 	Pwd:     pwd,
	// 	Balance: 0,
	// }

	newAccount := models.NewAccount(name, utils.Md5Salt(pwd))
	// allAccounts[name] = *newAccount

	// 序列化成json格式，存放到文件中
	data, _ := json.Marshal(newAccount)
	os.WriteFile(filepath, data, 0666)

	return "注册成功", true, newAccount
}

func ShowBalanceService() (string, int, bool) {
	if currentAccount == nil {
		return "未登录", 0, false
	}
	return "余额查询成功", currentAccount.ShowBalance(), true
}

func ShowBalanceDetailService() (string, bool) {
	if currentAccount == nil {
		return "未登录", false
	}
	return "明细查询成功", currentAccount.ShowBalanceDetail()
}

func UpBalanceService() string {
	if currentAccount == nil {
		return "未登录"
	}
	return currentAccount.UpBalance()
}

func DownBalance() string {
	if currentAccount == nil {
		return "未登录"
	}
	return currentAccount.DownBalance()
}
