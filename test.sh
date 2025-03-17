#!/bin/bash

function check_status() {
    grep "error" /var/log/wifi.log  # 假设函数内执行某个操作
    return $?  # 显式返回上一个命令的状态
}
function check_direct() {
    ls /nonexistent_directory  # 执行一个可能失败的命令
    if [ $? -eq 0 ]; then
        echo "目录存在"
    else
        echo "目录不存在，错误码: $?"  # 输出错误信息（例如返回2表示文件未找到）
    fi
}
check_direct
echo "函数返回状态: $?"  # 输出grep命令的结果（例如0表示找到匹配，1表示未找到）