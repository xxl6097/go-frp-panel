#!/bin/bash


function upgradeVersion() {
  version="v0.0.8"
  version=$(increment_version "$version")
  echo "-->$version"
}

function increment_version() {
    local version_part=$1
    if [ "$version_part" = "" ]; then
      version_part="v0.0.0"
    fi
    local prefix="${version_part%%[0-9.]*}"  # 提取前缀（删除数字/点后的所有内容）
    local version="${version_part#$prefix}"  # 提取版本号（删除前缀后的剩余部分）
    # 分割版本号
    IFS='.' read -ra parts <<< "$version"
    local major=${parts[0]}
    local minor=${parts[1]}
    local patch=${parts[2]}
    patch=$((patch + 1))
    if [[ $patch -ge 100 ]]; then
        minor=$((minor + 1))
        patch=0
        # 检查次版本是否需要进位
        if [[ $minor -ge 100 ]]; then
            major=$((major + 1))
            minor=0
        fi
    fi
    # 重组并返回新版本号
    echo "${prefix}${major}.${minor}.${patch}"
}

upgradeVersion

