#!/bin/bash
GITHUB_USER="xxl6097"
REPO_NAME="go-frp-panel"
TOKEN="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"

# 获取所有 Release ID 列表
releases=$(curl -s -H "Authorization: token $TOKEN" \
  "https://api.github.com/repos/$GITHUB_USER/$REPO_NAME/releases" | jq -r '.[].id')

# 循环删除
for release_id in $releases; do
  echo "Deleting Release ID: $release_id"
  curl -X DELETE -H "Authorization: token $TOKEN" \
    "https://api.github.com/repos/$GITHUB_USER/$REPO_NAME/releases/$release_id"
done