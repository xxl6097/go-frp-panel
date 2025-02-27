package utils

import (
	"archive/zip"
	"github.com/xxl6097/glog/glog"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Write(filePath string, content []byte) error {
	// 写入文件
	return os.WriteFile(filePath, content, 0644)
}
func Read(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteToml(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0o600)
}
func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0o600)
}

func ReadToml(filePath string) ([]byte, error) {
	bb, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return bb, nil
}

func AppendStringToFile(filePath, content string) error {
	// 以追加模式打开文件，如果文件不存在则创建
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	// 确保在函数结束时关闭文件
	defer file.Close()
	// 追加字符串内容到文件
	_, err = io.WriteString(file, content)
	return err
}

func Delete(filePath string, args ...string) {
	var title string
	if args != nil && len(args) > 0 {
		title = args[0]
	}
	if err := os.Remove(filePath); err != nil {
		glog.Infof("%s 文件删除失败: %s,%v\n", title, filePath, err)
		return
	}
	glog.Infof("%s 文件删除成功: %s\n", title, filePath)
}

func DirCheck(path string) error {
	// 检查目录是否存在
	if _, err := os.Stat(path); err == nil {
		// 存在
		return nil
	} else if !os.IsNotExist(err) {
		// 其他错误
		return err
	}
	// 不存在，创建
	return os.MkdirAll(path, 0755)
}

func EnsureDir(path string) error {
	// 检查目录是否存在
	if _, err := os.Stat(path); err == nil {
		// 存在，删除
		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
		return os.MkdirAll(path, 0755)
	} else if !os.IsNotExist(err) {
		// 其他错误
		return err
	}
	// 不存在，创建
	return os.MkdirAll(path, 0755)
}

// Unzip 函数用于解压指定的ZIP文件到指定的目标目录
func Unzip(src, dest string) error {
	// 打开ZIP文件
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// 遍历ZIP文件中的每个文件或目录
	for _, f := range r.File {
		// 构建目标文件的路径
		fpath := filepath.Join(dest, f.Name)

		// 如果是目录，创建目录
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// 如果是文件，创建文件并写入内容
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		// 将ZIP文件中的内容复制到目标文件中
		_, err = io.Copy(outFile, rc)

		// 关闭文件和读取器
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func Zip(dir, dst string) error {
	// 创建目标ZIP文件
	zipFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 初始化ZIP写入器
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// 遍历目录
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过根目录自身
		if path == dir {
			return nil
		}

		// 创建ZIP条目头信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 修正条目名称（保留相对路径）
		header.Name = strings.TrimPrefix(strings.Replace(path, dir, "", 1), string(filepath.Separator))
		if info.IsDir() {
			header.Name += "/" // 目录需以/结尾
		} else {
			header.Method = zip.Deflate // 启用压缩
		}

		// 写入条目头
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		// 若是文件，则写入内容
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			return err
		}
		return nil
	})
}

// 根据后缀判断文件类型（仅后缀匹配）
func IsZipOrJson(filename string) (isZip, isJSON bool) {
	ext := strings.ToLower(filepath.Ext(filename)) // 统一转为小写
	isZip = ext == ".zip"
	isJSON = ext == ".json"
	return isZip, isJSON
}
