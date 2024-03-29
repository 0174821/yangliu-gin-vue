## go module

- go module是go的包管理工具，允许程序包不放在src下面
- 通过 go mod init 项目包名
  可以初始化一个项目名
- 如果是需要写web项目，需要通过一下命令安装依赖包
  go get -u github.com/gin-gonic/gin
  下载完成之后，在项目包名这个文件夹下的go.mod文件中会添加所增加的依赖包

## 推送github

```shell
git add README.md
git config --global user.email "2459770156@qq.com"  #已配置可以跳过
git config --global user.name "0174821"  #已配置可以跳过
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/0174821/yangliu-gin-vue.git
git config --global http.postBuffer 524288000
git config --global http.sslVerify "false"
git push -u origin main
```

## 临时存储和持久化存储

- 临时存储的实现是通过将数据保存在内存当中，当程序重新运行的时候，数据将会丢失

```GO
// 使用内存来临时保存数据
var allAccounts = make(map[string]Account, 100)
// 通过该方式从变量中取值和像变量中写值
account := allAccounts[name]
allAccounts[name] = *newAccount
```

- 持久化存储是将数据库以json格式的形式存放在文件当中，在存放数据的时候通过序列化成json格式后再存放到文件当中，从文件中取数据是从文件中取完后通过反序列成结构体数据类型，然后对数据进行使用

```go
// 打开json文件，反序列化处账号信息
	accountjson, _ := ioutil.ReadFile(filepath)
	var account Account
	json.Unmarshal(accountjson, &account)
	fmt.Println(account)

// 序列化成json格式，存放到文件中
	data, _ := json.Marshal(newAccount)
	os.WriteFile(filepath, data, 0666)
```

## utils的使用

- 对于共同的代码，我们可以将该段代码分离出来，将所需要的功能都写到分离出来的utils目录下面，如下将一下代码分离出去

```go
// 如何判断当前路径下是否存在一个文件？
	// filepath := "./accounts/" + name + ".json"
	// _, err := os.Stat(filepath)
	// if !os.IsNotExist(err) {
	// 	fmt.Println("当前用户名已经被注册")
	// 	return
	// }

// 在utils中实现该函数
package utils

import (
	// "fmt"
	"fmt"
	"os"
	"path"
)

func IsExist(name string) (string, bool) {
	// filepath := "./accounts/" + name + ".json"
	filepath := path.Join("./accounts", fmt.Sprint(name, ".json"))
	_, err := os.Stat(filepath)
	return filepath, !os.IsNotExist(err)
}
```



## MD5的使用

在保存密码的过程中，因为使用的是明文保存，所以这里使用md5对其进行加密处理

```go
func Md5Salt(str string) string {
	salt := "猜猜我是谁"

	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte(salt))
	ret := fmt.Sprintf("%x", h.Sum(nil))
	return ret
}
```

## 日志格式和级别的设置

```go
import log
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
```

## MVC架构

- 在model中实现结构体对象的声明和结构体对象的方法
- 在service中实现程序中间逻辑
- 在main中实现与用户之间的交互

