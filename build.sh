#!/bin/bash
module=$(grep "module" go.mod | cut -d ' ' -f 2)
#appname=$(basename $module)
#appname="acfrps"
#appdir="./cmd/frps"
#DisplayName="AcFrps网络代理程序"
#Description="一款基于GO语言的网络代理服务程序"
#builddir="./release"
options=("windows:amd64" "windows:arm64" "linux:amd64" "linux:arm64" "linux:arm:7" "linux:arm:5" "linux:mips64" "linux:mips64le" "linux:mips:softfloat" "linux:mipsle:softfloat" "linux:riscv64" "linux:loong64" "darwin:amd64" "darwin:arm64" "freebsd:amd64" "android:arm64")
#options=("windows:amd64" "linux:amd64")
version=$(git tag -l "v[0-99]*.[0-99]*.[0-99]*" --sort=-creatordate | head -n 1)
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
	"runtime"
)
func init() {
	OsType = runtime.GOOS
	Arch = runtime.GOARCH
}
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
	OsType       string // 操作系统
	Arch         string // cpu类型
	BinName      string // 运行文件名称，包含平台架构
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
	sb.WriteString(fmt.Sprintf("OsType:\t%s\n", OsType))
	sb.WriteString(fmt.Sprintf("Arch:\t%s\n", Arch))
	sb.WriteString(fmt.Sprintf("BinName:\t%s\n", BinName))
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
          #echo "env CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} ${flags} go build"
          filename=$(basename "$dstFilePath")  # 输出 "file.txt"
          binName="-X '${versionDir}.BinName=${filename}'"
          env CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} ${flags} go build -trimpath -ldflags "$ldflags $binName -linkmode internal" -o ${dstFilePath} ${appdir}
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
      #echo "env CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} ${flags} go build"
      filename=$(basename "$dstFilePath")  # 输出 "file.txt"
      # shellcheck disable=SC2089
      binName="-X '${versionDir}.BinName=${filename}'"
      #echo "--->env CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} ${flags} go build -trimpath -ldflags "$ldflags $binName -linkmode internal" -o ${dstFilePath} ${appdir}"
      env CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} ${flags} go build -trimpath -ldflags "$ldflags $binName -linkmode internal" -o ${dstFilePath} ${appdir}
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
  version=$(increment_version "$version")
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

function buildLdflags() {
  #os_name=$(uname -s)
  #echo "os type $os_name"
  appname=$1
  DisplayName=$2
  Description=$3
  APP_NAME=${appname}
  #BUILD_VERSION=$(if [ "$(git describe --tags --abbrev=0 2>/dev/null)" != "" ]; then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
  BUILD_TIME=$(TZ=Asia/Shanghai date "+%Y-%m-%d %H:%M:%S")
  GIT_REVISION=$(git rev-parse --short HEAD)
  #GIT_BRANCH=$(git name-rev --name-only HEAD)
  #GIT_BRANCH=$(git tag -l "v[0-99]*.[0-99]*.[0-99]*" --sort=-creatordate | head -n 1)
  GO_VERSION=$(go version)
  # shellcheck disable=SC2089
  local ldflags="-s -w\
 -X '${versionDir}.DisplayName=${DisplayName}_${version}'\
 -X '${versionDir}.Description=${Description}'\
 -X '${versionDir}.AppName=${APP_NAME}'\
 -X '${versionDir}.AppVersion=${version}'\
 -X '${versionDir}.BuildVersion=${version}'\
 -X '${versionDir}.BuildTime=${BUILD_TIME}'\
 -X '${versionDir}.GitRevision=${GIT_REVISION}'\
 -X '${versionDir}.GitBranch=${version}'\
 -X '${versionDir}.GoVersion=${GO_VERSION}'"
  echo "$ldflags"
}


function push() {
  git add .
  git commit -m "$(date '+%Y-%m-%d %H:%M:%S') by ${USER}"
  echo "提交代码"
  git push
}

function quickTagAndPush() {
  git add .
  git commit -m "release ${version}"
  git tag -a $version -m "release ${version}"
  git push origin $version
  push
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
      #quickTagAndPush
      push
  else
      echo "编译失败，错误码: $?"  # 输出错误信息（例如返回2表示文件未找到）
  fi
}

function buildFrpc() {
    appname="acfrpc"
    appdir="./cmd/frpc"
    DisplayName="AcFrpc网络代理程序"
    Description="一款基于GO语言的网络代理服务程序"
    builddir="./release/frpc"
    rm -rf ${builddir}
    build $builddir $appname "$version" $appdir $DisplayName $Description "$1"
    #upload $builddir $appname "$version"
}

