package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"io"
	"net/http"
	"os"
)

func DownLoadGeneric(baseUrl, username, password, destPath string) error {
	glog.Debug("DownLoadGeneric", baseUrl)
	// 创建HTTP请求对象
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return fmt.Errorf("创建请求失败: %v\n", err)
	}
	// 添加Basic认证头
	credentials := fmt.Sprintf("%s:%s", username, password) // 替换实际用户名密码
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	req.Header.Add("Authorization", "Basic "+encodedCredentials)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求发送失败: %v\n", err)
		return fmt.Errorf("请求发送失败: %v\n", err)
	}
	defer resp.Body.Close()
	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("认证失败或资源不可用，状态码: %d\n", resp.StatusCode)
		return fmt.Errorf("认证失败或资源不可用，状态码: %d\n", resp.StatusCode)
	}

	// 创建本地文件
	file, err := os.Create(destPath)
	if err != nil {
		fmt.Printf("文件创建失败: %v\n", err)
		return fmt.Errorf("文件创建失败: %v\n", err)
	}
	defer file.Close()

	// 流式写入文件（避免内存溢出）
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("文件下载失败: %v\n", err)
		return fmt.Errorf("文件下载失败: %v\n", err)
	}
	fmt.Println("文件下载成功！")
	return nil
}

func UploadGeneric(baseUrl, method, cfgpath string, username, password string) error {
	glog.Debug("UploadGeneric", baseUrl)
	file, err := os.Open(cfgpath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return err
	}
	defer file.Close()

	// 读取文件内容
	fileData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("无法读取文件内容:", err)
		return err
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(fileData))
	if err != nil {
		fmt.Println("无法创建请求:", err)
		return err
	}
	// 设置基本认证
	req.SetBasicAuth(username, password)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求发送失败:", err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("无法读取响应内容:", err)
		return err
	}

	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(respBody))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("上传失败，状态码: %d", resp.StatusCode)
	}

	//var result map[string]interface{}
	//err = json.Unmarshal(respBody, &result)
	//if err != nil {
	//	fmt.Println("json解析错误:", err)
	//	return err
	//}
	//status := result["status"].(float64)
	//if status != 200 {
	//	return fmt.Errorf("coding code %d", status)
	//}
	return nil
}
