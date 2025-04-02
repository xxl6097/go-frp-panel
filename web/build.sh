#!/bin/bash

function buildFrpc() {
  cd ./web/frpc && npm run build
  echo "frpc编译完成"
  pwd
  ls -lh
}

function buildFrps() {
  cd ./web/frpc && npm run build
  echo "frps编译完成"
  pwd
  ls -lh
}

function buildFrpcAndFrpsAll() {
#  sudo apt update && sudo apt install make
#  make --version
  # 下载并解压Node.js
#  wget https://nodejs.org/dist/v18.12.0/node-v18.12.0-linux-x64.tar.xz
#  tar xvf node-v18.12.0-linux-x64.tar.xz

  ver=18.12.0
  wget -q https://nodejs.org/dist/v${ver}/node-v${ver}-linux-x64.tar.xz
  tar xf node-v${ver}-linux-x64.tar.xz

  sudo rm -r ./node
  sudo mv -f node-v${ver}-linux-x64 node

  ls -lh ./node/bin

  sudo rm -r /usr/bin/node
  sudo rm -r /usr/bin/npm

  # 创建软链接到系统路径
  sudo ln -s ./node/bin/node /usr/bin/node
  sudo ln -s ./node/bin/npm /usr/bin/npm
  source ~/.bashrc
  npm install -g n
  n 18.12.0
  echo "打印 node"
  node -v  # 应输出如v16.20.0
  echo "打印 npm"
  npm -v   # 应输出如8.19.4

  echo "全局安装npm-run-all"
  npm install -g npm-run-all
  echo "打印run-p版本"
  run-p --version      # 全局安装时使用
  echo "全局安装 vite"
  npm install -g vite
  echo "打印 vite"
  vite --version  # 全局安装
  echo "全局安装 vue-tsc"
  npm install -g vue-tsc
  echo "打印 vue-tsc"
  vue-tsc --version
  #echo "打印npm list npm-run-all"
  #npm list npm-run-all  # 本地依赖

  buildFrpc &
  #buildFrps &
  wait  # 等待所有后台进程结束
  echo "所有任务完成"
}


buildFrpcAndFrpsAll