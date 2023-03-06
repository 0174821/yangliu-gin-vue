# yangliu-gin-vue

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


