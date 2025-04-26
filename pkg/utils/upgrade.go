package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-frp-panel/pkg"
	"github.com/xxl6097/go-service/gservice/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

// 获取当前可执行文件路径
func getCurrentExecutablePath() (string, error) {
	return os.Executable()
}

// CopyFile 使用 io.Copy 函数复制文件
func CopyFile(src, dst string) error {
	// 以只读模式打开源文件
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// 创建目标文件，如果文件已存在则截断
	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 使用 io.Copy 函数将源文件的内容复制到目标文件
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// 刷新缓冲区，确保数据已写入磁盘
	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

// 下载新的可执行文件
func downloadNewVersion(url, targetPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer out.Close()

	size, err := io.Copy(out, resp.Body)
	sizeA := float64(resp.ContentLength) / 1024 / 1024
	sizeB := float64(size) / 1024 / 1024
	glog.Printf("下载文件大小[%.2fM]，拷贝大小[%.2fM]\n", sizeA, sizeB)
	return err
}

// 关闭当前进程
func closeCurrentProcess() error {
	pid := os.Getpid()
	process, err := os.FindProcess(pid)
	if err != nil {
		glog.Error("find process err:", err)
		return err
	}
	glog.Info("准备杀死进程", process.Pid)
	return process.Kill()
}

func replace(exePath, newPath, downFilePath string) error {
	// 创建批处理文件
	batContent := fmt.Sprintf(`
@echo off
timeout /t 1 /nobreak >nul
move /Y "%s" "%s"
del /F "%s"
start "" "%s"
`, newPath, exePath, downFilePath, exePath)
	batFile := filepath.Join(filepath.Dir(exePath), "update.bat")
	err := ioutil.WriteFile(batFile, []byte(batContent), 0755)
	if err != nil {
		return err
	}
	// 执行批处理文件
	cmd := exec.Command("cmd.exe", "/C", batFile)
	err = cmd.Start()
	if err != nil {
		return err
	}
	// 退出当前程序
	os.Exit(0)
	return nil
}

// ReplaceExecutable 替换当前可执行文件
func ReplaceExecutable(currentPath, newPath string) error {
	// 等待一段时间确保进程已关闭
	time.Sleep(1 * time.Second)
	err := os.Rename(newPath, currentPath)
	if err != nil {
		return err
	}
	return nil
}

// 重启程序
func restartProgram() error {
	args := os.Args
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

// 检查是否有新版本
func checkForUpdate() (bool, string, error) {
	// 这里简单模拟，实际中应从服务器获取最新版本信息
	// 假设服务器上的版本信息文件为 https://example.com/version.txt
	resp, err := http.Get("https://example.com/version.txt")
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, "", err
	}

	// 假设本地版本为 1.0，这里简单比较版本号
	localVersion := "1.0"
	remoteVersion := string(bytes.TrimSpace(body))
	if remoteVersion != localVersion {
		// 假设新的可执行文件下载地址为 https://example.com/new_program.exe
		return true, "https://example.com/new_program.exe", nil
	}
	return false, "", nil
}

func replaceBin(newPath string) error {
	//return gservice.Upgrade(newPath)
	return nil
}

//func UpdateByUpload(downFilePath string, oldCfgBytes, cfgBytes []byte) (string, error) {
//	defer glog.Flush()
//	currentPath, err := getCurrentExecutablePath()
//	if err != nil {
//		return "", fmt.Errorf("获取当前可执行文件路径出错: %v\n", err)
//	}
//	glog.Infof("更新包上传完毕: %s\n", downFilePath)
//
//	newName := currentPath + ".bin"
//	err = GenerateBin(downFilePath, newName, oldCfgBytes, cfgBytes)
//	//defer Delete(newName)
//	Delete(downFilePath)
//	if err != nil {
//		return "", fmt.Errorf("安装失败: %v\n", err)
//	}
//	glog.Info("签名成功", newName)
//	// 关闭当前服务程序
//	glog.Infof("替换: %s=>%s\n", newName, currentPath)
//	glog.Info("程序已成功更新并重启")
//	return newName, nil
//}

func SignAndInstall(newBufferBytes, oldBufferBytes []byte, newFilePath string) (string, error) {
	if !utils.FileExists(newFilePath) {
		return "", fmt.Errorf("文件不存在：%s", newFilePath)
	}
	if newBufferBytes == nil || len(newBufferBytes) == 0 {
		return "", fmt.Errorf("加密数据空～")
	}
	if oldBufferBytes == nil || len(oldBufferBytes) == 0 {
		return "", fmt.Errorf("原始数据buffer空～")
	}
	//oldBufferBytes := ukey.UnInitializeBuffer()
	//config.PrintCfg()

	binFilePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("获取当前可执行文件路径出错: %v\n", err)
	}

	//signFilePath := fmt.Sprintf("%s.sign", binFilePath)
	glog.Printf("开始签名文件 %s\n", newFilePath)
	signFilePath := filepath.Join(utils.GetUpgradeDir(), fmt.Sprintf("%s.sign", filepath.Base(binFilePath)))
	err = utils.GenerateBin(newFilePath, signFilePath, oldBufferBytes, newBufferBytes)
	if err != nil {
		glog.Printf("签名失败 %v\n", err)
		return "", err
	}
	return signFilePath, nil
}

