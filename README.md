# open-api-tool
接口包初始化工具

## 安装 

```bash
go install github.com/open-api-go/open-api-tool@latest
```

## 使用

```bash
open-api-tool new -n {项目名} -d {项目目录}
```

## 建议

最新版本的protoc-gen-go要求go_package必须含有/，且会生成到$GOPATH/src目录下，所以建议把工程文件放到$GOPATH/src/git域名/git_group/目录下。

如 https://github.com/open-api-go/wechat-mp 则该工程文件为 $GOPATH/src/github.com/open-api-go/wechat-mp