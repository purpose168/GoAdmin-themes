#!/bin/bash

# GoAdmin 主题依赖更新脚本
# 功能：更新 go.mod 中的所有依赖包到最新版本
# 更新日期：2026-01-04

set -e  # 遇到错误时立即退出

# 输出颜色定义
GREEN="\033[0;32m"
YELLOW="\033[1;33m"
RED="\033[0;31m"
NC="\033[0m"  # No Color

echo "========================================"
echo -e "${GREEN}GoAdmin 主题依赖更新脚本${NC}"
echo "========================================"

# 检查Go是否安装
echo -e "\n${YELLOW}[检查] 验证Go环境...${NC}"
if ! command -v go &> /dev/null; then
    echo -e "${RED}错误: Go 未安装，请先安装Go${NC}"
    exit 1
fi
echo -e "${GREEN}✅ Go已安装: $(go version)${NC}"

# 检查当前目录是否有go.mod文件
echo -e "\n${YELLOW}[检查] 验证go.mod文件...${NC}"
if [ ! -f "go.mod" ]; then
    echo -e "${RED}错误: 当前目录没有go.mod文件，请在项目根目录执行此脚本${NC}"
    exit 1
fi
echo -e "${GREEN}✅ go.mod文件存在${NC}"

echo -e "\n========================================"
echo -e "${YELLOW}开始更新依赖...${NC}"
echo -e "========================================"

# 步骤 1：整理现有依赖，清理无效依赖
echo -e "\n${YELLOW}[步骤 1/3] 整理现有依赖...${NC}"
go mod tidy

# 步骤 2：更新所有直接依赖到最新版本
echo -e "\n${YELLOW}[步骤 2/3] 更新直接依赖...${NC}"
go get -u ./... || {
    echo -e "${RED}警告: 更新直接依赖时出现错误，将继续执行...${NC}"
}

# 步骤 3：再次整理依赖，确保依赖关系正确
echo -e "\n${YELLOW}[步骤 3/3] 再次整理依赖...${NC}"
go mod tidy

# 显示更新后的依赖摘要
echo -e "\n========================================"
echo -e "${GREEN}✅ 依赖更新完成！${NC}"
echo -e "========================================"
echo -e "\n📋 更新摘要："
echo -e "- 已整理现有依赖关系"
echo -e "- 已尝试更新所有直接依赖到最新版本"
echo -e "- 已确保依赖关系的正确性"
echo -e "\n📦 更新后的go.mod文件："
echo -e "----------------------------------------"
cat go.mod
echo -e "----------------------------------------"
echo -e "\n💡 建议后续操作："
echo -e "1. 运行 '${YELLOW}go build ./...${NC}' 测试编译是否成功"
echo -e "2. 运行 '${YELLOW}go test ./...${NC}' 执行所有测试"
echo -e "3. 如有问题，可使用 '${YELLOW}git diff go.mod${NC}' 查看具体变更"
echo -e "4. 如需指定版本更新，可使用 '${YELLOW}go get package@version${NC}'"
echo -e "\n${GREEN}🎉 更新脚本执行完毕！${NC}"
