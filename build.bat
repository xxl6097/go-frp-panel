@echo off
rem Turn off command echoing in the command prompt so that only the output results are shown.
echo Current operating system: Windows
set "sourceFile=./cmd/service/service.go"
set "APP_NAME=acfrps"
set "PKG=github.com/xxl6097/go-frp-panel/pkg"
set "DISPLAYNAME=AcFrps网络代理程序"
set "DESRIBTION=一款基于GO语言的网络代理服务程序"

:menu
cls
echo Please select the target system and architecture to compile:
echo 1. Windows amd64
echo 2. Windows arm64
echo 3. Linux amd64
echo 4. Linux arm64
echo 5. macOS amd64
echo 6. macOS arm64
echo 7. Compile all
echo 8. Exit
set /p choice=Enter the option number:

rem Perform corresponding actions based on the user's choice.
if "%choice%"=="1" (
    call :build windows amd64
    goto end
) else if "%choice%"=="2" (
    call :build windows arm64
    goto end
) else if "%choice%"=="3" (
    call :build linux amd64
    goto end
) else if "%choice%"=="4" (
    call :build linux arm64
    goto end
) else if "%choice%"=="5" (
    call :build darwin amd64
    goto end
) else if "%choice%"=="6" (
    call :build darwin arm64
    goto end
) else if "%choice%"=="7" (
    call :build windows amd64
    call :build windows arm64
    call :build linux amd64
    call :build linux arm64
    call :build darwin amd64
    call :build darwin arm64
    goto end
) else if "%choice%"=="8" (
    exit /b
) else (
    echo Invalid option. Please try again.
    pause
    goto menu
)

:build
    set "A=-X '%PKG%.AppName=%APP_NAME%'"
    set "B=-X ^'%PKG%.DisplayName=%DISPLAYNAME%'"
    set "C=-X '%PKG%.Description=%DESRIBTION%'"
    set "D=-X '%PKG%.BuildTime=%date% %time%'"
    set LDFLAGS=-ldflags "-s -w %A% %B% %C% %D%"
    RD /S /Q "./dist"
    set "baseOutputName=./dist/%APP_NAME%_%1_%2"
    if "%1"=="windows" (
        set "outputName=%baseOutputName%.exe"
    ) else (
        set "outputName=%baseOutputName%"
    )
    echo Compiling %1/%2...
    rem Set the GOOS and GOARCH environment variables and execute the compilation command.
    (
        set "GOOS=%1"
        set "GOARCH=%2"
        rem echo "go build %LDFLAGS% -o %outputName% %sourceFile%"
        go build %LDFLAGS% -o %outputName% %sourceFile%
    )
    rem Check if the compilation was successful.
    if %errorlevel% equ 0 (
        echo Compilation of %1/%2 succeeded. Generated file: %outputName%
    ) else (
        echo Compilation of %1/%2 failed.
    )
    rem pause
    exit /b

:end
exit /b