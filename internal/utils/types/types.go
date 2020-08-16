package types

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Byte(value interface{}) byte {
	return Uint8(value)
}

func Int(value interface{}) int {
	return int(Int32(value))
}

func Int8(value interface{}) int8 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return int8(value)
	case int8:
		return int8(value)
	case int16:
		return int8(value)
	case int32:
		return int8(value)
	case int64:
		return int8(value)
	case uint:
		return int8(value)
	case uint8:
		return int8(value)
	case uint16:
		return int8(value)
	case uint32:
		return int8(value)
	case uint64:
		return int8(value)
	case float32:
		return int8(value)
	case float64:
		return int8(value)
	case string:
		var result, err = strconv.ParseInt(value, 10, 64)
		if err == nil {
			return int8(result)
		} else {
			floatResult, err := strconv.ParseFloat(value, 64)
			if err == nil {
				return int8(floatResult)
			}
		}
	}
	return 0
}

func Int16(value interface{}) int16 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return int16(value)
	case int8:
		return int16(value)
	case int16:
		return int16(value)
	case int32:
		return int16(value)
	case int64:
		return int16(value)
	case uint:
		return int16(value)
	case uint8:
		return int16(value)
	case uint16:
		return int16(value)
	case uint32:
		return int16(value)
	case uint64:
		return int16(value)
	case float32:
		return int16(value)
	case float64:
		return int16(value)
	case string:
		var result, err = strconv.ParseInt(value, 10, 64)
		if err == nil {
			return int16(result)
		} else {
			floatResult, err := strconv.ParseFloat(value, 64)
			if err == nil {
				return int16(floatResult)
			}
		}
	}
	return 0
}

func Int64(value interface{}) int64 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return int64(value)
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case string:
		var result, err = strconv.ParseInt(value, 10, 64)
		if err == nil {
			return result
		} else {
			floatResult, err := strconv.ParseFloat(value, 64)
			if err == nil {
				return int64(floatResult)
			}
		}
	}
	return 0
}

func Uint64(value interface{}) uint64 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case int64:
		return uint64(value)
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return uint64(value)
	case float32:
		return uint64(value)
	case float64:
		return uint64(value)
	case string:
		var result, err = strconv.ParseInt(value, 10, 64)
		if err == nil {
			return uint64(result)
		} else {
			floatResult, err := strconv.ParseFloat(value, 64)
			if err == nil {
				return uint64(floatResult)
			}
		}
	}
	return 0
}

func Int32(value interface{}) int32 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return int32(value)
	case int8:
		return int32(value)
	case int16:
		return int32(value)
	case int32:
		return int32(value)
	case int64:
		return int32(value)
	case uint:
		return int32(value)
	case uint8:
		return int32(value)
	case uint16:
		return int32(value)
	case uint32:
		return int32(value)
	case uint64:
		return int32(value)
	case float32:
		return int32(value)
	case float64:
		return int32(value)
	case string:
		var result, err = strconv.ParseInt(value, 10, 32)
		if err == nil {
			return int32(result)
		} else {
			floatResult, err := strconv.ParseFloat(value, 32)
			if err == nil {
				return int32(floatResult)
			}
		}
	}
	return 0
}

func Uint(value interface{}) uint {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return uint(value)
	case int8:
		return uint(value)
	case int16:
		return uint(value)
	case int32:
		return uint(value)
	case int64:
		return uint(value)
	case uint:
		return uint(value)
	case uint8:
		return uint(value)
	case uint16:
		return uint(value)
	case uint32:
		return uint(value)
	case uint64:
		return uint(value)
	case float32:
		return uint(value)
	case float64:
		return uint(value)
	case string:
		var result, err = strconv.ParseInt(value, 10, 32)
		if err == nil {
			return uint(result)
		} else {
			floatResult, err := strconv.ParseFloat(value, 32)
			if err == nil {
				return uint(floatResult)
			}
		}
	}
	return 0
}

