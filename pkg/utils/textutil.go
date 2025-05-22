package utils

import (
	"encoding/json"
	"fmt"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/pelletier/go-toml/v2"
	"reflect"
	"unicode"
)

func StringContains(element string, data []string) bool {
	for _, v := range data {
		if element == v {
			return true
		}
	}
	return false
}

// filterData 递归过滤空字段、0字段和false字段，并确保整型字段没有小数点
func filterData(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			filteredValue := filterData(value)
			if !isEmpty(filteredValue) {
				result[key] = filteredValue
			}
		}
		return result
	case []interface{}:
		var result []interface{}
		for _, item := range v {
			filteredItem := filterData(item)
			if !isEmpty(filteredItem) {
				result = append(result, filteredItem)
			}
		}
		return result
	case float64:
		// 如果浮点数没有小数部分，则转换为整数
		if v == float64(int(v)) {
			return int(v)
		}
		return v
	default:
		return v
	}
}

// isEmpty 判断字段是否为空、0或false
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len() == 0
	default:
		return false
	}
}
func jsonToToml(jBytes []byte) ([]byte, error) {
	var data map[string]interface{}
	// 解析 JSON 数据
	err := json.Unmarshal(jBytes, &data)
	if err != nil {
		return nil, err
	}
	// 过滤数据
	newData := filterData(data)
	// 转换为 TOML
	tomlBytes, err := toml.Marshal(newData)
	if err != nil {
		return nil, err
	}
	return tomlBytes, nil
}

func tomlToJson(tomlBytes []byte) ([]byte, error) {
	var data map[string]interface{}
	err := toml.Unmarshal(tomlBytes, &data)
	if err != nil {
		return nil, err
	}
	// 将 map 转换为 JSON 字节切片
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}
	// 将 JSON 字节切片转换为字符串
	return jsonBytes, nil
}

func TestToml() {
	cfg := &v1.ServerConfig{
		BindPort: 6000,
		BindAddr: "0.0.0.0",
		WebServer: v1.WebServerConfig{
			Addr:     "0.0.0.0",
			Port:     7500,
			User:     "admin",
			Password: "admin",
		},
	}
	tomlBytes := ObjectToTomlText(cfg)
	fmt.Println(string(tomlBytes))
	jstr, err := tomlToJson(tomlBytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jstr))
	v1Data := v1.ServerConfig{}
	err = TomlTextToObject(tomlBytes, &v1Data)
	fmt.Println(err, v1Data)

}

func TomlTextToObject(tomlBytes []byte, obj interface{}) error {
	jstr, err := tomlToJson(tomlBytes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(jstr))
	err = json.Unmarshal(jstr, &obj)
	return err
}

func ObjectToTomlText(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	tomlStr, err := jsonToToml(b)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return tomlStr
}

// ToUpperFirst 将字符串的首字母转换为大写
func ToUpperFirst(s string) string {
	if s == "" {
		return s
	}
	// 将字符串转换为符文切片
	r := []rune(s)
	// 将首字符转换为大写
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