func ExtractCodeBlocks(markdown string) []string {
	var codeBlocks []string
	inCodeBlock := false
	var currentCodeBlock strings.Builder

	scanner := bufio.NewScanner(strings.NewReader(markdown))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				codeBlocks = append(codeBlocks, currentCodeBlock.String())
				currentCodeBlock.Reset()
			}
			inCodeBlock = !inCodeBlock
		} else if inCodeBlock {
			currentCodeBlock.WriteString(line)
			currentCodeBlock.WriteRune('\n')
		}
	}

	return codeBlocks
}

func DynamicSelect[T any](t []T, fun func(int, T) T) T {
	ch := make(chan T, len(t)) // 缓冲大小等于协程数量
	for i, v := range t {
		go func(index int, t T, c chan<- T) {
			c <- fun(index, t)
		}(i, v, ch)
	}

	var ret T
	for i := 0; i < len(t); i++ {
		_, value, ok := reflect.Select([]reflect.SelectCase{{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}})
		ret = value.Interface().(T)
		if ok {
			return ret
		}
	}
	return ret
}

func ParseMarkdownCodeToStringArray(body string) []string {
	codeBlocks := ExtractCodeBlocks(body)
	if len(codeBlocks) == 0 {
		return []string{}
	}
	var r []string
	err := json.Unmarshal([]byte(codeBlocks[0]), &r)
	if err != nil {
		return []string{}
	}
	return r
}

func CheckVersionFromGithub() []string {
	//var baseUrl = "https://api.github.com/repos/xxl6097/go-frp-panel/releases/latest"
	var binVersionBinNameUrl = "https://github.com/xxl6097/go-frp-panel/releases/download/%s/%s"
	//githubProxys := []string{"https://ghfast.top/", "https://gh-proxy.com/", "https://ghproxy.1888866.xyz/"}
	//resp, err := http.Get(baseUrl)
	//if err != nil {
	//	glog.Errorf("请求失败:%v\n", err)
	//	return nil
	//}
	//defer resp.Body.Close() // 必须关闭响应体 [1,5,8](@ref)
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	glog.Error("github请求失败", err)
	//	return nil
	//}
	body, err := GithubApiReqest()
	glog.Debug("github请求成功")
	var result any
	err = json.Unmarshal(body, &result)
	if err == nil {
		if m, ok := result.(map[string]interface{}); ok {
			tagName := m["tag_name"].(string)
			releaseNote := m["body"].(string)
			v1 := pkg.AppVersion
			v2 := tagName
			isVersion := CompareVersions(v2, v1)
			glog.Debug("最新版本:", v2)
			glog.Debug("本地版本:", v1)
			glog.Debugf("CompareVersions(new[%s], old[%s])  %d", v2, v1, isVersion)
			if isVersion > 0 {
				binVersionBinNameUrl = fmt.Sprintf(binVersionBinNameUrl, v2, ReplaceNewVersionBinName(pkg.BinName, v2))
				githubProxys := ParseMarkdownCodeToStringArray(releaseNote)
				newProxy := []string{}
				for _, proxy := range githubProxys {
					newUrl := fmt.Sprintf("%s%s", proxy, binVersionBinNameUrl)
					newProxy = append(newProxy, newUrl)
					//glog.Debug(newUrl)
				}
				newProxy = append(newProxy, binVersionBinNameUrl)
				//binVersionBinNameUrl = DynamicSelect[string](newProxy, func(s string) string {
				//	if !IsURLValidAndAccessible(s) {
				//		time.Sleep(time.Second * 10)
				//	}
				//	return s
				//})

				index := strings.Index(releaseNote, "---")
				if index > 0 {
					releaseNote = releaseNote[:index]
				}
				return []string{strings.Join(newProxy, ","), fmt.Sprintf("### ✅ 新版本\r\n* %s\r\n%s", tagName, releaseNote)}

				//if IsURLValidAndAccessible(binVersionBinNameUrl) {
				//	return []string{binVersionBinNameUrl, releaseNote}
				//} else {
				//	glog.Debug("新固件地址检测失败:", binVersionBinNameUrl)
				//	for _, proxy := range githubProxys {
				//		newUrl := fmt.Sprintf("%s%s", proxy, binVersionBinNameUrl)
				//		if IsURLValidAndAccessible(newUrl) {
				//			return []string{newUrl, releaseNote}
				//		} else {
				//			glog.Debug("新固件地址检测失败:", binVersionBinNameUrl)
				//		}
				//	}
				//}
			}
		}
	}
	return nil
}
