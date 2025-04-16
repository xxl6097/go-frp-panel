package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

var (
	ErrEntityInvalid      = errors.New(`common.ENTITY_INVALID`)
	ErrFailedVerification = errors.New(`common.ENTITY_CHECK_FAILED`)
)

// RemoveSlice 泛型删除所有匹配值的元素
func RemoveSlice[T comparable](slice []T, value T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

func If[T any](b bool, t, f T) T {
	if b {
		return t
	}
	return f
}

func Min[T int | int32 | int64 | uint | uint32 | uint64 | float32 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T int | int32 | int64 | uint | uint32 | uint64 | float32 | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func XOR(data []byte, key []byte) []byte {
	if len(key) == 0 {
		return data
	}
	for i := 0; i < len(data); i++ {
		data[i] = data[i] ^ key[i%len(key)]
	}
	return data
}

func GenRandByte(n int) []byte {
	secBuffer := make([]byte, n)
	rand.Reader.Read(secBuffer)
	return secBuffer
}

func GetStrUUID() string {
	return hex.EncodeToString(GenRandByte(16))
}

func GetUUID() []byte {
	return GenRandByte(16)
}

func GetMD5(data []byte) ([]byte, string) {
	hash := md5.New()
	hash.Write(data)
	result := hash.Sum(nil)
	hash.Reset()
	return result, hex.EncodeToString(result)
}

func FormatSize(size int64) string {
	sizes := []string{`B`, `KB`, `MB`, `GB`, `TB`, `PB`, `EB`, `ZB`, `YB`}
	i := 0
	for size >= 1024 && i < len(sizes)-1 {
		size /= 1024
		i++
	}
	return fmt.Sprintf(`%d%s`, size, sizes[i])
}

func BytesToString(b []byte, r ...int) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bytesPtr := sh.Data
	bytesLen := sh.Len
	switch len(r) {
	case 1:
		r[0] = If(r[0] > bytesLen, bytesLen, r[0])
		bytesLen -= r[0]
		bytesPtr += uintptr(r[0])
	case 2:
		r[0] = If(r[0] > bytesLen, bytesLen, r[0])
		bytesLen = If(r[1] > bytesLen, bytesLen, r[1]) - r[0]
		bytesPtr += uintptr(r[0])
	}
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: bytesPtr,
		Len:  bytesLen,
	}))
}

func StringToBytes(s string, r ...int) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	strPtr := sh.Data
	strLen := sh.Len
	switch len(r) {
	case 1:
		r[0] = If(r[0] > strLen, strLen, r[0])
		strLen -= r[0]
		strPtr += uintptr(r[0])
	case 2:
		r[0] = If(r[0] > strLen, strLen, r[0])
		strLen = If(r[1] > strLen, strLen, r[1]) - r[0]
		strPtr += uintptr(r[0])
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: strPtr,
		Len:  strLen,
		Cap:  strLen,
	}))
}

func GetSlicePrefix[T any](data *[]T, n int) *[]T {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(data))
	return (*[]T)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sliceHeader.Data,
		Len:  n,
		Cap:  n,
	}))
}

func GetSliceSuffix[T any](data *[]T, n int) *[]T {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(data))
	return (*[]T)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sliceHeader.Data + uintptr(sliceHeader.Len-n),
		Len:  n,
		Cap:  n,
	}))
}

func GetSliceChunk[T any](data *[]T, start, end int) *[]T {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(data))
	return (*[]T)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sliceHeader.Data + uintptr(start),
		Len:  end - start,
		Cap:  end - start,
	}))
}

func CheckBinaryPack(data []byte) (byte, byte, bool) {
	if len(data) >= 8 {
		if bytes.Equal(data[:4], []byte{34, 22, 19, 17}) {
			if data[4] == 20 || data[4] == 21 {
				return data[4], data[5], true
			}
		}
	}
	return 0, 0, false
}

func BytesToHexString(bytes []byte) string {
	var hexValues []string
	for _, b := range bytes {
		// 将每个字节转换为十六进制字符串，并添加 0x 前缀
		hexValues = append(hexValues, fmt.Sprintf("0x%02x", b))
	}
	// 使用逗号连接所有十六进制字符串
	return strings.Join(hexValues, ", ")
}

// DivideAndCeil 函数用于进行除法并向上取整
func DivideAndCeil(a, b int) int {
	// 将整数转换为 float64 类型进行除法运算
	result := float64(a) / float64(b)
	// 使用 math.Ceil 函数进行向上取整
	result = math.Ceil(result)
	// 将结果转换回整数类型
	return int(result)
}

func Divide(a, b int) int {
	return DivideAndCeil(a, b) * b
}

// IsWindows 判断是否为 Windows 系统
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func GetDataByJson[T any](r *http.Request) (*T, error) {
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func BindJSON[T any](r *http.Request) (*T, error) {
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func IsMacOs() bool {
	if strings.Compare(runtime.GOOS, "darwin") == 0 {
		return true
	}
	return false
}

func IsLinux() bool {
	if strings.Compare(runtime.GOOS, "linux") == 0 {
		return true
	}
	return false
}

func SplitVersion(v string) []string {
	// 去除前缀标识（如 "v1.2.3" → "1.2.3"）
	v = strings.TrimLeft(v, "v")
	return strings.Split(v, ".")
}

// CompareVersions 0:相等；1：v1>v2;-1:v1<v2
func CompareVersions(v1, v2 string) int {
	seg1 := SplitVersion(v1)
	seg2 := SplitVersion(v2)
	maxLen := int(math.Max(float64(len(seg1)), float64(len(seg2))))

	for i := 0; i < maxLen; i++ {
		num1 := getSegmentValue(seg1, i)
		num2 := getSegmentValue(seg2, i)

		if num1 > num2 {
			return 1 // v1 > v2
		} else if num1 < num2 {
			return -1 // v1 < v2
		}
	}
	return 0 // 相等
}

func getSegmentValue(seg []string, idx int) int {
	if idx >= len(seg) {
		return 0 // 自动补零处理长度不一致情况
	}
	num, _ := strconv.Atoi(seg[idx])
	return num
}

func GetVersionByFileName(filename string) string {
	re := regexp.MustCompile(`v\d+\.\d+\.\d+`)
	//fmt.Println(re.FindStringSubmatch(filename))
	return re.FindString(filename)
}

func ReplaceNewVersionBinName(filename, v string) string {
	re := regexp.MustCompile(`_v\d+\.\d+\.\d+_`)
	newName := re.ReplaceAllString(filename, fmt.Sprintf("_%s_", v)) // 替换为单个下划线
	fmt.Println(newName)
	return newName
}

func GetSelfSize() uint64 {
	// 获取当前可执行文件的路径
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("获取可执行文件路径时出错: %v\n", err)
		return 0
	}
	// 获取文件信息
	fileInfo, err := os.Stat(exePath)
	if err != nil {
		fmt.Printf("获取文件信息时出错: %v\n", err)
		return 0
	}

	// 获取文件大小
	fileSize := fileInfo.Size()
	fmt.Printf("本程序自身大小为: %d 字节\n", fileSize)
	return uint64(fileSize)
}
