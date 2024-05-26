package common

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

func DataToStructByTagSql(data map[string]string, obj interface{}) {
	// get the struct
	objValue := reflect.ValueOf(obj).Elem()
	// convert type and map data into struct
	for i := 0; i < objValue.NumField(); i++ {
		// get value of sql: data["ID"]
		value := data[objValue.Type().Field(i).Tag.Get("sql")]
		// get the field name: "ID"
		name := objValue.Type().Field(i).Name
		// get the field type: int64
		structFieldType := objValue.Field(i).Type()
		// get type of value; alternatively, directly assign as "string"
		val := reflect.ValueOf(value)

		// Conversion
		var err error
		if structFieldType != val.Type() {
			val, err = TypeConversion(value, structFieldType.Name()) //类型转换
			if err != nil {

			}
		}
		//设置类型值
		objValue.FieldByName(name).Set(val)
	}
}

func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	return reflect.ValueOf(value), errors.New("Unknown type：" + ntype)
}
