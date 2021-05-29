package assert

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

// assert val1 is equal to val2
func Equal(t *testing.T, val1, val2 interface{}) {
	EqualSkip(t, 2, val1, val2)
}

func EqualSkip(t *testing.T, skip int, val1, val2 interface{}) bool {
	if !IsEqual(val1, val2) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v does not equal %v\n", path.Base(file), line, val1, val2)
		t.Fail()
	}
	return true
}

// assert val1 is not equal to val2
func NotEqual(t *testing.T, val1, val2 interface{}) {
	NotEqualSkip(t, 2, val1, val2)
}

func NotEqualSkip(t *testing.T, skip int, val1, val2 interface{}) {

	if IsEqual(val1, val2) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v should not be equal %v\n", path.Base(file), line, val1, val2)
		t.Fail()
	}
}

func IsEqual(val1, val2 interface{}) bool {
	v1 := reflect.ValueOf(val1)
	v2 := reflect.ValueOf(val2)

	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}

	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}

	if !v1.IsValid() && !v2.IsValid() {
		return true
	}

	switch v1.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v1.IsNil() {
			v1 = reflect.ValueOf(nil)
		}
	}

	switch v2.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v2.IsNil() {
			v2 = reflect.ValueOf(nil)
		}
	}

	v1Underlying := reflect.Zero(reflect.TypeOf(v1)).Interface()
	v2Underlying := reflect.Zero(reflect.TypeOf(v2)).Interface()

	if v1 == v1Underlying {
		if v2 == v2Underlying {
			goto CASE4
		} else {
			goto CASE3
		}
	} else {
		if v2 == v2Underlying {
			goto CASE2
		} else {
			goto CASE1
		}
	}

CASE1:
	return reflect.DeepEqual(v1.Interface(), v2.Interface())
CASE2:
	return reflect.DeepEqual(v1.Interface(), v2)
CASE3:
	return reflect.DeepEqual(v1, v2.Interface())
CASE4:
	return reflect.DeepEqual(v1, v2)
}

// assert map has key
func MapHasKey(t *testing.T, m interface{}, i interface{}) {
	ValueIsNullSkip(t, m, 2)
	m1 := interfaceToMapSkip(t, m, 2)
	//m1 := m.(map[interface{}]interface{})
	if !TypeIsComparableSkip(t, i, 2) {
		return
	}
	_, ok := m1[i]
	if !ok {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d %v map has not key %v\n", path.Base(file), line, m, i)
		t.Fail()
	}
}

// assert map has value
func MapHasValue(t *testing.T, m interface{}, i interface{}) {
	ValueIsNullSkip(t, m, 2)

	if !TypeIsComparableSkip(t, i, 2) {
		return
	}

	m1 := interfaceToMapSkip(t, m, 2)
	mv := reflect.ValueOf(m1)
	for _, key := range mv.MapKeys() {
		if IsEqual(mv.MapIndex(key).Interface(), i) {
			return
		}
	}

	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%d %v map not has value %v \n", path.Base(file), line, m, i)
	t.Fail()
}

// assert map not has key
func MapNotHasKey(t *testing.T, m interface{}, i interface{}) {
	ValueIsNullSkip(t, m, 2)

	m1 := interfaceToMapSkip(t, m, 2)

	if !TypeIsComparableSkip(t, i, 2) {
		return
	}
	_, ok := m1[i]
	if ok {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d %v map has key %v \n", path.Base(file), line, m, i)
		t.Fail()
	}
}

// map only has same keys
func MapHasKeysOnly(t *testing.T, m interface{}, keys []interface{}) {
	m1 := interfaceToMapSkip(t, m, 2)
	allKeys := getMapKeys(m1)
	for _, key := range allKeys {
		if !contains(keys, key) {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("%s:%d %v not only contains keys %v \n", path.Base(file), line, m, keys)
			t.Fail()
		}
	}
}

