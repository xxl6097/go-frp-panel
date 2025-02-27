package utils

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"io"
	"os"
	"strings"
)

// PrintByteArrayAsConstant 把字节数组以常量字节数组的形式打印出来
func PrintByteArrayAsConstant(bytes []byte) string {
	sb := strings.Builder{}
	sb.WriteString("[]byte{")
	for i, b := range bytes {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("0x%02X", b))
	}
	sb.WriteString("}")
	return sb.String()
}

// HexEscapeToBytes 将十六进制的转义序列转换回字节数组
func HexEscapeToBytes(hexEscape string) ([]byte, error) {
	// 移除所有的 \x 前缀
	hexStr := strings.ReplaceAll(hexEscape, "\\x", "")
	// 解码十六进制字符串为字节数组
	return hex.DecodeString(hexStr)
}

func BytesToHexEscape(bytes []byte) string {
	result := ""
	for _, b := range bytes {
		result += fmt.Sprintf("\\x%02x", b)
	}
	return result
}

func PrintProgress(index int64, totalSteps int64) {
	barLength := 50
	// 计算进度条填充的字符数量
	filledLength := int(float64(barLength) * float64(index) / float64(totalSteps))
	// 生成进度条字符串
	bar := ""
	for j := 0; j < filledLength; j++ {
		bar += "="
	}
	for j := 0; j < barLength-filledLength; j++ {
		bar += " "
	}
	progress := int(float64(index) / float64(totalSteps) * 100)
	// 打印进度条和百分比
	glog.Printf("\r[%s] %d%%", bar, progress)
}

func GenerateBin1(srcFile, dstFile string, b byte, size int, cfgBytes []byte) error {
	if size <= 0 || cfgBytes == nil || len(cfgBytes) <= 0 {
		src, err := os.Open(srcFile) // can not use args[0], on Windows call openp2p is ok(=openp2p.exe)
		if err != nil {
			return err
		}
		defer src.Close()
		//将本程序复制到目标为止，目标文件名称为配置文件的名称
		dst, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0775)
		if err != nil {
			return err
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
		return nil
	}
	tpl, er := os.Open(srcFile)
	if er != nil {
		return er
	}
	defer tpl.Close()
	// 创建目标文件
	dst, er1 := os.Create(dstFile)
	if er1 != nil {
		return er1
	}
	defer dst.Close()
	writer := bufio.NewWriter(dst)
	cfgBuffer := bytes.Repeat([]byte{b}, size)
	prevBuffer := make([]byte, 0)
	for {
		thisBuffer := make([]byte, 1024)
		n, err := tpl.Read(thisBuffer)
		thisBuffer = thisBuffer[:n]
		tempBuffer := append(prevBuffer, thisBuffer...)
		bufIndex := bytes.Index(tempBuffer, cfgBuffer)
		if bufIndex > -1 {
			tempBuffer = bytes.Replace(tempBuffer, cfgBuffer, cfgBytes, -1)
		}
		temp := tempBuffer[:len(prevBuffer)]
		// 将缓冲区的数据写入目标文件
		_, errr := writer.Write(temp)
		if errr != nil {
			break
		}
		prevBuffer = tempBuffer[len(prevBuffer):]
		if err != nil {
			break
		}
	}
	if len(prevBuffer) > 0 {
		_, err1 := writer.Write(prevBuffer)
		if err1 != nil {
			return err1
		}
		prevBuffer = nil
	}
	return writer.Flush()
}

func GenerateBin2(srcFile, dstFile string, oldCfgBytes, cfgBytes []byte) error {
	if oldCfgBytes == nil || len(oldCfgBytes) <= 0 || cfgBytes == nil || len(cfgBytes) <= 0 {
		return fmt.Errorf("oldCfgBytes or cfgBytes is nil")
	}
	tpl, er := os.Open(srcFile)
	if er != nil {
		glog.Printf("打开源文件失败：%v\n", er)
		return er
	}
	defer tpl.Close()
	// 创建目标文件
	dst, er1 := os.Create(dstFile)
	if er1 != nil {
		glog.Printf("创建目标文件失败：%v\n", er)
		return er1
	}
	defer dst.Close()
	var fileSize int64
	if stat, err := tpl.Stat(); err == nil {
		fileSize = stat.Size()
		sizeB := float64(stat.Size()) / 1024 / 1024
		glog.Printf("正在安装%s[大小：%.2fMB]到%s\n", stat.Name(), sizeB, dstFile)
	}

	writer := bufio.NewWriter(dst)
	prevBuffer := make([]byte, 0)
	var tempSize int64
	var isReplace bool
	var errMsg error
	tempProgress := -1
	index := -1
	for {
		thisBuffer := make([]byte, 1024)
		n, err := tpl.Read(thisBuffer)
		if err == io.EOF {
			break
		} else if err != nil {
			glog.Printf("\n\n读取源文件数据失败[%s]，读取大小：%d，错误信息：%v\n", srcFile, n, err)
			errMsg = err
			return err
		}
		tempSize += int64(n)
		thisBuffer = thisBuffer[:n]
		tempBuffer := append(prevBuffer, thisBuffer...)
		index = bytes.Index(tempBuffer, oldCfgBytes)
		if index > -1 {
			//glog.Println("找到密钥位置了，替换...")
			tempBuffer = bytes.Replace(tempBuffer, oldCfgBytes, cfgBytes, -1)
			isReplace = true
		}
		temp := tempBuffer[:len(prevBuffer)]

		//PrintProgress(tempSize, fileSize)
		progress := int(float64(tempSize) / float64(fileSize) * 100)
		if progress >= tempProgress {
			glog.Println("安装进度", progress, "%")
			tempProgress = progress
			tempProgress += 5
		}

		// 将缓冲区的数据写入目标文件
		nn, errr := writer.Write(temp)
		if errr == io.EOF {
			break
		} else if errr != nil {
			glog.Printf("\n\n将缓冲区的数据写入目标文件失败：%d %v\n", nn, errr)
			errMsg = errr
			return errr
		}
		prevBuffer = tempBuffer[len(prevBuffer):]
	}
	glog.Println()
	if !isReplace {
		glog.Println("密钥未被替换", index)
	}
	if len(prevBuffer) > 0 {
		_, err1 := writer.Write(prevBuffer)
		if err1 != nil {
			return err1
		}
		prevBuffer = nil
	}

	errMsg = writer.Flush()
	// 给文件赋予执行权限（0755）
	errMsg = os.Chmod(dstFile, 0755)
	if errMsg != nil {
		glog.Printf("赋予文件执行权限时出错: %v\n", errMsg)
	}
	return errMsg
}

