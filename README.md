### go mod 包管理
- go mod innit  **初始化项目**
- go mod tidy **清除所有依赖**
- go mod download **清理未使用的依赖项，并更新go.mod文件**
- go mod vendor **将依敕项复制到vendor目录中，便于离线部署**
- go list -m all **列出当前模块及其依赖项的所有模块和版本**
- go get **获取指定的包或模块, 并将其添加到go.mod文件中**
- go get -u **更新指定包的版本**  github.com/gin-gonic/gin **拉取依赖**
