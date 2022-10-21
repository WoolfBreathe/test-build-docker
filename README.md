# Github Action Build Docker Image

## 使用

1. 编辑image.json文件

2. 编译打包 SET CGO_ENABLED=0&&SET GOOS=linux&&SET GOARCH=amd64&&go build -ldflags="-s -v" -o build_shell

3. 上传分支