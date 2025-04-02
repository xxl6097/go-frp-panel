#!/bin/bash

API_URL="https://api.github.com/repos/xxl6097/go-frp-panel/releases/latest"
DOWNLOAD_URL=$(curl -s $API_URL | grep "browser_download_url" | grep "acfrp" | cut -d'"' -f4)

echo "--->${DOWNLOAD_URL}"

