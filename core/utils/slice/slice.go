package slice

import (
	"reflect"
	"strings"

	"github.com/treeyh/soc-go-common/core/errors"
)

// SliceUniqueString 通过map去重slice
func SliceUniqueString(s []string) []string {
	size := len(s)
	if size == 0 {
		return []string{}
	}

	m := make(map[string]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]string, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}
	return ret
}

// SliceUniqueInt64 通过map去重slice
func SliceUniqueInt64(s []int64) []int64 {
	size := len(s)
	if size == 0 {
		return []int64{}
	}

	m := make(map[int64]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]int64, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}
	return ret
}

// EqualString 判断string是否存在数组中
func EqualString(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// ContainString 判断string是否包含在数组的string中
func ContainString(v string, sl []string) bool {
	for _, vv := range sl {
		if strings.Contains(v, vv) {
			return true
		}
	}
	return false
}

// HasPrefixString 判断string是否包含在前缀数组的string中
func HasPrefixString(v string, sl []string) bool {
	for _, vv := range sl {
		if strings.HasPrefix(v, vv) {
			return true
		}
	}
	return false
}

// ContainIface 判断对象是否存在数组中
func ContainIface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// Contain 判断对象是否包含在数组的string中
func Contain(list interface{}, obj interface{}) (bool, errors.AppError) {
	targetValue := reflect.ValueOf(list)
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
		return false, nil
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
		return false, nil
	}
	return false, errors.NewAppError(errors.ObjectNotArray)

}

// ToSlice 对象转list
func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
