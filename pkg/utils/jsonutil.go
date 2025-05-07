package utils

import (
	"encoding/gob"
	"encoding/json"
	"os"
)

// SaveToFile 序列化并保存
func SaveToFile[T any](obj T, filename string) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromFile 从文件加载并反序列化
func LoadFromFile[T any](filename string) (T, error) {
	var p T
	data, err := os.ReadFile(filename)
	if err != nil {
		return p, err
	}
	if err := json.Unmarshal(data, &p); err != nil {
		return p, err
	}
	return p, nil
}

func SaveWithGob[T any](obj T, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(obj)
}

func LoadWithGob[T any](filename string) (T, error) {
	var p T
	file, err := os.Open(filename)
	if err != nil {
		return p, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&p); err != nil {
		return p, err
	}
	return p, nil
}
