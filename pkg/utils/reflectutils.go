package utils

import (
	"fmt"
	"reflect"
	"unsafe"
)

func GetFuncInstance[T any](name string, obj interface{}) *T {
	value := reflect.ValueOf(obj).Elem()
	// 通过方法名查找方法
	nameField := value.MethodByName(name)
	if nameField.IsValid() {
		nameField = nameField.Elem()
		instance := (*T)(unsafe.Pointer(nameField.UnsafeAddr()))
		return instance
	}
	return nil
}

func GetPointerInstance[T any](name string, obj interface{}) *T {
	value := reflect.ValueOf(obj).Elem()
	nameField := value.FieldByName(name)
	if nameField.IsValid() {
		nameField = nameField.Elem()
		instance := (*T)(unsafe.Pointer(nameField.UnsafeAddr()))
		return instance
	}
	return nil
}

func GetStructInstance[T any](name string, obj interface{}) *T {
	value := reflect.ValueOf(obj).Elem()
	nameField := value.FieldByName(name)
	if nameField.IsValid() {
		instance := (*T)(unsafe.Pointer(nameField.UnsafeAddr()))
		return instance
	}
	return nil
}

// SetFieldValue 使用反射给结构体的字段赋值
func SetFieldValue(obj interface{}, fieldName string, value interface{}) error {
	// 获取传入对象的反射值
	objValue := reflect.ValueOf(obj).Elem()
	// 获取字段的反射值
	field := objValue.FieldByName(fieldName)
	// 检查字段是否存在且可设置
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", fieldName)
	}
	//if !field.CanSet() {
	//	return fmt.Errorf("field %s is not settable", fieldName)
	//}
	// 将传入的值转换为反射值
	valueReflect := reflect.ValueOf(value)
	// 检查值的类型是否匹配
	if !valueReflect.Type().AssignableTo(field.Type()) {
		return fmt.Errorf("value type %v is not assignable to field type %v", valueReflect.Type(), field.Type())
	}
	// 给字段赋值
	field.Set(valueReflect)
	return nil
}