func GenerateBin(scrFilePath, dstFilePath string, oldBytes, newBytes []byte) error {
	// 打开原文件
	srcFile, err := os.Open(scrFilePath)
	if err != nil {
		return fmt.Errorf("无法打开文件: %v[%s]", err, scrFilePath)
	}
	defer srcFile.Close()

	var srcFileSize int64
	if stat, err := srcFile.Stat(); err == nil {
		srcFileSize = stat.Size()
		sizeB := float64(stat.Size()) / 1024 / 1024
		glog.Printf("%s[大小：%.2fMB]%s\n", stat.Name(), sizeB, dstFilePath)
	}

	tmpFile, err := os.Create(dstFilePath)
	if err != nil {
		return fmt.Errorf("无法创建临时文件: %v[%s]", err, dstFilePath)
	}
	defer tmpFile.Close()

	reader := bufio.NewReader(srcFile)
	prevBuffer := make([]byte, 0)
	isReplace := false
	var indexSize int64
	newFileSize := int64(0)
	tempProgress := -1
	for {
		thisBuffer := make([]byte, Divide(len(oldBytes), 1024))
		n, err2 := reader.Read(thisBuffer)
		if err2 != nil && err2 != io.EOF {
			return fmt.Errorf("读取文件时出错: %v[%s]", err2, scrFilePath)
		}
		indexSize += int64(n)
		thisBuffer = thisBuffer[:n]
		tempBuffer := append(prevBuffer, thisBuffer...)
		index := bytes.Index(tempBuffer, oldBytes)
		if index > -1 {
			glog.Printf("找到位置[%d]了，签名...\n", index)
			isReplace = true
			tempBuffer = bytes.Replace(tempBuffer, oldBytes, newBytes, -1)
		}
		// 写入前一次的
		writeSize, err1 := tmpFile.Write(tempBuffer[:len(prevBuffer)])
		if err1 != nil {
			return fmt.Errorf("1写入临时文件时出错: %v[%s]", err1, dstFilePath)
		}

		newFileSize += int64(writeSize)
		progress := int(float64(indexSize) / float64(srcFileSize) * 100)
		if progress >= tempProgress {
			glog.Printf("程序签名:%v%s\n", progress, "%")
			tempProgress = progress
			tempProgress += 5
		}

		//前一次的+本次的转给 prev
		prevBuffer = tempBuffer[len(prevBuffer):]
		//if err != nil {
		//	break
		//}
		if n == 0 || err2 != nil {
			break // 文件读取完毕
		}
	}
	if len(prevBuffer) > 0 {
		writeSize, err1 := tmpFile.Write(prevBuffer)
		if err1 != nil {
			return fmt.Errorf("2写入临时文件时出错: %v[%s]", err1, dstFilePath)
		}
		newFileSize += int64(writeSize)
		prevBuffer = nil
	}
	glog.Printf("原始文件大小：%d  %s\n", indexSize, scrFilePath)
	glog.Printf("目标文件大小：%d  %s\n", indexSize, dstFilePath)
	// 给文件赋予执行权限（0755）
	errMsg := os.Chmod(dstFilePath, 0755)
	if errMsg != nil {
		return fmt.Errorf("赋予文件执行权限时出错: %v\n", errMsg)
	}
	if !isReplace {
		glog.Printf("oldBytes[%d]--->%v\n", len(oldBytes), oldBytes)
		glog.Printf("newBytes[%d]--->%v\n", len(newBytes), newBytes)
		return errors.New("位置没找到，数据未替换😭")
	}
	err1 := srcFile.Close()
	if err1 != nil {
		glog.Error("srcFile.Close", err1)
	}
	err1 = tmpFile.Close()
	if err1 != nil {
		glog.Error("tmpFile.Close", err1)
	}

	return nil
}
