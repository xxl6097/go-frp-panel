#!/bin/bash

function buildFrpc() {
  cd frpc && make build
  echo "frpc编译完成"
  pwd
  ls -lh
}

function buildFrps() {
  cd frps && make build
  echo "frps编译完成"
  pwd
  ls -lh
}

function buildFrpcAndFrpsAll() {
  buildFrpc &
  buildFrps &
  wait  # 等待所有后台进程结束
  echo "所有任务完成"
}


buildFrpcAndFrpsAll