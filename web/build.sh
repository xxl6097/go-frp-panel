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

  wget https://nodejs.org/dist/v20.19.0/node-v20.19.0-linux-x64.tar.xz
  tar xvf node-v20.19.0-linux-x64.tar.xz


#  rm -r /usr/local/bin/node
#  rm -r /usr/local/bin/npm
  # 创建软链接到系统路径
  sudo ln -s ./node-v20.19.0-linux-x64/bin/node /usr/local/bin/node
  sudo ln -s ./node-v20.19.0-linux-x64/bin/npm /usr/local/bin/npm
  echo "打印 node"
  node -v  # 应输出如v16.20.0
  echo "打印 npm"
  npm -v   # 应输出如8.19.4

  echo "全局安装npm-run-all"
  npm install -g npm-run-all
  echo "打印run-p版本"
  run-p --version      # 全局安装时使用
  #echo "打印npm list npm-run-all"
  #npm list npm-run-all  # 本地依赖

  buildFrpc &
  #buildFrps &
  wait  # 等待所有后台进程结束
  echo "所有任务完成"
}


buildFrpcAndFrpsAll