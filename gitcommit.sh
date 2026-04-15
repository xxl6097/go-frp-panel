#!/bin/bash
git_url="${SHELL_HOST:-https://cnb.cool/clife/golang/cp/-/git/raw/main}/gitcommit.sh"
bash <(curl -sSL "$git_url")
