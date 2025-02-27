package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func GenBin1() {
	// 文件路径
	filePath := "example.txt"
	// 临时文件路径
	tmpFilePath := "example.txt.tmp"

	// 打开原文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 创建临时文件
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		log.Fatalf("无法创建临时文件: %v", err)
	}
	defer tmpFile.Close()

	// 要查找的字节数组
	oldBytes := []byte("old_pattern_4000_bytes_...") // 替换为实际的 4000 字节数据
	// 要替换的字节数组
	newBytes := []byte("new_pattern_4000_bytes_...") // 替换为实际的 4000 字节数据

	// 使用 bufio.Reader 逐块读取文件
	reader := bufio.NewReader(file)
	buffer := make([]byte, 4096) // 4KB 缓冲区
	var window []byte            // 用于存储滑动窗口的数据

	for {
		// 读取一块数据
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatalf("读取文件时出错: %v", err)
		}
		if n == 0 {
			break // 文件读取完毕
		}

		// 将新读取的数据添加到滑动窗口
		window = append(window, buffer[:n]...)

		// 在滑动窗口中查找目标字节数组
		index := bytes.Index(window, oldBytes)
		if index != -1 {
			// 如果找到目标字节数组，写入窗口中找到目标之前的部分
			_, err := tmpFile.Write(window[:index])
			if err != nil {
				log.Fatalf("写入临时文件时出错: %v", err)
			}
			// 写入替换的字节数组
			_, err = tmpFile.Write(newBytes)
			if err != nil {
				log.Fatalf("写入临时文件时出错: %v", err)
			}
			// 更新滑动窗口，移除已处理的部分
			window = window[index+len(oldBytes):]
		}

		// 如果滑动窗口过大，写入部分数据以避免内存占用过高
		if len(window) > len(oldBytes)*2 {
			_, err := tmpFile.Write(window[:len(window)-len(oldBytes)])
			if err != nil {
				log.Fatalf("写入临时文件时出错: %v", err)
			}
			window = window[len(window)-len(oldBytes):]
		}

		// 如果文件读取完毕，写入剩余的滑动窗口数据
		if err == io.EOF {
			_, err := tmpFile.Write(window)
			if err != nil {
				log.Fatalf("写入临时文件时出错: %v", err)
			}
			break
		}
	}

	// 关闭文件
	file.Close()
	tmpFile.Close()

	// 将临时文件重命名为原文件
	err = os.Rename(tmpFilePath, filePath)
	if err != nil {
		log.Fatalf("无法重命名文件: %v", err)
	}

	fmt.Println("文件内容已成功替换。")
}

func ReplaceFileContent(scrFile, dstFile string, oldBytes, newBytes []byte) error {
	// 打开原文件
	file, err := os.Open(scrFile)
	if err != nil {
		return fmt.Errorf("无法打开文件: %v[%s]", err, scrFile)
	}
	defer file.Close()
	// 创建临时文件
	tmpFile, err := os.Create(dstFile)
	if err != nil {
		return fmt.Errorf("无法创建临时文件: %v[%s]", err, dstFile)
	}
	defer tmpFile.Close()
	// 使用 bufio.Reader 逐块读取文件
	reader := bufio.NewReader(file)
	buffer := make([]byte, 4096) // 4KB 缓冲区
	var window []byte            // 用于存储滑动窗口的数据

	for {
		// 读取一块数据
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("读取文件时出错: %v[%s]", err, scrFile)
		}
		if n == 0 {
			break // 文件读取完毕
		}
		// 将新读取的数据添加到滑动窗口
		window = append(window, buffer[:n]...)
		// 在滑动窗口中查找目标字节数组
		index := bytes.Index(window, oldBytes)
		if index != -1 {
			fmt.Printf("找到位置[%d]了，替换...\n", index)
			// 如果找到目标字节数组，写入窗口中找到目标之前的部分
			_, err := tmpFile.Write(window[:index])
			if err != nil {
				return fmt.Errorf("1写入临时文件时出错: %v[%s]", err, dstFile)
			}
			// 写入替换的字节数组
			_, err = tmpFile.Write(newBytes)
			if err != nil {
				return fmt.Errorf("2写入临时文件时出错: %v[%s]", err, dstFile)
			}
			// 更新滑动窗口，移除已处理的部分
			window = window[index+len(oldBytes):]
		}

		// 如果滑动窗口过大，写入部分数据以避免内存占用过高
		if len(window) > len(oldBytes)*2 {
			_, err := tmpFile.Write(window[:len(window)-len(oldBytes)])
			if err != nil {
				return fmt.Errorf("3写入临时文件时出错: %v[%s]", err, dstFile)
			}
			window = window[len(window)-len(oldBytes):]
		}

		// 如果文件读取完毕，写入剩余的滑动窗口数据
		if err == io.EOF {
			_, err := tmpFile.Write(window)
			if err != nil {
				return fmt.Errorf("4写入临时文件时出错: %v[%s]", err, dstFile)
			}
			break
		}
	}

	// 关闭文件
	file.Close()
	tmpFile.Close()
	// 将临时文件重命名为原文件
	//err = os.Rename(dstFile, scrFile)
	//if err != nil {
	//	return fmt.Errorf("无法重命名文件: %v[%s]", err, scrFile)
	//}
	fmt.Println("文件内容已成功替换。")
	return nil
}

func Find(scrFile string, oldBytes []byte) error {
	// 打开原文件
	file, err := os.Open(scrFile)
	if err != nil {
		return fmt.Errorf("无法打开文件: %v[%s]", err, scrFile)
	}
	defer file.Close()
	// 使用 bufio.Reader 逐块读取文件
	reader := bufio.NewReader(file)
	buffer := make([]byte, 4096) // 4KB 缓冲区
	var window []byte            // 用于存储滑动窗口的数据

	for {
		// 读取一块数据
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("读取文件时出错: %v[%s]", err, scrFile)
		}
		if n == 0 {
			break // 文件读取完毕
		}
		// 将新读取的数据添加到滑动窗口
		window = append(window, buffer[:n]...)
		// 在滑动窗口中查找目标字节数组
		index := bytes.Index(window, oldBytes)
		if index != -1 {
			fmt.Println("a->找到了。。。", index)
			// 更新滑动窗口，移除已处理的部分
			window = window[index+len(oldBytes):]
		}

		// 如果滑动窗口过大，写入部分数据以避免内存占用过高
		if len(window) > len(oldBytes)*2 {
			//fmt.Println("b->找到了。。。", index)
			window = window[len(window)-len(oldBytes):]
		}

		// 如果文件读取完毕，写入剩余的滑动窗口数据
		if err == io.EOF {
			break
		}
	}

	// 关闭文件
	file.Close()
	// 将临时文件重命名为原文件
	//err = os.Rename(dstFile, scrFile)
	//if err != nil {
	//	return fmt.Errorf("无法重命名文件: %v[%s]", err, scrFile)
	//}
	//fmt.Println("文件内容已成功替换。")
	return nil
}
