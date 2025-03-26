#!/bin/bash

function showBuildDir() {
  # 检查是否输入路径参数
  if [ -z "$1" ]; then
      echo "用法: $0 <路径>"
      exit 1
  fi

  # 验证路径是否存在且为目录
  if [ ! -d "$1" ]; then
      echo "错误: 路径 '$1' 不存在或不是目录！"
      exit 1
  fi

  # 获取指定路径下的所有直接子目录（非递归）
  dirs=()
  while IFS= read -r dir; do
      dirs+=("$dir")
  done < <(find "$1" -maxdepth 1 -type d ! -path "$1" | sort)

  # 检查是否有子目录
  if [ ${#dirs[@]} -eq 0 ]; then
      echo "路径 '$1' 下没有子目录！"
      exit 0
  fi

  # 生成交互式菜单
  echo "请选择要操作的目录："
  PS3="输入序号 (1-${#dirs[@]}): "
  select dir in "${dirs[@]}"; do
      if [[ -n "$dir" ]] && [[ $REPLY -ge 1 && $REPLY -le ${#dirs[@]} ]]; then
          echo "您选择的目录是: $dir"
          break
#          return $dir
      else
          echo "无效输入！请输入有效序号。"
      fi
  done
}

function build() {
    read -p "请选择：" index
      if [ $index -le 3 ]; then
        echo "无效输入！请输入有效序号。"
      fi
}

build