func Uint8(value interface{}) uint8 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return uint8(value)
	case int8:
		return uint8(value)
	case int16:
		return uint8(value)
	case int32:
		return uint8(value)
	case int64:
		return uint8(value)
	case uint:
		return uint8(value)
	case uint8:
		return uint8(value)
	case uint16:
		return uint8(value)
	case uint32:
		return uint8(value)
	case uint64:
		return uint8(value)
	case float32:
		return uint8(value)
	case float64:
		return uint8(value)
	case string:
		var result, err = strconv.ParseInt(value, 10, 32)
		if err == nil {
			return uint8(result)
		} else {
			floatResult, err := strconv.ParseFloat(value, 32)
			if err == nil {
				return uint8(floatResult)
			}
		}
	}
	return 0
}

func Uint16(value interface{}) uint16 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return uint16(value)
	case int8:
		return uint16(value)
	case int16:
		return uint16(value)
	case int32:
		return uint16(value)
	case int64:
		return uint16(value)
	case uint:
		return uint16(value)
	case uint8:
		return uint16(value)
	case uint16:
		return uint16(value)
	case uint32:
		return uint16(value)
	case uint64:
		return uint16(value)
	case float32:
		return uint16(value)
	case float64:
		return uint16(value)
	case string:
		var result, err = strconv.ParseInt(value, 10, 32)
		if err == nil {
			return uint16(result)
		} else {
			floatResult, err := strconv.ParseFloat(value, 32)
			if err == nil {
				return uint16(floatResult)
			}
		}
	}
	return 0
}

func Uint32(value interface{}) uint32 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return uint32(value)
	case int8:
		return uint32(value)
	case int16:
		return uint32(value)
	case int32:
		return uint32(value)
	case int64:
		return uint32(value)
	case uint:
		return uint32(value)
	case uint8:
		return uint32(value)
	case uint16:
		return uint32(value)
	case uint32:
		return uint32(value)
	case uint64:
		return uint32(value)
	case float32:
		return uint32(value)
	case float64:
		return uint32(value)
	case string:
		var result, err = strconv.ParseInt(value, 10, 32)
		if err == nil {
			return uint32(result)
		} else {
			floatResult, err := strconv.ParseFloat(value, 32)
			if err == nil {
				return uint32(floatResult)
			}
		}
	}
	return 0
}

func Int32Value(value interface{}) (int32, error) {
	if value == nil {
		return 0, errors.New("value should not be nil")
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return int32(value), nil
	case int8:
		return int32(value), nil
	case int16:
		return int32(value), nil
	case int32:
		return int32(value), nil
	case int64:
		return int32(value), nil
	case uint:
		return int32(value), nil
	case uint8:
		return int32(value), nil
	case uint16:
		return int32(value), nil
	case uint32:
		return int32(value), nil
	case uint64:
		return int32(value), nil
	case float32:
		return int32(value), nil
	case float64:
		return int32(value), nil
	case string:
		var result, err = strconv.ParseInt(value, 10, 32)
		if err == nil {
			return int32(result), nil
		} else {
			floatResult, err := strconv.ParseFloat(value, 32)
			if err == nil {
				return int32(floatResult), nil
			}
			return 0, err
		}
	}
	return 0, nil
}

func Float64(value interface{}) float64 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return float64(value)
	case int8:
		return float64(value)
	case int16:
		return float64(value)
	case int32:
		return float64(value)
	case int64:
		return float64(value)
	case uint:
		return float64(value)
	case uint8:
		return float64(value)
	case uint16:
		return float64(value)
	case uint32:
		return float64(value)
	case uint64:
		return float64(value)
	case float32:
		return float64(value)
	case float64:
		return float64(value)
	case string:
		floatResult, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return floatResult
		}
	}
	return 0
}

func Float32(value interface{}) float32 {
	if value == nil {
		return 0
	}

	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		return float32(value)
	case int8:
		return float32(value)
	case int16:
		return float32(value)
	case int32:
		return float32(value)
	case int64:
		return float32(value)
	case uint:
		return float32(value)
	case uint8:
		return float32(value)
	case uint16:
		return float32(value)
	case uint32:
		return float32(value)
	case uint64:
		return float32(value)
	case float32:
		return float32(value)
	case float64:
		return float32(value)
	case string:
		floatResult, err := strconv.ParseFloat(value, 32)
		if err == nil {
			return float32(floatResult)
		}
	}
	return 0
}

