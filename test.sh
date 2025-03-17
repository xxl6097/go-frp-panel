#!/bin/bash

# 获取所有分支列表（包含远程分支）
git fetch origin > /dev/null 2>&1
branches=($(git branch -a | grep -v "HEAD" | sed 's/^* //' | sed 's/remotes\///'))

# 生成分支菜单
echo "可更新的分支列表："
select branch in "${branches[@]}"; do
    if [[ -n "$branch" ]]; then
        echo "正在更新分支：$branch"
        git checkout "$branch" > /dev/null 2>&1
        git pull origin "$branch"
        break
    else
        echo "输入无效，请重新选择。"
    fi
done