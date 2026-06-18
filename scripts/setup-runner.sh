#!/bin/bash
set -e

# ============================================================
# HelloBlog - GitHub Actions 自托管 Runner 安装脚本
# 在阿里云服务器上运行此脚本
# 解决服务器无法连接 github.com 拉取代码的问题
# ============================================================

REPO="greenpea30/HelloBlog"
RUNNER_VERSION="2.322.0"

echo "🔍 1/4 创建 runner 目录..."
mkdir -p ~/actions-runner && cd ~/actions-runner

echo "🔍 2/4 下载 GitHub Actions Runner..."
curl -o actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz \
  -L https://github.com/actions/runner/releases/download/v${RUNNER_VERSION}/actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz

echo "🔍 3/4 解压..."
tar xzf actions-runner-linux-x64-${RUNNER_VERSION}.tar.gz

echo ""
echo "============================================"
echo "✅ 下载完成！接下来需要配置并启动 Runner"
echo "============================================"
echo ""
echo "请打开以下网址获取配置命令："
echo "  https://github.com/${REPO}/settings/actions/runners/new"
echo ""
echo "然后依次执行："
echo "  1. cd ~/actions-runner"
echo "  2. 粘贴上面网址中的 ./config.sh 命令"
echo "  3. ./run.sh  （测试运行，Ctrl+C 停止）"
echo "  4. 安装为服务：sudo ./svc.sh install && sudo ./svc.sh start"
echo ""
echo "配置完成后，回到 VS Code 继续后续步骤。"
