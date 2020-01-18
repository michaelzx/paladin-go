package utils

import (
	"errors"
	"reflect"
)

func IsZeroValue(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
func StructName(structPtr interface{}) (string, error) {
	if err := IsStructPtr(structPtr); err != nil {
		return "", err
	}
	typeName := reflect.TypeOf(structPtr).Name()
	return typeName, nil
}
func StructFields(t reflect.Type) []reflect.StructField {
	rte := t.Elem()
	var fields = make([]reflect.StructField, rte.NumField())
	for i := 0; i < rte.NumField(); i++ {
		field := rte.Field(i)
		fields = append(fields, field)
	}
	return fields
}
func StructFieldIdxMap(t reflect.Type) map[string]int {
	rte := t.Elem()
	fieldMap := make(map[string]int)
	for i := 0; i < rte.NumField(); i++ {
		field := rte.Field(i)
		fieldMap[field.Name] = i
	}
	return fieldMap
}

func IsPtr(source interface{}) bool {
	rv := reflect.ValueOf(source)
	return rv.Kind() == reflect.Ptr
}
func IsStruct(source interface{}) bool {
	rv := reflect.ValueOf(source)
	rv = rv.Elem()
	return rv.Kind() == reflect.Struct
}
func IsMap(source interface{}) bool {
	rv := reflect.ValueOf(source)
	rv = rv.Elem()
	return rv.Kind() == reflect.Map
}

func IsStructPtr(source interface{}) error {
	rv := reflect.ValueOf(source)
	if rv.Kind() != reflect.Ptr {
		return nil
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return errors.New("variable must be struct")
	}
	return nil
}