function buildFrps() {
    appname="acfrps"
    appdir="./cmd/frps"
    DisplayName="AcFrps网络代理程序"
    Description="一款基于GO语言的网络代理服务程序"
    builddir="./release/frps"
    rm -rf ${builddir}
    build $builddir $appname "$version" $appdir $DisplayName $Description "$1"
    #upload $builddir $appname "$version"
}

function buildFrpcAndFrpsAll() {
  rm -rf ${builddir}
  buildFrpc 2 &
  buildFrps 2 &
  wait  # 等待所有后台进程结束
  builddir="./release"
  echo "所有任务完成"
}

function buildFrpcAndFrpsAllForGithubRelease() {
  echo "===>version:${version}"
  buildFrpcAndFrpsAll
  mkdir -p ./release/packages
  cp -f ./release/frpc/* ./release/packages
  cp -f ./release/frps/* ./release/packages
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

function github_release() {
    REPO="xxl6097/go-frp-panel"  # 替换为你的GitHub仓库
    TAG="${version}"  # 替换为你的标签
    RELEASE_NAME="${version}"  # 替换为你的发布名称
    DESCRIPTION="基于GO语言的网络代理服务程序"  # 替换为你的发布描述
    TOKEN=$(cat .token)  # 替换为你的GitHub Token
    # 定义要扫描的目录
    DIRECTORY="./release"
    # 初始化一个空数组
    FILES=()
    # 使用find命令扫描目录，并将结果添加到数组中
    while IFS= read -r file; do
        FILES+=("$file")
    done < <(find "$DIRECTORY" -type f)
    # 打印数组内容
#    echo "Found files:"
#    printf '%s\n' "${FILES[@]}"

    # 创建一个新的release
    response=$(curl -s -X POST \
      -H "Authorization: token $TOKEN" \
      -H "Accept: application/vnd.github.v3+json" \
      https://api.github.com/repos/$REPO/releases \
      -d "{
        \"tag_name\": \"$TAG\",
        \"target_commitish\": \"main\",
        \"name\": \"$RELEASE_NAME\",
        \"body\": \"$DESCRIPTION\",
        \"draft\": false,
        \"prerelease\": false
      }")

    # 提取release的上传URL
    upload_url=$(echo "$response" | jq -r .upload_url | sed -e "s/{?name,label}//")

    # 检查创建release是否成功
    if [ "$upload_url" == "null" ]; then
      echo "Failed to create release"
      echo "$response"
      exit 1
    fi

    # 上传附件文件
    for FILE_PATH in "${FILES[@]}"; do
      FILE_NAME=$(basename "$FILE_PATH")
      echo "Uploading $FILE_NAME..."
      curl -s -X POST \
        -H "Authorization: token $TOKEN" \
        -H "Content-Type: $(file -b --mime-type "$FILE_PATH")" \
        --data-binary @"$FILE_PATH" \
        "$upload_url?name=$FILE_NAME"
      echo "$FILE_NAME uploaded successfully."
    done

    echo "All files uploaded successfully."
}


function buildAllUploadGithub() {
  github_release
}

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
# shellcheck disable=SC2120
function buildDir() {
  showBuildDir ./cmd
  builddir="./release/${dir}"
  appname=$(basename "$dir")
  appdir=${dir}
  disname="${dir}应用程序"
  describe="一款基于GO语言的${dir}程序"
  rm -rf ${builddir}
  buildMenu $builddir $appname "$version" $appdir $disname $describe
}

function main() {
  upgradeVersion
  echo "1、编译Frps"
  echo "2、编译Frpc"
  echo "3、编译全部"
  echo "4、编译目录"
  read -p "请选择：" index
  if [ $index == 1 ]; then
    buildFrpsMenu
  elif [ $index == 2 ]; then
    buildFrpcMenu
  elif [ $index == 3 ]; then
    buildFrpcAndFrpsAll
  elif [ $index == 4 ]; then
    buildDir
  fi
  #提交代码
#  if [ $index -le 3 ]; then
#      gitCommit
#  fi
  #gitCommit
}

function buildWeb() {
  chmod +x ./web/build.sh
  ./web/build.sh
}
function bootstrap() {
  if [ $# -ge 2 ] && [ -n "$2" ]; then
    version=$2
  fi
  writeVersionGoFile
  case $1 in
  all) (buildFrpcAndFrpsAllForGithubRelease) ;;
    *) (main)  ;;
  esac
}

bootstrap $1 $2