func Bool(value interface{}) bool {
	if value == nil {
		return false
	}

	var kind = reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Bool:
		return value.(bool)
	}
	return Int64(value) > 0
}

func String(value interface{}) string {
	if value == nil {
		return ""
	}
	valueString, ok := value.(string)
	if ok {
		return valueString
	}
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", value)
	}
	return fmt.Sprintf("%#v", value)
}

func Compare(value1 interface{}, value2 interface{}) bool {
	if value1 == nil {
		return false
	}

	switch value1 := value1.(type) {
	case bool:
		return Int(value1) > Int(value2)
	case int:
		return Int(value1) > Int(value2)
	case int8:
		return Int8(value1) > Int8(value2)
	case int16:
		return Int16(value1) > Int16(value2)
	case int32:
		return Int32(value1) > Int32(value2)
	case int64:
		return Int64(value1) > Int64(value2)
	case uint:
		return Uint(value1) > Uint(value2)
	case uint8:
		return Uint8(value1) > Uint8(value2)
	case uint16:
		return Uint16(value1) > Uint16(value2)
	case uint32:
		return Uint32(value1) > Uint32(value2)
	case uint64:
		return Uint64(value1) > Uint64(value2)
	case float32:
		return Float32(value1) > Float32(value2)
	case float64:
		return Float64(value1) > Float64(value2)
	case string:
		return String(value1) > String(value2)
	}
	return String(value1) > String(value2)
}

// 判断是否为数字
func IsNumber(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}
	return false
}

// 判断是否为整形数字
func IsInteger(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

// 判断是否为浮点型数字
func IsFloat(value interface{}) bool {
	switch value.(type) {
	case float32, float64:
		return true
	}
	return false
}

// 判断是否为nil
func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}

	return reflect.ValueOf(value).IsNil()
}

// 转换Slice类型
func Slice(fromSlice interface{}, toSliceType reflect.Type) (interface{}, error) {
	if fromSlice == nil {
		return nil, errors.New("'fromSlice' should not be nil")
	}

	fromValue := reflect.ValueOf(fromSlice)
	if fromValue.Kind() != reflect.Slice {
		return nil, errors.New("'fromSlice' should be slice")
	}

	if toSliceType.Kind() != reflect.Slice {
		return nil, errors.New("'toSliceType' should be slice")
	}

	v := reflect.Indirect(reflect.New(toSliceType))
	count := fromValue.Len()
	toElemKind := toSliceType.Elem().Kind()
	for i := 0; i < count; i++ {
		elem := fromValue.Index(i)
		elemVar := elem.Interface()
		switch toElemKind {
		case reflect.Int:
			v = reflect.Append(v, reflect.ValueOf(Int(elemVar)))
		case reflect.Int8:
			v = reflect.Append(v, reflect.ValueOf(Int8(elemVar)))
		case reflect.Int16:
			v = reflect.Append(v, reflect.ValueOf(Int16(elemVar)))
		case reflect.Int32:
			v = reflect.Append(v, reflect.ValueOf(Int32(elemVar)))
		case reflect.Int64:
			v = reflect.Append(v, reflect.ValueOf(Int64(elemVar)))
		case reflect.Uint:
			v = reflect.Append(v, reflect.ValueOf(Uint(elemVar)))
		case reflect.Uint8:
			v = reflect.Append(v, reflect.ValueOf(Uint8(elemVar)))
		case reflect.Uint16:
			v = reflect.Append(v, reflect.ValueOf(Uint16(elemVar)))
		case reflect.Uint32:
			v = reflect.Append(v, reflect.ValueOf(Uint32(elemVar)))
		case reflect.Uint64:
			v = reflect.Append(v, reflect.ValueOf(Uint64(elemVar)))
		case reflect.Bool:
			v = reflect.Append(v, reflect.ValueOf(Bool(elemVar)))
		case reflect.Float32:
			v = reflect.Append(v, reflect.ValueOf(Float32(elemVar)))
		case reflect.Float64:
			v = reflect.Append(v, reflect.ValueOf(Float64(elemVar)))
		case reflect.String:
			v = reflect.Append(v, reflect.ValueOf(String(elemVar)))
		}
	}
	return v.Interface(), nil
}
