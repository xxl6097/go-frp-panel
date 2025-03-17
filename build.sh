#!/bin/bash
module=$(grep "module" go.mod | cut -d ' ' -f 2)
#appname=$(basename $module)
#appname="acfrps"
#appdir="./cmd/frps"
#DisplayName="AcFrps网络代理程序"
#Description="一款基于GO语言的网络代理服务程序"
#builddir="./dist"
options=("windows:amd64" "windows:arm64" "linux:amd64" "linux:arm64" "linux:arm:7" "linux:arm:5" "linux:mips64" "linux:mips64le" "linux:mips:softfloat" "linux:mipsle:softfloat" "linux:riscv64" "linux:loong64" "darwin:amd64" "darwin:arm64" "freebsd:amd64" "android:arm64")
#options=("windows:amd64" "windows:arm64" "linux:amd64" "linux:arm64")
version=$(git tag -l "[0-99]*.[0-99]*.[0-99]*" --sort=-creatordate | head -n 1)
versionDir="$module/pkg"

function writeVersionGoFile() {
  if [ ! -d "./pkg" ]; then
    mkdir "./pkg"
  fi
cat <<EOF > ./pkg/version.go
package pkg
import (
	"fmt"
	"strings"
)
var (
	AppName      string // 应用名称
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
	DisplayName  string // 服务显示名
	Description  string // 服务描述信息
)
// Version 版本信息
func Version() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("App Name:\t%s\n", AppName))
	sb.WriteString(fmt.Sprintf("App Version:\t%s\n", AppVersion))
	sb.WriteString(fmt.Sprintf("Build version:\t%s\n", BuildVersion))
	sb.WriteString(fmt.Sprintf("Build time:\t%s\n", BuildTime))
	sb.WriteString(fmt.Sprintf("Git revision:\t%s\n", GitRevision))
	sb.WriteString(fmt.Sprintf("Git branch:\t%s\n", GitBranch))
	sb.WriteString(fmt.Sprintf("Golang Version: %s\n", GoVersion))
	sb.WriteString(fmt.Sprintf("DisplayName:\t%s\n", DisplayName))
	sb.WriteString(fmt.Sprintf("Description:\t%s\n", Description))
	fmt.Println(sb.String())
	return sb.String()
}
EOF
}

# builddir：输出目录
# appname：应用名称
# version：应用版本
# appdir：main.go目录
# disname：显示名
# describe：描述
function buildMenu() {
  builddir=$1
  appname=$2
  version=$3
  appdir=$4
  disname=$5
  describe=$6
  ldflags=$(buildLdflags $appname $disname $describe)
  PS3="请选择需要编译的平台："
  select arch in "${options[@]}"; do
      if [[ -n "$arch" ]]; then
        IFS=":" read -r os arch extra <<< "$arch"
          dstFilePath=${builddir}/${appname}_${version}_${os}_${arch}
          flags='';
          if [ "${os}" = "linux" ] && [ "${arch}" = "arm" ] && [ "${extra}" != "" ] ; then
            if [ "${extra}" = "7" ]; then
              flags=GOARM=7;
              dstFilePath=${builddir}/${appname}_${version}_${os}_${arch}hf
            elif [ "${extra}" = "5" ]; then
              flags=GOARM=5;
              dstFilePath=${builddir}/${appname}_${version}_${os}_${arch}
            fi;
          elif [ "${os}" = "windows" ] ; then
            dstFilePath=${builddir}/${appname}_${version}_${os}_${arch}.exe
          elif [ "${os}" = "linux" ] && ([ "${arch}" = "mips" ] || [ "${arch}" = "mipsle" ]) && [ "${extra}" != "" ] ; then
            flags=GOMIPS=${extra};
          fi;
          echo "build：GOOS=${os} GOARCH=${arch} ${flags} ==>${dstFilePath}"
          env CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} ${flags} go build -trimpath -ldflags "$ldflags -s -w -linkmode internal" -o ${dstFilePath} ${appdir}
          return $?
      else
        echo "输入无效，请重新选择。"
      fi
  done
}

# builddir：输出目录
# appname：应用名称
# version：应用版本
# appdir：main.go目录
# disname：显示名
# describe：描述
function buildAll() {
  builddir=$1
  appname=$2
  version=$3
  appdir=$4
  disname=$5
  describe=$6
  ldflags=$(buildLdflags $appname $disname $describe)
  for arch in "${options[@]}"; do
      IFS=":" read -r os arch extra <<< "$arch"
      #echo "OS: $os | Arch: $arch | extra: ${extra}"
      dstFilePath=${builddir}/${appname}_${version}_${os}_${arch}
      flags='';
      if [ "${os}" = "linux" ] && [ "${arch}" = "arm" ] && [ "${extra}" != "" ] ; then
        if [ "${extra}" = "7" ]; then
          flags=GOARM=7;
          dstFilePath=${builddir}/${appname}_${version}_${os}_${arch}hf
        elif [ "${extra}" = "5" ]; then
          flags=GOARM=5;
          dstFilePath=${builddir}/${appname}_${version}_${os}_${arch}
        fi;
      elif [ "${os}" = "windows" ] ; then
        dstFilePath=${builddir}/${appname}_${version}_${os}_${arch}.exe
      elif [ "${os}" = "linux" ] && ([ "${arch}" = "mips" ] || [ "${arch}" = "mipsle" ]) && [ "${extra}" != "" ] ; then
        flags=GOMIPS=${extra};
      fi;
      echo "build：GOOS=${os} GOARCH=${arch} ${flags} ==>${dstFilePath}"
      env CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} ${flags} go build -trimpath -ldflags "$ldflags -s -w -linkmode internal" -o ${dstFilePath} ${appdir}
  done
}

