# 默认目标：编译所有主题并格式化代码
# 执行顺序：
# 1. 编译 adminlte 主题
# 2. 编译 sword 主题
# 3. 格式化所有 Go 代码
all:
	make -C ./adminlte
	make -C ./sword
	make fmt

# 代码格式化目标
# 功能：
# - 使用 go fmt 格式化所有 Go 代码
# - 使用 goimports 整理导入语句并格式化代码
# 参数说明：
# - GO111MODULE=off：关闭 Go 模块模式
# - -l：列出需要格式化的文件
# - -w：直接修改文件而不输出到标准输出
fmt:
	GO111MODULE=off go fmt ./...
	GO111MODULE=off goimports -l -w .