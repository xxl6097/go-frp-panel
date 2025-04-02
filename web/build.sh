#!/bin/bash

function buildFrpc() {
  cd ./web/frpc && make build
  echo "frpc编译完成"
  pwd
  ls -lh
}

function buildFrps() {
  cd ./web/frps && make build
  echo "frps编译完成"
  pwd
  ls -lh
}

function buildFrpcAndFrpsAll() {
  sudo apt update && sudo apt install make
  make --version
  curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
  source ~/.bashrc  # 或 source ~/.zshrc
  nvm install node
  nvm install 18.11.12
  npm install npm-run-all --save-dev
  npx run-p --version  # 本地安装时使用
  run-p --version      # 全局安装时使用
  npm list npm-run-all  # 本地依赖
  node -v  # 输出 v16.20.0
  npm -v   # 自动匹配对应版本（如 npm 8.19.4）
  buildFrpc &
  buildFrps &
  wait  # 等待所有后台进程结束
  echo "所有任务完成"
}


buildFrpcAndFrpsAll