function build() {
  #echo "---->$1 $2 $3 $4 $5 $6 $7"
  if [ $7 -eq 1 ]; then
    buildMenu $1 $2 $3 $4 $5 $6
  else
    buildAll $1 $2 $3 $4 $5 $6
  fi
}

function upgradeVersion() {
  if [ "$version" = "" ]; then
    version="0.0.0"
  else
    v3=$(echo $version | awk -F'.' '{print($3);}')
    v2=$(echo $version | awk -F'.' '{print($2);}')
    v1=$(echo $version | awk -F'.' '{print($1);}')
    if [[ $(expr $v3 \>= 99) == 1 ]]; then
      v3=0
      if [[ $(expr $v2 \>= 99) == 1 ]]; then
        v2=0
        v1=$(expr $v1 + 1)
      else
        v2=$(expr $v2 + 1)
      fi
    else
      v3=$(expr $v3 + 1)
    fi
    version="$v1.$v2.$v3"
  fi
}


function buildLdflags() {
  #os_name=$(uname -s)
  #echo "os type $os_name"
  appname=$1
  DisplayName=$2
  Description=$3
  APP_NAME=${appname}
  BUILD_VERSION=$(if [ "$(git describe --tags --abbrev=0 2>/dev/null)" != "" ]; then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
  BUILD_TIME=$(TZ=Asia/Shanghai date "+%Y-%m-%d %H:%M:%S")
  GIT_REVISION=$(git rev-parse --short HEAD)
  GIT_BRANCH=$(git name-rev --name-only HEAD)
  GO_VERSION=$(go version)
  # shellcheck disable=SC2089
  local ldflags="-s -w\
 -X '${versionDir}.DisplayName=${DisplayName}_v${version}'\
 -X '${versionDir}.Description=${Description}'\
 -X '${versionDir}.AppName=${APP_NAME}'\
 -X '${versionDir}.AppVersion=${version}'\
 -X '${versionDir}.BuildVersion=${BUILD_VERSION}'\
 -X '${versionDir}.BuildTime=${BUILD_TIME}'\
 -X '${versionDir}.GitRevision=${GIT_REVISION}'\
 -X '${versionDir}.GitBranch=${GIT_BRANCH}'\
 -X '${versionDir}.GoVersion=${GO_VERSION}'"
  echo "$ldflags"
}

function initCommArgs() {
  upgradeVersion
  #echo "version:${version}"
  writeVersionGoFile
}

function tagAndGitPush() {
    git add .
    git commit -m "release ${version}"
    git tag -a $version -m "release ${version}"
    git push origin $version
}

function upload() {
  builddir=$1
  appname=$2
  version=$3
  ls ${builddir}
  if [ $? -eq 0 ]; then
      echo "上传文件 ${builddir} /soft/${appname}/${version}"
      bash <(curl -s -S -L http://uuxia.cn:8087/up) ${builddir} /soft/${appname}/${version}
  else
      echo "上传失败，错误码: $?"  # 输出错误信息（例如返回2表示文件未找到）
  fi
}

function gitCommit() {
  if [ $? -eq 0 ]; then
      echo "编译成功，git提交代码..."
      tagAndGitPush
  else
      echo "编译失败，错误码: $?"  # 输出错误信息（例如返回2表示文件未找到）
  fi
}

function buildFrpc() {
    appname="acfrpc"
    appdir="./cmd/frpc"
    DisplayName="AcFrpc网络代理程序"
    Description="一款基于GO语言的网络代理服务程序"
    builddir="./dist/frpc"
    rm -rf ${builddir}
    build $builddir $appname "$version" $appdir $DisplayName $Description "$1"
    upload $builddir $appname "$version"
}

function buildFrps() {
    appname="acfrps"
    appdir="./cmd/frps"
    DisplayName="AcFrps网络代理程序"
    Description="一款基于GO语言的网络代理服务程序"
    builddir="./dist/frps"
    rm -rf ${builddir}
    build $builddir $appname "$version" $appdir $DisplayName $Description "$1"
    upload $builddir $appname "$version"
}

function buildFrpcAndFrpsAll() {
  buildFrpc 2 &
  buildFrps 2 &
  wait  # 等待所有后台进程结束
  builddir="./dist"
  echo "所有任务完成"
}

function buildFrpcMenu() {
  clear
  echo "1、Frpc编译菜单"
  echo "2、编译全部"
  read -p "请选择：" index
  buildFrpc $index
}

function buildFrpsMenu() {
  clear
  echo "1、Frps编译菜单"
  echo "2、编译全部"
  read -p "请选择：" index
  buildFrps $index
}

function main() {
  initCommArgs
  echo "1、编译Frps"
  echo "2、编译Frpc"
  echo "3、编译全部"
  read -p "请选择：" index
  if [ $index == 1 ]; then
    buildFrpsMenu
  elif [ $index == 2 ]; then
    buildFrpcMenu
  else
    buildFrpcAndFrpsAll
  fi
  gitCommit
}

main