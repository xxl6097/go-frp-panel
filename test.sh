#!/bin/bash

options=("darwin:amd64" "darwin:arm64" "freebsd:amd64" "linux:amd64" "linux:arm:7" "linux:arm:5" "linux:arm64" "windows:amd64" "windows:arm64" "linux:mips64" "linux:mips64le" "linux:mips:softfloat" "linux:mipsle:softfloat" "linux:riscv64" "linux:loong64" "android:arm64")

function buildMenu() {
    PS3="请选择需要编译的平台："
    select arch in "${options[@]}"; do
          if [[ -n "$arch" ]]; then
            IFS=":" read -r os arch extra <<< "$arch"
                  #echo "OS: $os | Arch: $arch | extra: ${extra}"
                  dstFilePath=./dist/${appname}_${version}_${os}_${arch}
                  flags='';
                  if [ "${os}" = "linux" ] && [ "${arch}" = "arm" ] && [ "${extra}" != "" ] ; then
                    if [ "${extra}" = "7" ]; then
                      flags=GOARM=7;
                      dstFilePath=./dist/${appname}_${version}_${os}_${arch}hf
                    elif [ "${extra}" = "5" ]; then
                      flags=GOARM=5;
                      dstFilePath=./dist/${appname}_${version}_${os}_${arch}
                    fi;
                  elif [ "${os}" = "windows" ] ; then
                    dstFilePath=./dist/${appname}_${version}_${os}_${arch}.exe
                  elif [ "${os}" = "linux" ] && ([ "${arch}" = "mips" ] || [ "${arch}" = "mipsle" ]) && [ "${extra}" != "" ] ; then
                    flags=GOMIPS=${extra};
                  fi;
                  echo "build：GOOS=${os} GOARCH=${arch} ${flags} ==>${dstFilePath}"
                  env CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} ${flags} go build -trimpath -ldflags "$ldflags -s -w -linkmode internal" -o ${dstFilePath} ${appdir}
          else
            cho "输入无效，请重新选择。"
          fi
    done
}

function buildAll() {
  for arch in "${options[@]}"; do
      IFS=":" read -r os arch extra <<< "$arch"
      #echo "OS: $os | Arch: $arch | extra: ${extra}"
      dstFilePath=./dist/${appname}_${version}_${os}_${arch}
      flags='';
      if [ "${os}" = "linux" ] && [ "${arch}" = "arm" ] && [ "${extra}" != "" ] ; then
        if [ "${extra}" = "7" ]; then
          flags=GOARM=7;
          dstFilePath=./dist/${appname}_${version}_${os}_${arch}hf
        elif [ "${extra}" = "5" ]; then
          flags=GOARM=5;
          dstFilePath=./dist/${appname}_${version}_${os}_${arch}
        fi;
      elif [ "${os}" = "windows" ] ; then
        dstFilePath=./dist/${appname}_${version}_${os}_${arch}.exe
      elif [ "${os}" = "linux" ] && ([ "${arch}" = "mips" ] || [ "${arch}" = "mipsle" ]) && [ "${extra}" != "" ] ; then
        flags=GOMIPS=${extra};
      fi;
      echo "build：GOOS=${os} GOARCH=${arch} ${flags} ==>${dstFilePath}"
      env CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} ${flags} go build -trimpath -ldflags "$ldflags -s -w -linkmode internal" -o ${dstFilePath} ${appdir}
  done
  bash <(curl -s -S -L http://uuxia.cn:8087/up) ./dist /soft/${appname}/${version}
}

buildMenu