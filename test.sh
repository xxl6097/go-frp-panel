#!/bin/bash

# 定义上传的文件路径
FILE_PATH="/Users/uuxia/Downloads/ClashX.dmg"
# 定义上传的目标地址
UPLOAD_URL="http://uuxia.cn:8087/soft/acfrps/v0.0.11"

# 使用 curl 上传文件并显示进度
curl  -# -H "Authorization: Basic YWRtaW46aGV0MDAyNDAy" -F "file=@$FILE_PATH" "$UPLOAD_URL"