// assert slice container
func Contains(t *testing.T, l interface{}, v interface{}) {
	s := interfaceToSliceSkip(t, l, 2)
	if contains(s ,v) {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%d %v not contains %v \n", path.Base(file), line, l, v)
	t.Fail()
}

func contains(l []interface{}, v interface{}) bool {
	for _, item := range l {
		if IsEqual(v, item) {
			return true
		}
	}
	return false
}

// assert slice not contains v
func NotContains(t *testing.T, l interface{}, v interface{}) {
	s := interfaceToSliceSkip(t, l, 2)
	if contains(s ,v) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d %v contains %v \n", path.Base(file), line, l, v)
		t.Fail()
	}
}

// assert string s contains substr
func StringContainsString(t *testing.T, s, substr string) {
	if !strings.Contains(s, substr) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d string '%v' not container string '%v'\n", path.Base(file), line, s, substr)
		t.Fail()
	}
}

// assert string s ignore case contains substr
func StringContainsStringIgnoringCase(t *testing.T, s, substr string) {
	s1 := strings.ToLower(s)
	substr1 := strings.ToLower(substr)
	if !strings.Contains(s1, substr1) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d string '%v' ignore case not container string '%v' \n", path.Base(file), line, s, substr)
		t.Fail()
	}
}

// assert string s not contains substr
func StringNotContainsString(t *testing.T, s, substr string) {
	if strings.Contains(s, substr) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d string '%v' container string '%v' \n", path.Base(file), line, s, substr)
		t.Fail()
	}
}

// assert string s ignore case not contains substr
func StringNotContainsStringIgnoringCase(t *testing.T, s, substr string) {
	s1 := strings.ToLower(s)
	substr1 := strings.ToLower(substr)
	if strings.Contains(s1, substr1) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d string '%v' ingore case container string '%v' \n", path.Base(file), line, s, substr)
		t.Fail()
	}
}

func TypeIsComparable(t *testing.T, i interface{})  {
	TypeIsComparableSkip(t, i, 2)
}

func TypeIsComparableSkip(t *testing.T, i interface{}, skip int) bool {
	if nil == i || reflect.TypeOf(i).Comparable() {
		return true
	}
	_, file, line, _ := runtime.Caller(skip)
	fmt.Printf("%s:%d %v is not comparable\n", path.Base(file), line, i)
	t.Fail()
	return false
}

func ValueIsNullSkip(t *testing.T, i interface{}, skip int)  {
	isNull := false
	if nil == i {
		isNull = true
	} else if !reflect.ValueOf(i).IsValid() {
		isNull = true

	} else {
		switch reflect.TypeOf(i).Kind() {
		case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
			isNull = reflect.ValueOf(i).IsNil()
		}
	}
	if isNull {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v is not null\n", path.Base(file), line, i)
		t.Fail()
	}
}

type InterfaceStructure struct {
	pt uintptr
	pv uintptr
}

// transform interface{} to InterfaceStructure
func AsInterfaceStructure(i interface{}) InterfaceStructure {
	return *(*InterfaceStructure)(unsafe.Pointer(&i))
}


// get all keys of map
func getMapKeys(m map[interface{}]interface{}) []interface{} {j := 0
	keys := make([]interface{}, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func interfaceToMapSkip(t *testing.T,i interface{}, skip int) map[interface{}]interface{} {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Map {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d  '%v' is not map type is %s\n", path.Base(file), line, i, v.Kind().String())
		t.Fail()
	}
	m := make(map[interface{}]interface{})
	for _, key := range v.MapKeys() {
		m[key.Interface()] = v.MapIndex(key).Interface()
	}
	return m
}

func interfaceToSliceSkip(t *testing.T, i interface{}, skip int) []interface{} {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d  '%v' is not slice、array type is %s\n", path.Base(file), line, i, v.Kind().String())
		t.Fail()
	}
	m := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		m[i] = v.Index(i).Interface()
	}
	return m
}

func isMap(i interface{}) bool {
	v := reflect.ValueOf(i)
	return v.Kind() == reflect.Map
}